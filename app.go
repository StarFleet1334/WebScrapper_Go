package main

import (
	"github.com/gocolly/colly"
	"web/utils"
)

func main() {
	c := colly.NewCollector(
		colly.CacheDir("./cache_dir"),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
	)
	utils.Initialization(c, "a[href]", "href")
	utils.LinkToVisit("http://go-colly.org/", c)

}
