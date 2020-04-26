package crawler

import (
	"fmt"
	"net/http"

	"github.com/zbioe/tor-crawler/loader/tor"
)

type Error struct {
	Err			error
	Operation	string
	Feedback	string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Operation, e.Feedback, e.Err.Error())
}

func Crawl(url string) (Content, error) {
	fmt.Printf("crawling: %s\n", url)
	response, err := request(url)
	if err != nil {
		return Content{}, &Error{err, "crawl", "requests"}
	}
	defer response.Body.Close()
	fmt.Printf("parse: %s\n", url)
	return Parse(response.Body)
}

func request(url string) (*http.Response, error) {
	fmt.Printf("request: %s\n", url)
	client, err := tor.New()
	if err != nil {
		return nil, &Error{err, "tor", "new client"}
	}
	return client.Get(url)
}