package tor

import (
	"fmt"
	"net/url"
	"net/http"
	"golang.org/x/net/proxy"
)

const (
	ENDPOINT = "socks5://127.0.0.1:9050"
	EDIALERPROXY = "Failed to create new dialer"
	EPARSEPROXY  = "Failed to obtain new proxy"
)

func errFmt(s string, err error) error {
	return fmt.Errorf("%s: %s", s, err)
}

func New() (*http.Client, error) {
	u, err := url.Parse(ENDPOINT)
	if err != nil {
		return &http.Client{}, errFmt(EPARSEPROXY, err)
	}

	dialer, err := proxy.FromURL(u, proxy.Direct)
	if err != nil {
		return &http.Client{}, errFmt(EDIALERPROXY, err)
	}
	
	return &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}, nil
}