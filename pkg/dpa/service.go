package dpa

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type Service struct {
	client *Client
}

func NewService(clientURL, clientApiEndpoint string, verbose bool, authToken *oauth2.Token) (*Service, error) {
	// Validate Bearer Token was provided
	tokenType := authToken.Type()
	if tokenType != "Bearer" {
		return nil, fmt.Errorf("dpa: invalid token type provided %s, expected type is bearer token", tokenType)
	}

	// Validate Bearer Token is still valid
	if !authToken.Valid() {
		return nil, fmt.Errorf("dpa: token is invalid")
	}

	tr := &Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: authToken.AccessToken,
			TokenType:   authToken.TokenType,
			Expiry:      authToken.Expiry,
		}),
	}

	return &Service{
		client: NewClient(
			&http.Client{Transport: tr},
			Options{
				ApiURL:  fmt.Sprintf("%s/%s", clientURL, clientApiEndpoint),
				Verbose: verbose,
			},
		),
	}, nil
}
