package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

func randomAgent() string {
	var (
		lowchars = "abcdefghijklmnopqrstuvwxyz"
		upchars  = strings.ToUpper(lowchars)
		chars    = lowchars + upchars
	)
	b := make([]byte, rand.Intn(5)+20)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func main() {
	const FRurl = "https://www.reddit.com/r/FashionReps/search?q=flair_name%3A\"GIFTBAG\"&restrict_sr=&t=hour1&sort=new"
	color.Set(color.FgCyan)
	logo := fmt.Sprintf("" +
		" _______ __  ___ __   _______             \n" +
		"|   _   |__.'  _|  |_|   _   .-----.-----.\n" +
		"|.  |___|  |   _|   _|.  1   |  _  |  _  |\n" +
		"|.  |   |__|__| |____|.  _   |_____|___  |\n" +
		"|:  1   |            |:  1    \\    |_____|\n" +
		"|::.. . |            |::.. .  /           \n" +
		"`-------'            `-------'            \n")
	fmt.Println(logo)
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randomAgent())
		fmt.Println("[STATUS]: fetching giftbags...")
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("[ERROR]:", err)
	})
	var links []string
	c.OnHTML("a[data-click-id]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.Contains(link, "/sharepack/") {
			return
		}
		links = append(links, link)
	})
	c.Visit(FRurl)
	for _, link := range links {
		if len(link) > 5 {
			fmt.Println("[GIFTBAG]:", link)
		}
	}
}
