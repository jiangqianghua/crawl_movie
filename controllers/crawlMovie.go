package controllers

import (
	"crawl_movie/models"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	// "time"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	models.ConnectRedis("127.0.0.1:6379")
	// 入口url
	sUrl := "https://movie.douban.com/subject/30171425/"
	// 添加入口url
	models.PutinQueue(sUrl)
	var movieInfo models.MovieInfo
	c.Ctx.WriteString("<html>")
	var index int64 = 0
	for {
		length := models.GetQueueLength()
		if length == 0 {
			break
		}

		sUrl = models.PopfromQueue()
		// 判断是否被访问过
		if models.IsVisit(sUrl) {
			continue
		}
		rsp := httplib.Get(sUrl)
		sMovieHtml, err := rsp.String()
		if err != nil {
			panic(err)
		}

		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
		if movieInfo.Movie_name != "" {
			movieInfo.Id = index
			movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_type = models.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_on_time = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_span = models.GetMovieRunningTime(sMovieHtml)
			if _, err := models.AddMovie(&movieInfo); err != nil {
				panic(err)
			}
			index++
		}

		urls := models.GetMovieUrls(sMovieHtml)

		for _, url := range urls {
			models.PutinQueue(url)
			c.Ctx.WriteString("<br>" + url + "</br>")
		}
		// 把当前url添加到已访问的队列中
		models.AddToSet(sUrl)
		//休息1秒
		//time.Sleep(time.Second)

	}
	c.Ctx.WriteString("<br>爬取完成。。。</br>")
	c.Ctx.WriteString("</html>")
	//c.Ctx.WriteString(models.GetMovieDirector(sMovieHtml) + " | ")
	//c.Ctx.WriteString(models.GetMovieName(sMovieHtml) + " | ")
	//c.Ctx.WriteString(models.GetMovieMainCharacters(sMovieHtml))
	//c.Ctx.WriteString(models.GetMovieGrade(sMovieHtml))
	// c.Ctx.WriteString(models.GetMovieOnTime(sMovieHtml))

	// var movieInfo models.MovieInfo
	// movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
	// movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
	// movieInfo.Movie_type = models.GetMovieGenre(sMovieHtml)
	// movieInfo.Movie_on_time = models.GetMovieOnTime(sMovieHtml)
	// movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
	// movieInfo.Movie_span = models.GetMovieRunningTime(sMovieHtml)
	// if _, err := models.AddMovie(&movieInfo); err != nil {
	// 	panic(err)
	// }
	//c.Ctx.WriteString(models.GetMovieDirector(sMovieHtml) + " | ")

	// urls := models.GetMovieUrls(sMovieHtml)

	// for _, url := range urls {
	// 	models.PutinQueue(url)
	// }

	//c.Ctx.WriteString(fmt.Sprintf("%v", models.GetMovieUrls(sMovieHtml)))
}
