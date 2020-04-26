package crawler_test

import (
    "bufio"
	"io"
	"os"
	"testing"

	"github.com/zbioe/tor-crawler/crawler"
)

func fileToReader(t *testing.T, rpath string) io.Reader {
	f, err := os.Open("test_files/" + rpath)
	if err != nil {
		t.Fatal("fileToReader os open:", err)
	}
	return bufio.NewReader(f)
}

func TestParseSuccess(t *testing.T) {
	r := fileToReader(t, "dreammarket.html")
	content, err := crawler.Parse(r)
	if err != nil || len(content.Links) <= 0 {
		t.Fatal("parse fails with no nil err")
	}
	t.Log(content)
}
