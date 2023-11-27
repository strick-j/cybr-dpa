package dpa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func TestGetPublicKey(t *testing.T) {
	var tests = []struct {
		name     string
		input    interface{}
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:    "Invalid string provided",
			input:   "AWS",
			wantErr: true,
		},
		{
			name: "Valid Map provided",
			input: map[string]string{
				"workspaceId":   "123280068473",
				"workspaceType": "AWS",
			},
			header:   http.StatusOK,
			response: `{"publicKey":"sha-rsa AWRKJWKJEXAMPLE13dZYA"}`,
			sleep:    1 * time.Millisecond,
			wantErr:  false,
		},
		{
			name: "Invalid Map Provided",
			input: map[string]string{
				"workspaceId":   "123280068473",
				"workspaceType": "AWS",
				"workspaceName": "Example",
			},
			wantErr: true,
		},
		{
			name: "Timeout",
			input: map[string]string{
				"workspaceId":   "123280068473",
				"workspaceType": "AWS",
			},
			header:   http.StatusOK,
			response: `{"publicKey":"sha-rsa AWRKJWKJEXAMPLE13dZYA"}`,
			sleep:    6 * time.Second,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(tt.sleep)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(int(tt.header))
				w.Write([]byte(tt.response))
			}))
			defer ts.Close()

			// Valid Token
			token := &oauth2.Token{
				AccessToken: "123",
				TokenType:   "bearer",
				Expiry:      time.Now().Add(5 * time.Hour),
			}
			// Valid Service using httptest New Server URL
			ns, _ := NewService(ts.URL, "api", false, token)

			_, _, err := ns.GetPublicKey(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("GetPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetPublicKeyScript(t *testing.T) {
	var tests = []struct {
		name     string
		input    interface{}
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:    "Invalid string provided",
			input:   "AWS",
			wantErr: true,
		},
		{
			name: "Valid Map provided",
			input: map[string]string{
				"workspaceId":   "123280068473",
				"workspaceType": "AWS",
			},
			header:   http.StatusOK,
			response: `{"base64_cmd":"c3254EXAMPLEg=="}`,
			sleep:    1 * time.Millisecond,
			wantErr:  false,
		},
		{
			name: "Invalid Map Provided",
			input: map[string]string{
				"workspaceId":   "123280068473",
				"workspaceType": "AWS",
				"workspaceName": "Example",
			},
			wantErr: true,
		},
		{
			name: "Timeout",
			input: map[string]string{
				"workspaceId":   "123280068473",
				"workspaceType": "AWS",
			},
			header:   http.StatusOK,
			response: `{"base64_cmd":"c3254EXAMPLEg=="}`,
			sleep:    6 * time.Second,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(tt.sleep)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(int(tt.header))
				w.Write([]byte(tt.response))
			}))
			defer ts.Close()

			// Valid Token
			token := &oauth2.Token{
				AccessToken: "123",
				TokenType:   "bearer",
				Expiry:      time.Now().Add(5 * time.Hour),
			}
			// Valid Service using httptest New Server URL
			ns, _ := NewService(ts.URL, "api", false, token)

			_, _, err := ns.GetPublicKeyScript(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetPublicKeyScript() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("GetPublicKeyScript() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
