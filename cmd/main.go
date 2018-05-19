package main

import (
	"fmt"
	"flag"
	"sync"

	"github.com/iuryfukuda/tor-crawler/crawler"
)

type output struct {
	url string	`json:"url"`
	raw string	`json:"raw"`
}

var (
	concurrence int
	fileOutput  string
	link 		string
)

func init(){
	flag.StringVar(&fileOutput, "output", "output.json", "output file")
	flag.StringVar(&link, "link", "", "input link")
	flag.IntVar(&concurrence, "concurrence", 8, "number of concurrence")
	flag.Parse()
}

func main(){
	links := make(chan string)
	contents := make(chan crawler.Content)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		links <- link
		for link := range links {
			crawler.Crawl(link, contents, &wg)
		}
	}()
	go func() {
		for content := range contents {
			for _, link := range content.Links {
				wg.Add(1)
				links <- link
			}
			fmt.Println(content)
		}
	}()
	wg.Wait()
	fmt.Println("closing channels")
	close(links)
	close(contents)
}