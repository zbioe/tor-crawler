package tor

import (
	"fmt"
	"net/http"
	"golang.org/x/net/proxy"
)

const (
	ENDPOINT = "127.0.0.1:9050"
	EDIALERPROXY = "Failed to create new dialer"
)

func errFmt(s string, err error) error {
	return fmt.Errorf("%s: %s", s, err)
}

func New() (*http.Client, error) {
	dialer, err := proxy.SOCKS5("tcp", ENDPOINT, nil, proxy.Direct)
	if err != nil {
		return &http.Client{}, errFmt(EDIALERPROXY, err)
	}

	return &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}, nil
}