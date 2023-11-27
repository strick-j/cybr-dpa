package dpa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

// Both tests test the validation of the field names of the input struct
// Expect error for both tests
func TestParameterValidation_FieldNameValidation(t *testing.T) {
	var tests = []struct {
		name    string
		input   interface{}
		want    string
		wantErr bool
	}{
		{
			name: "Valid Field Names",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "linux",
				ConnectorType: "AWS",
			},
			wantErr: false,
		},
		{
			name: "Not Valid Field Name 1",
			input: struct{ OS, ConnectorType string }{
				OS:            "linux",
				ConnectorType: "AWS",
			},
			wantErr: true,
		},
		{
			name: "Not Valid Field Name 2",
			input: struct{ ConnectorOS, Type string }{
				ConnectorOS: "linux",
				Type:        "AWS",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parameterValidation(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GenerateScript() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("GenerateScript() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

// All tests test the validation of the field values of the input struct
// Expect error for all tests
func TestParameterValidation_FieldValueValidation(t *testing.T) {
	var tests = []struct {
		name    string
		input   interface{}
		want    string
		wantErr bool
	}{
		{
			name: "Not Valid Field Value 1",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "ubuntu",
				ConnectorType: "AWS",
			},
			wantErr: true,
		},
		{
			name: "Not Valid Field Value 2",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "windows10",
				ConnectorType: "AWS",
			},
			wantErr: true,
		},
		{
			name: "Not Valid Field Value 3",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "windows",
				ConnectorType: "OCI",
			},
			wantErr: true,
		},
		{
			name: "Valid Values",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "windows",
				ConnectorType: "AWS",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parameterValidation(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GenerateScript() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("GenerateScript() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGenerateScript(t *testing.T) {
	var tests = []struct {
		name     string
		input    interface{}
		header   http.ConnState
		sleep    time.Duration
		response string
		want     string
		wantErr  bool
	}{
		{
			name: "Not Valid Field Values",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "ubuntu",
				ConnectorType: "AWS",
			},
			wantErr: true,
		},
		{
			name: "Not Valid Field Name",
			input: struct{ OS, ConnectorType string }{
				OS:            "linux",
				ConnectorType: "AWS",
			},
			wantErr: true,
		},
		{
			name:    "Not Valid Type",
			input:   "test",
			wantErr: true,
		},
		{
			name: "Valid Values",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "windows",
				ConnectorType: "AWS",
			},
			header:   http.StatusOK,
			sleep:    1 * time.Millisecond,
			response: `{"script_url":"https://example.com","bash_cmd":"curl -sSL https://example.com | bash"}`,
			wantErr:  false,
		},
		{
			name: "Status Bad Request",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "windows",
				ConnectorType: "AWS",
			},
			header:   http.StatusBadRequest,
			sleep:    1 * time.Millisecond,
			response: `{"code":"DPA_CONNECTOR_SETUPS_SCRIPT_INVALID_INPUT","message":"Invalid body format","description":"Body should be a dictionary"}`,
			wantErr:  false,
		},
		{
			name: "Timeout",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "windows",
				ConnectorType: "AWS",
			},
			header:   http.StatusOK,
			sleep:    6 * time.Second,
			response: `{"script_url":"https://example.com","bash_cmd":"curl -sSL https://example.com | bash"}`,
			wantErr:  false,
		},
		{
			name: "Status Not Found",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "windows",
				ConnectorType: "AWS",
			},
			header:  http.StatusNotFound,
			sleep:   1 * time.Millisecond,
			wantErr: true,
		},
		{
			name: "Status Forbidden",
			input: struct{ ConnectorOS, ConnectorType string }{
				ConnectorOS:   "windows",
				ConnectorType: "AWS",
			},
			header:  http.StatusForbidden,
			sleep:   1 * time.Millisecond,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Response
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(int(tt.header))
				time.Sleep(tt.sleep)
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
			_, _, err := ns.GenerateScript(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GenerateScript() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("GenerateScript() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
