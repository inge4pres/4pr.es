// shorty.go
package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
	"time"
)

var db *sql.DB
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var coder Coder

type Coder struct {
	Length    uint
	Shrt, Url string
}

func init() {
	var err error
	db, err = sql.Open("mysql", "shortener:passwd@/short")
	if err != nil {
		panic(err)
	}
	coder.Length = 8
}

func main() {
	m := martini.Classic()
	//Landing page
	m.Use(render.Renderer(render.Options{
		Layout:     "index",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))
	m.Get("/", func(r render.Render) {
		r.HTML(200, "", "")
	})
	//Create entry for shortened URL
	m.Post("/", func(r *http.Request) string {
		short, err := createUrl(r.FormValue("url"))
		if err != nil {
			return "Error shortening URL " + r.FormValue("url") + "\n" + err.Error()
		}
		return "Short URL: " + short + " \n"
	})
	//Redirection to original URL
	m.Get("/:short", func(params martini.Params, w http.ResponseWriter, r *http.Request) string {
		redir, err := getUrl(params["short"])
		if err != nil {
			return "404 :(\nURL " + params["short"] + " not found\n"
		}
		http.Redirect(w, r, redir, 301)
		return "Redirect OK :D"
	})

	fmt.Println("Shortening URLS on localhost:1337")
	m.RunOnAddr(":1337")
}

func createUrl(input string) (string, error) {
	coder.Url = input
	coder.Shrt = Shorten(coder.Length)
	/*_ = db.QueryRow("SELECT id from ursl where short = $1", coder.Shrt).Scan(&id)
		fmt.Printf("Short URL %s is alredy gone!", coder.Shrt)
	} */
	_, err := db.Exec("INSERT INTO urls VALUES (null, ?, ?, null)", coder.Url, coder.Shrt)
	if err != nil {
		return "", err
	}
	return coder.Shrt, nil
}

func getUrl(short string) (string, error) {
	var redir string
	err := db.QueryRow("SELECT url FROM urls WHERE short = ?", short).Scan(&redir)
	if err != nil {
		return "", err
	}
	return redir, nil
}

func Shorten(c uint) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, c)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
