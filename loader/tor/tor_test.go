// +build integration

package tor_test

import (
	"testing"
	"io/ioutil"

	"github.com/iuryfukuda/tor-crawler/loader/tor"
)

const (
	validEndpoint = "socks5://127.0.0.1:9050"
	invalidEndpoint = "teste://127.0.0.1:9050"
)

func TestGetClient(t *testing.T) {
	client, err := tor.New(validEndpoint)
	if err != nil {
		t.Fatalf("unexpected can't get client: %T", err)
	}
	resp, err := client.Get("http://check.torproject.org")
	if err != nil {
		t.Fatalf("Failed to issue GET request: %T\n", err)
	}
	defer resp.Body.Close()
	t.Logf("GET returned: %v\n", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read the body: %T\n", err)
	}
	t.Logf("Body:\n%s", body)
}

func TestCantGetClient(t *testing.T) {
	_, err := tor.New(invalidEndpoint)
	if err == nil {
		t.Fatal("unexpected nil in err")
	}
}