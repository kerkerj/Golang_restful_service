package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func NewApi() Api {
	return &ApiMartini{}
}

type Api interface {
	Handler() http.Handler
}

type ApiMartini struct{}

func SetupDB() *sql.DB {
	//db, err := sql.Open("mysql", "kerkerj:n4zj6fu4@tcp(back2.pageplates.com:3306)/golang?charset=utf8")
	db, err := sql.Open(DBType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		Username,
		Password,
		Url,
		Port,
		DefaultDB,
	))
	PanicIf(err, nil)

	return db
}

func PanicIf(err error, r render.Render) {
	if err != nil {
		//panic(err.Error())
		//fmt.Println(err.Error())
		r.JSON(404, map[string]interface{}{
			"error":  "404",
			"detail": err,
		})
	}
}

func (a *ApiMartini) Handler() http.Handler {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Map(SetupDB())

	m.Use(func(res http.ResponseWriter, req *http.Request) {
		if req.Header.Get("API-KEY") != "secret123" {
			res.WriteHeader(http.StatusUnauthorized)
		}
	})

	// Set route
	m.Get("/", root)

	m.Group("/page", func(r martini.Router) {
		r.Get("/", getPageAll)
		r.Get("/:id", getPage)
		r.Post("/", binding.Json(Page{}), postPage)
		r.Put("/:id", binding.Json(Page{}), putPage)
		r.Delete("/:id", binding.Json(Page{}), deletePage)
	})

	m.NotFound(notFound)

	return m
}

func main() {
	api := NewApi()
	http.ListenAndServe(":3000", api.Handler())
}
