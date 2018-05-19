// define the loader to get a client with proxy
package loader

import "net/http"

type Loader interface {
	New(endpoint string) (http.Client, error)
}