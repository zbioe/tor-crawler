package main

import (
	"log"
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
	url 		string
	links		chan string
	contents	chan crawler.Content
	wg			sync.WaitGroup
)

func init(){
	flag.StringVar(&fileOutput, "output", "output.json", "output file")
	flag.StringVar(&url, "url", "", "input url")
	flag.IntVar(&concurrence, "concurrence", 8, "number of concurrence")
	flag.Parse()
	links = make(chan string)
	contents = make(chan crawler.Content)
}

func workerCrawler() {
	for link := range links {
		go func() {
			defer wg.Done()
			content, error := crawler.Crawl(link)
			if error != nil {
				log.Printf("Problem with Crawl of %s:%s:", link, error)
				return
			}
			contents <- content
		}()
	}
}


func poolerLinks() {
	for content := range contents {
		for _, link := range content.Links {
			wg.Add(1)
			links <- link
		}
	}
}

func main(){
	wg.Add(1)
	links <- url
	for i := 0; i < concurrence; i++ {
		go workerCrawler()
		go poolerLinks()
	}
	wg.Wait()
	close(links)
	close(contents)
}