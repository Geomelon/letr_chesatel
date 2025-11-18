// source: https://ru-brightdata.com/blog/how-tos-ru/web-scraping-go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var EmptyFlag = true

func main() {
	for {
		clearScreen()
		babydoll := colly.NewCollector()
		babydoll.UserAgent = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

		babydoll.OnHTML(".user-server b", func(userLink *colly.HTMLElement) {
			fmt.Println(" ", userLink.Text)
			EmptyFlag = false
		})

		fmt.Println("Players Online:")

		scrubServer("POWER TECH", babydoll)
		scrubServer("HEAVY TECH", babydoll)
		scrubServer("MAGIC", babydoll)
		scrubServer("MISUSE TECH", babydoll)
		scrubServer("NEVER CHANGE", babydoll)

		time.Sleep(60 * time.Second)
	}
}

func scrubServer(serverName string, babydoll *colly.Collector) {
	EmptyFlag = true
	fmt.Println("\n" + serverName)
	words := strings.Fields(serverName)
	var firstLetters string
	for _, word := range words {
		firstLetters += strings.ToLower(string(word[0]))
	}
	babydoll.Visit("https://letragon.ru/servers/" + firstLetters)
	if EmptyFlag {
		fmt.Println("---пусто---")
	}
}

func clearScreen() {
	command := exec.Command("cmd", "/c", "cls")
	command.Stdout = os.Stdout
	command.Run()
}
