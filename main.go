package main

import (
	"log"
	"sample/crawl"
	"sample/db"
	"time"
)

func main() {
	db.InitDB()
restart:
	res, err := crawl.RunCrawl()
	if err != nil {
		time.Sleep(3 * time.Second)
		log.Println(err)
		goto restart
	}
	crawl.Parse1(res)
}
