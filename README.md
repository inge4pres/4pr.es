# __4pr.es__
## A URL shortener in Go
[4pres](https://4pr.es/) is yet another URL shortener, it will give you a short link for a long URL.

It is written in 120 lines of [Go](http://golang.org/), it uses the powerful [go-martini](https://github.com/go-martini) web framework and the beautiful [basscss](http://www.basscss.com/) as CSS style.

Short links have no expiration, for now :D, you can get short URL via any browser or other HTTP client.

Visit [4pres](https://4pr.es/) with your web browser

  __or__
  
curl --data "url=http://YOURLONGURLHERE/SOME%20CONTEXT/FILE.EXT" http://4pr.es/ | grep ">4pr.es/"

#### Install your own

    git clone https://github.com/inge4pres/4pr.es.git
    go get github.com/go-sql-driver/mysql
    go get github.com/go-martini/martini
    go get github.com/martini-contrib/render
    mysql -u root -p < db.sql 
    go run shorty.go -p 12345 -d yourdomain.com [-s]
