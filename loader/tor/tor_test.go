// +build integration

package tor_test

import (
	"testing"
	"io/ioutil"

	"github.com/iuryfukuda/tor-crawler/loader/tor"
)

func TestGetClient(t *testing.T) {
	client, err := tor.New()
	if err != nil {
		t.Fatalf("unexpected can't get client: %s", err)
	}
	resp, err := client.Get("http://check.torproject.org")
	if err != nil {
		t.Fatalf("Failed to issue GET request: %s\n", err)
	}
	defer resp.Body.Close()
	t.Logf("GET returned: %v\n", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read the body: %s\n", err)
	}
	t.Logf("Body:\n%s", body)
}