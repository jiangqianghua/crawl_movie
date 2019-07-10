package controllers

import (
	"crawl_movie/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	sUrl := "https://movie.douban.com/subject/30171425/"
	rsp := httplib.Get(sUrl)
	sMovieHtml, err := rsp.String()
	if err != nil {
		panic(err)
	}
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

	models.ConnectRedis("127.0.0.1:6379")
	urls := models.GetMovieUrls(sMovieHtml)

	for _, url := range urls {
		models.PutinQueue(url)
	}

	c.Ctx.WriteString(fmt.Sprintf("%v", models.GetMovieUrls(sMovieHtml)))
}
