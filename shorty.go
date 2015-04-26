// shorty.go
package main

import (
	"database/sql"
	"fmt"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
	"math/rand"
	"net/http"
	"time"
)

var siteurl = "4pr.es/"
var db *sql.DB
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var coder Coder

type Coder struct {
	Length    uint
	Shrt, Url string
}

type Web struct {
	Get             bool
	Banner, Content string
	Err             error
}

func init() {
	var err error
	db, err = sql.Open("mysql", "shortener:passwd@/short")
	if err != nil {
		panic(err)
	}
	coder.Length = 6
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout:     "index",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))
	//Landing page
	m.Get("/", func(r render.Render) {
		var land Web
		land.Get = true
		land.Banner = "Get short URL for"
		land.Content = ""
		r.HTML(200, "get", land)
	})
	//Create entry for shortened URL
	m.Post("/", func(req *http.Request, r render.Render) {
		short, err := createUrl(req.FormValue("url"))
		var post Web
		post.Get = false
		if err != nil {
			post.Banner = "Error :("
			post.Content = "Something did not work while trying to shorten URL " +
				req.FormValue("url") + "\n" + err.Error()
			post.Err = err
			r.HTML(500, "error", post)
		}
		post.Content = siteurl + short
		r.HTML(200, "post", post)
	})
	//Redirection to original URL
	m.Get("/:short", func(params martini.Params, w http.ResponseWriter, req *http.Request, r render.Render) {
		err := getUrl(params["short"], w, req)
		if err != nil {
			var e404 Web
			e404.Get = false
			e404.Err = err
			e404.Banner = "404 :("
			e404.Content = "URL " + params["short"] + " not found!"
			r.HTML(404, "error", e404)
		}
	})
	fmt.Println("Shortening URLS on localhost:1337")
	m.RunOnAddr(":1337")
}

func createUrl(input string) (string, error) {
	coder.Url = input
	coder.Shrt = Shorten(coder.Length)
	_, err := db.Exec("INSERT INTO urls VALUES (null, ?, ?, null)", coder.Url, coder.Shrt)
	if err != nil {
		return "", err
	}
	return coder.Shrt, nil
}

func getUrl(short string, w http.ResponseWriter, req *http.Request) error {
	var redir string
	err := db.QueryRow("SELECT url FROM urls WHERE short = ?", short).Scan(&redir)
	if err != nil {
		return err
	}
	http.Redirect(w, req, redir, 301)
	return nil
}

func Shorten(c uint) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, c)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
