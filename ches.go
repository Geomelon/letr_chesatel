// source: https://ru-brightdata.com/blog/how-tos-ru/web-scraping-go
package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	babydoll := colly.NewCollector()
	babydoll.UserAgent = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	// var users []string
	babydoll.OnHTML(".user-server b", func(userLink *colly.HTMLElement) {
		fmt.Println(userLink.Text)
	})

	fmt.Println("Players Online:")
	babydoll.Visit("https://letragon.ru/servers/mt/")
}
