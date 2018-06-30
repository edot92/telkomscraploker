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
	urlFb := "https://www.facebook.com/groups/532806206744264"
	crawl.UrlFb = urlFb
	res, err := crawl.RunCrawl()
	if err != nil {
		time.Sleep(3 * time.Second)
		log.Println(err)
		goto restart
	}
	reset := crawl.Parse1(res, urlFb)
	if reset {
		goto restart
	}
	log.Println("SUKSES PROSES SCARPING")
}
