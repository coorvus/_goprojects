package main

import (
	"coorvus/web-scraper/engine"
)

var PublicLink = "https://deodatkanabo.vercel.app" // "https://scrape-me.dreamsofcode.io"

func main() {
	// fmt.Println(os.Args)
	engine.Start(PublicLink)
}
