// shorty.go
package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

var db *sql.DB
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var coder Coder

type Coder struct {
	Length    uint `6`
	Shrt, Url string
}

func init() {
	var err error
	db, err = sql.Open("mysql", "shortener:passwd@/short")
	if err != nil {
		panic(err)
	}
}

func main() {
	m := martini.Classic()
	//Insert
	m.Get("/", func() {
		m.Use(render.Renderer(render.Options{
			Layout:  "index",
			Charset: "UTF-8",
		}))

	})
	//Create entry for shortened URL
	m.Post("/", func() {
		createUrl()
		m.Use(render.Renderer(render.Options{
			Layout:  "post",
			Charset: "UTF-8",
		}))
	})
	//Redirection to original URL
	m.Get("/:id", func(params martini.Params) {
		redirectToUrl(params["id"])
	})

	fmt.Println("Serving on localhost:3000")
	m.Run()
}

func createUrl(input string) (string, int, error) {
	coder.Url = input
	coder.Shorten()
	var id int
	err := db.QueryRow("INSERT INTO urls (url) VALUES ($1) returning id", url).Scan(&id)
	if err != nil {
		return "", 0, err
	}
	return coder.Shrt, id, nil
}

func (c *Coder) Shorten() {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, c.Length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	c.Shrt = string(b)
}
