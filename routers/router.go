package routers

import (
	"crawl_movie/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//http://127.0.0.1:8080/crawl_movie
	beego.Router("/crawl_movie", &controllers.CrawlMovieController{}, "*:CrawlMovie")
}
