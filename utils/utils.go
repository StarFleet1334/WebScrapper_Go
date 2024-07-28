package utils

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"sync"
)

var visitedURLs = struct {
	sync.RWMutex
	m map[string]bool
}{m: make(map[string]bool)}

func Initialization(c *colly.Collector, selectorPattern string, attributeName string) {
	c.OnHTML(selectorPattern, func(e *colly.HTMLElement) {
		link := e.Attr(attributeName)
		if link == "" {
			fmt.Println("No link found for attribute:", attributeName)
			return
		}
		visitedURLs.RLock()
		_, visited := visitedURLs.m[link]
		visitedURLs.RUnlock()

		if !visited {
			visitedURLs.Lock()
			visitedURLs.m[link] = true
			visitedURLs.Unlock()
			err := e.Request.Visit(link)
			if err != nil {
				fmt.Println("There was an error while fetching or finding attributed")
				fmt.Println(err)
				return
			}
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
}

func LinkToVisit(linkName string, c *colly.Collector) {
	cacheDir := "cache_dir"
	cacheFilePath := GetCacheFilePath(cacheDir, linkName)

	// Debugging print
	fmt.Println("Checking cache for:", linkName)

	if data, err := ReadCache(cacheFilePath); err == nil {
		fmt.Printf("Loading from cache: %s\n", linkName)
		// Use the cached data
		fmt.Println(string(data))
		return
	} else {
		fmt.Println("Cache miss for:", linkName)
	}

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Caching response for:", r.Request.URL.String())
		err1 := os.MkdirAll(cacheDir, os.ModePerm)
		if err1 != nil {
			return
		}
		err := WriteCache(cacheFilePath, r.Body)
		if err != nil {
			fmt.Printf("Error saving to cache: %s\n", err)
		}
	})

	err := c.Visit(linkName)
	if err != nil {
		fmt.Printf("Some Error occurred: %s", err)
		return
	}
}
