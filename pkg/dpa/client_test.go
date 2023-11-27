package dpa

import "net/http"

type TestOptions struct {
	ApiURL  string
	Verbose bool
}

type TestClient struct {
	httpClient *http.Client
	options    *Options
}
