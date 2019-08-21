package main

import (
	"fmt"
	"math/rand"
	"os"
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
	b := make([]byte, rand.Intn(20)+5)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func main() {
	const FashionReps = "https://www.reddit.com/r/FashionReps/search?q=flair_name%3A\"GIFTBAG\"&restrict_sr=1&sort=new&t=hour"
	var links []string
	var quotes = []string{
		"inbound into customs...", "seizing your haul...", "legit checking...", "calling you out...",
		"checking washtags...", "declaring 19$ for 20kg...", "taking QC pics...", "GP'ing the peppa pig flip flops...",
		"tracking the package...", "getting rid of the shoebox...", "using simple packaging...", "uploading with QCsuite...",
		"lurking on goat.pet/fr/sb", "saving money for the haul...", "flexin on 'em...", "making ur broke ass look like u got money..."}
	cyan := color.New(color.FgCyan)
	red := color.New(color.FgRed)
	color.Set(color.FgCyan)
	logo := fmt.Sprintf("" +
		" _______ __  ___ __   _______             \n" +
		"|   _   |__.'  _|  |_|   _   .-----.-----.\n" +
		"|.  |___|  |   _|   _|.  1   |  _  |  _  |\n" +
		"|.  |   |__|__| |____|.  _   |_____|___  |\n" +
		"|:  1   |            |:  1    \\    |_____|\n" +
		"|::.. . |            |::.. .  /           \n" +
		"`-------'            `-------'            \n")
	print("\033[H\033[2J")
	fmt.Println(logo)
	color.Unset()
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randomAgent())
		cyan.Printf("%s ", "[STATUS]:")
		fmt.Printf(quotes[rand.Intn(len(quotes))] + "\n")
	})
	c.OnError(func(_ *colly.Response, err error) {
		red.Printf("%s ", "[ERROR]:")
		fmt.Printf("%s\n\n", "your connection is not 1:1!")
		os.Exit(1)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.Contains(link, "/sharepack/") {
			return
		}
		if len(links) > 0 {
			if links[len(links)-1] != link {
				links = append(links, link)
			}
		} else {
			links = append(links, link)
		}
	})
	c.Visit(FashionReps)
	if len(links) > 1 {
		cyan.Printf("%s ", "[STATUS]:")
		fmt.Printf("found %d new giftbags!\n", len(links))
	} else {
		if len(links) == 0 {
			cyan.Printf("%s ", "[STATUS]:")
			fmt.Printf("%s", "sry fam, no giftbags right now\n")
		} else {
			cyan.Printf("%s ", "[STATUS]:")
			fmt.Printf("found %d new giftbag!\n", len(links))
		}
	}
	if len(links) != 0 {
		for _, link := range links {
			if len(link) > 3 {
				cyan.Printf("%s", "[LINK]: ")
				fmt.Printf("%s\n", link)
			}
		}
	}
	fmt.Println("")
	os.Exit(0)
}
