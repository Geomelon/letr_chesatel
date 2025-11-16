// source: https://ru-brightdata.com/blog/how-tos-ru/web-scraping-go
package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	emptyServer := true
	babydoll := colly.NewCollector()
	babydoll.UserAgent = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	babydoll.OnHTML(".user-server b", func(userLink *colly.HTMLElement) {
		fmt.Println(" ", userLink.Text)
		emptyServer = false
	})

	fmt.Println("Players Online:")
	fmt.Println("\nPOWER TECH")
	babydoll.Visit("https://letragon.ru/servers/pt/")
	if emptyServer {
		fmt.Println("---пусто---")
	}
	emptyServer = true
	fmt.Println("\nHEAVY TECH")
	babydoll.Visit("https://letragon.ru/servers/ht/")
	if emptyServer {
		fmt.Println("---пусто---")
	}
	emptyServer = true
	fmt.Println("\nMAGIC")
	babydoll.Visit("https://letragon.ru/servers/m/")
	if emptyServer {
		fmt.Println("---пусто---")
	}
	emptyServer = true
	fmt.Println("\nMISUSE TECH")
	babydoll.Visit("https://letragon.ru/servers/mt/")
	if emptyServer {
		fmt.Println("---пусто---")
	}
	emptyServer = true
	fmt.Println("\nNEVER CHANGE")
	babydoll.Visit("https://letragon.ru/servers/nc/")
	if emptyServer {
		fmt.Println("---пусто---")
	}
	emptyServer = true

}
