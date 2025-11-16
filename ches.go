// source: https://ru-brightdata.com/blog/how-tos-ru/web-scraping-go
package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Check boobies")
	babydoll := colly.NewCollector()
	babydoll.Visit("https://letragon.ru/servers/mt/")

}
