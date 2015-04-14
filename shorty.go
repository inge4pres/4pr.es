// shorty.go
package main

import (
	"database/sql"
	"math/rand"
	"time"
	_ "fmt"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var mixtape = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Shorten(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(mixtape))]
    }
    return string(b)
}

type Coder struct {
	Length	uint
	Shrt		Shorten(Length)
	Url		string
}

func init() {
	var err error
	db, err = sql.Open("mysql", "shortener:passwd@/short")
}

func main() {
	m := martini.Classic()
	m.Get("/", func() {
		Printf("Welcome to s4pres ULR Shortener!")
	})
	fmt.Prinln("Service ing localhost:3000")
	m.Run()
}
