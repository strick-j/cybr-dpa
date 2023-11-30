package dpa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
	"golang.org/x/oauth2"
)

func TestListSettings_BadRequest(t *testing.T) {
	want := "error message"
	// Mock Response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code":"DPA_INVALID_VALUE","message":"error message","description":"error description"}`))
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

	_, got, err := ns.ListSettings(context.Background())
	if got.Message != want {
		t.Errorf("got %v, wanted %v", got.Message, want)
	}
	if err != nil {
		t.Errorf("ListSettings() error = %v, wantNoErr", err)
	}
}

func TestListSettings_Timeout(t *testing.T) {
	// Mock Response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(6 * time.Second)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code":"DPA_INVALID_VALUE","message":"error message","description":"error description"}`))
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

	_, _, err := ns.ListSettings(context.Background())
	if err == nil {
		t.Errorf("ListSettings() got no error = %v, wantErr", err)
	}
}

func TestListSettings(t *testing.T) {
	want := 3600
	// Mock Response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"mfaCaching":{"isMfaCachingEnabled":true,"keyExpirationTimeSec":3600},"sshCommandAudit":{"isCommandParsingForAuditEnabled":true,"shellPromptForAudit":"(.*)[>#\\$]$"}}`))
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

	got, _, err := ns.ListSettings(context.Background())
	if got.MfaCaching.KeyExpirationTimeSec != want {
		t.Errorf("got %v, wanted %v", got.MfaCaching.KeyExpirationTimeSec, want)
	}
	if err != nil {
		t.Errorf("ListSettings() error = %v, wantNoErr", err)
	}
}

func TestListSettingsFeature(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:     "Invalid string provided",
			input:    "mfacaching",
			header:   http.StatusBadRequest,
			response: `{"code":"400","message":"Bad Request","description":"value is not a valid enumeration member; permitted: 'MFA_CACHING', 'STANDING_ACCESS', 'SSH_COMMAND_AUDIT', 'RDP_FILE_TRANSFER', 'CERTIFICATE_VALIDATION' (field: featureName)"}`,
			sleep:    1 * time.Millisecond,
			wantErr:  false,
		},
		{
			name:     "Valid string provided",
			input:    "MFA_CACHING",
			header:   http.StatusOK,
			response: `{"feature_name":"MFA_CACHING","feature_conf":{"is_mfa_caching_enabled":true,"key_expiration_time_sec":3600}}`,
			sleep:    1 * time.Millisecond,
			wantErr:  false,
		},
		{
			name:     "Invalid Token",
			input:    "MFA_CACHING",
			header:   http.StatusUnauthorized,
			response: `{"code":"DPA_AUTHENTICATION_TOKEN_VALIDATION_FAILED","message":"Authentication failed. If the issue persists, please contact your system administrator.","description":"Authentication token validation failed"}`,
			sleep:    1 * time.Millisecond,
			wantErr:  false,
		},
		{
			name:     "Timeout",
			input:    "MFA_CACHING",
			header:   http.StatusOK,
			response: `{"feature_name":"MFA_CACHING","feature_conf":{"is_mfa_caching_enabled":true,"key_expiration_time_sec":3600}}`,
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

			_, _, err := ns.ListSettingsFeature(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ListSettingsFeature() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("ListSettingsFeature() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestUpdateSettings(t *testing.T) {
	var tests = []struct {
		name     string
		input    interface{}
		response string
		header   http.ConnState
		sleep    time.Duration
		wantErr  bool
	}{
		{
			name:    "Invalid input type",
			input:   "String is not a valid input type",
			sleep:   1 * time.Millisecond,
			wantErr: true,
		},
		{
			name: "Timeout",
			input: struct {
				isMfaCachingEnabled  bool
				keyExpirationTimeSec int
			}{
				isMfaCachingEnabled:  true,
				keyExpirationTimeSec: 3600,
			},
			sleep:    6 * time.Second,
			header:   http.StatusOK,
			response: `{"feature_name":"MFA_CACHING","feature_conf":{"is_mfa_caching_enabled":true,"key_expiration_time_sec":3600}}`,
			wantErr:  true,
		},
		{
			name: "Valid Input",
			input: struct {
				isMfaCachingEnabled  bool
				keyExpirationTimeSec int
			}{
				isMfaCachingEnabled:  true,
				keyExpirationTimeSec: 3600,
			},
			header:   http.StatusOK,
			sleep:    1 * time.Millisecond,
			response: `{"feature_name":"MFA_CACHING","feature_conf":{"is_mfa_caching_enabled":true,"key_expiration_time_sec":3600}}`,
			wantErr:  false,
		},
		{
			name: "Invalid Input",
			input: types.Settings{
				MfaCaching: types.MfaCaching{
					IsMfaCachingEnabled:  true,
					KeyExpirationTimeSec: 3600,
				},
			},
			sleep:    1 * time.Millisecond,
			header:   http.StatusBadRequest,
			response: `{"code":"400","message":"Bad Request","description":"extra fields not permitted (field: mfaCachingConfiguration)"}`,
			wantErr:  false,
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

			_, _, err := ns.UpdateSettings(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("UpdateSettings() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("UpdateSettings() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
