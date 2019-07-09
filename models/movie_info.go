package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                   int64
	Movie_id             int64
	Movie_name           string
	Movie_pic            string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
}

func init() {
	orm.Debug = true // 调试模式会打印sql日志
	orm.RegisterDataBase("default", "mysql", "root:Jiang123!@tcp(180.76.105.202:3306)/douban_db?charset=utf8", 30)
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func AddMovie(movie_info *MovieInfo) (int64, error) {
	id, err := db.Insert(movie_info)
	return id, err
}

func GetMovieDirector(movieHtml string) string {

	if movieHtml == "" {
		return ""
	}
	//<a href="/celebrity/1274313/" rel="v:directedBy">邱礼涛</a>
	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])

}
