package crawler

import (
	"io"
	"fmt"
	"bytes"
	"regexp"
	"golang.org/x/net/html"
)

type Content struct {
	Raw		string
	Links	[]string
}

var (
	b bytes.Buffer
	reLink = regexp.MustCompile(`(https?)://?(?:www)?(\S*?\.onion)\b`)
)

func Parse(r io.Reader) (Content, error) {
	fmt.Println("parsing")
	raw, err := toRaw(r)
	if err != nil {
		return Content{}, &Error{err, "parse", "to raw"}
	}
	return Content{Raw: raw, Links: getLinks(raw)}, nil
}

func toRaw(r io.Reader) (string, error) {
	fmt.Println("to raw")
	node, err := html.Parse(r)
	if err != nil {
		return "", &Error{err, "reader", "to node"}
	}
	
	err = html.Render(&b, node)
	if err != nil {
		return "", &Error{err, "html", "render node"}
	}
	return b.String(), nil
}

func getLinks(s string) []string {
	fmt.Println("get links")
	return reLink.FindAllString(s, -1)
}