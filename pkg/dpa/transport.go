package dpa

import (
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

// Transport is an http.RoundTripper that makes requests,
// wrapping a base RoundTripper and adding an Authorization header
// with a bearer token.
//
// Transport is a low-level mechanism.
type Transport struct {
	Source oauth2.TokenSource
	Base   http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqBodyClosed := false
	if req.Body != nil {
		defer func() {
			if !reqBodyClosed {
				req.Body.Close()
			}
		}()
	}

	if t.Source == nil {
		return nil, errors.New("dpa: Transport's Source is nil")
	}

	token, err := t.Source.Token()
	if err != nil {
		return nil, err
	}

	r := req.Clone(req.Context()) // per RoundTripper contract
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	token.SetAuthHeader(r)

	// req.Body is assumed to be closed by the base RoundTripper.
	reqBodyClosed = true

	return t.base().RoundTrip(r)
}

func (t *Transport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}
