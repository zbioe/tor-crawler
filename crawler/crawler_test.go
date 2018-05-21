// +build integration

package crawler_test

import (
	"testing"

	"github.com/iuryfukuda/tor-crawler/crawler"
)

func TestGetContent(t *testing.T) {
	_, err := crawler.Crawl("http://7ep7acrkunzdcw3l.onion")
	if err != nil {
		t.Fatalf("unexpected can't get content: %s", err)
	}
}