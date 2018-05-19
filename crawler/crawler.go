package crawler

import (
	"fmt"
	"sync"
	"net/http"

	"github.com/iuryfukuda/tor-crawler/loader/tor"
)

type (
	Error struct {
		Err			error
		Operation	string
		Feedback	string
	}
)

func (e *Error) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Operation, e.Feedback, e.Err.Error())
}

func Crawl(link string, contents chan<- Content, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Crawling: %s\n", link)
	content, err := crawl(link)
	if err != nil {
		fmt.Printf("Unexpected: %s", err)
	}
	contents <- content
}

func crawl(url string) (Content, error) {
	fmt.Printf("crawling: %s\n", url)
	response, err := request(url)
	if err != nil {
		return Content{}, &Error{err, "crawl", "requests"}
	}
	defer response.Body.Close()
	return Parse(response.Body)
}

func request(url string) (*http.Response, error) {
	fmt.Printf("request: %s", url)
	client, err := tor.New()
	if err != nil {
		return nil, &Error{err, "tor", "new client"}
	}
	return client.Get(url)
}