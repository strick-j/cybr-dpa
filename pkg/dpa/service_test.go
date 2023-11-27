package dpa

import (
	"log"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func TestNewServiceValidToken(t *testing.T) {
	clientUrl := "https://testing.cyberark.cloud"
	clientApiEndpoint := "api"
	verbose := false
	token := &oauth2.Token{
		AccessToken: "123",
		TokenType:   "bearer",
		Expiry:      time.Now().Add(5 * time.Hour),
	}

	s, err := NewService(clientUrl, clientApiEndpoint, verbose, token)
	if err != nil {
		t.Errorf("got invalid with valid options; want valid: %s", err)
	}

	log.Printf("%s", s.client.httpClient.Transport)
}

func TestNewServiceExpiredToken(t *testing.T) {
	clientUrl := "https://testing.cyberark.cloud"
	clientApiEndpoint := "api"
	verbose := false
	token := &oauth2.Token{
		AccessToken: "123",
		TokenType:   "bearer",
		Expiry:      time.Now().Add(-5 * time.Hour),
	}

	_, err := NewService(clientUrl, clientApiEndpoint, verbose, token)
	if err == nil {
		t.Error("got valid with expired token; want invalid")
	}
}

func TestNewServiceInvalidTokenType(t *testing.T) {
	clientUrl := "https://testing.cyberark.cloud"
	clientApiEndpoint := "api"
	verbose := false
	token := &oauth2.Token{
		AccessToken: "123",
		TokenType:   "basic",
		Expiry:      time.Now().Add(5 * time.Hour),
	}

	_, err := NewService(clientUrl, clientApiEndpoint, verbose, token)
	if err == nil {
		t.Error("got valid with basic token; want invalid")
	}
}

func TestNewServiceEmptyToken(t *testing.T) {
	clientUrl := "https://testing.cyberark.cloud"
	clientApiEndpoint := "api"
	verbose := false
	token := &oauth2.Token{}

	_, err := NewService(clientUrl, clientApiEndpoint, verbose, token)
	if err == nil {
		t.Error("got valid with empty token; want invalid")
	}
}

func TestNewServiceClientUrl(t *testing.T) {
	want := "https://testing.cyberark.cloud/api"
	clientUrl := "https://testing.cyberark.cloud"
	clientApiEndpoint := "api"
	verbose := false
	token := &oauth2.Token{
		AccessToken: "123",
		TokenType:   "bearer",
		Expiry:      time.Now().Add(5 * time.Hour),
	}

	got, err := NewService(clientUrl, clientApiEndpoint, verbose, token)
	if err != nil {
		t.Errorf("got error; wanted no error: %err", err)
	}

	if got.client.options.ApiURL != want {
		t.Errorf("got %s, wanted %s", got.client.options.ApiURL, want)
	}
}
