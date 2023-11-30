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

// Valid token used in all tests
var validToken = &oauth2.Token{
	AccessToken: "123",
	TokenType:   "bearer",
	Expiry:      time.Now().Add(5 * time.Hour),
}

func TestListTargetSets(t *testing.T) {
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
			input:   "a0e12345-789e-12ab-abcd-d898f4cc810e",
			wantErr: true,
		},
		{
			name:   "No Query",
			header: http.StatusOK,
			response: `{
				"target_sets": [
					{
						"name": "example.com",
						"provision_format": "<user>-<session-guid>",
						"description": null,
						"enable_certificate_validation": true,
						"secret_type": "PCloudAccount",
						"secret_id": "a0e12345-789e-12ab-abcd-d898f4cc810e",
						"type": "Domain"
					}
				],
				"b64_last_evaluated_key": null
			}`,
			sleep:   1 * time.Millisecond,
			wantErr: false,
		},
		{
			name: "Simple Valid Query",
			input: map[string]string{
				"strongAccountId": "a0e12345-789e-12ab-abcd-d898f4cc810e",
			},
			response: `{
				"target_sets": [
					{
						"name": "example.com",
						"provision_format": "<user>-<session-guid>",
						"description": null,
						"enable_certificate_validation": true,
						"secret_type": "PCloudAccount",
						"secret_id": "a0e12345-789e-12ab-abcd-d898f4cc810e",
						"type": "Domain"
					}
				],
				"b64_last_evaluated_key": null
			}`,
			header:  http.StatusOK,
			sleep:   1 * time.Millisecond,
			wantErr: false,
		},
		{
			name: "Valid query timeout",
			input: map[string]string{
				"strongAccountId": "a0e12345-789e-12ab-abcd-d898f4cc810e",
			},
			response: `{"target_sets": [],"b64_last_evaluated_key": null}`,
			header:   http.StatusOK,
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

			// Valid Service using httptest New Server URL
			ns, _ := NewService(ts.URL, "api", false, validToken)

			_, _, err := ns.ListTargetSets(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ListTargetSets() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("ListTargetSets() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestAddTargetSet(t *testing.T) {
	var tests = []struct {
		name     string
		input    interface{}
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:    "Invalid Input",
			input:   "String input not slice",
			wantErr: true,
		},
		{
			name: "Invalid Struct Bad Request",
			input: struct {
				isMfaCachingEnabled bool
			}{
				isMfaCachingEnabled: true,
			},
			response: `{
				"tenant_id": "28515795-2bad-4468-8eb7-026a68520adf",
				"exception": "Validation error occurred while doing target set bulk operation [1 validation error for TargetSetsBulkMappingDto\ntarget_sets_mapping\n  field required (type=value_error.missing)]"
			}`,
			header:  http.StatusBadRequest,
			sleep:   1 * time.Millisecond,
			wantErr: false,
		},
		{
			name:     "Add Target Set Timeout",
			input:    []string{},
			response: `{"results":[]}`,
			header:   http.StatusOK,
			sleep:    6 * time.Second,
			wantErr:  true,
		},
		{
			name: "Add Target Set Unhandled Error",
			input: struct {
				isMfaCachingEnabled bool
			}{
				isMfaCachingEnabled: true,
			},
			response: `{"results":[]}`,
			header:   http.StatusTooManyRequests,
			sleep:    1 * time.Millisecond,
			wantErr:  true,
		},
		{
			name: "Add Target Set Success",
			input: types.TargetSetMapping{
				StrongAccountID: "1239-809e-45ab-abef-d424244cc810e",
				TargetSets: []types.TargetSets{
					{
						Name:                        "test.com",
						ProvisionFormat:             "<user>-<session-guid>",
						EnableCertificateValidation: true,
						SecretType:                  "PCloudAccount",
						SecretID:                    "1239-809e-45ab-abef-d424244cc810e",
						Type:                        "Domain",
					},
				},
			},
			response: `{
				"results": [
					{
						"strong_account_id": "1239-809e-45ab-abef-d424244cc810e",
						"target_set_name": "test.com",
						"success": true
					}
				]
			}`,
			header:  http.StatusOK,
			sleep:   1 * time.Millisecond,
			wantErr: false,
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

			// Valid Service using httptest New Server URL
			ns, _ := NewService(ts.URL, "api", false, validToken)

			_, _, err := ns.AddTargetSet(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("AddTargetSet() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("AddTargetSet() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestDeleteTargetSet(t *testing.T) {
	var tests = []struct {
		name     string
		input    interface{}
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:    "Invalid Input",
			input:   "String input not slice",
			wantErr: true,
		},
		{
			name:     "Empty Slice",
			input:    []string{},
			response: `{"results":[]}`,
			header:   http.StatusOK,
			sleep:    1 * time.Millisecond,
			wantErr:  false,
		},
		{
			name:     "Delete Target Set Timeout",
			input:    []string{},
			response: `{"results":[]}`,
			header:   http.StatusOK,
			sleep:    6 * time.Second,
			wantErr:  true,
		},
		{
			name:     "Valid Deletion",
			input:    []string{"example.com"},
			response: `{"results":[{"strong_account_id": null,"target_set_name":"example.com","success": true}]}`,
			header:   http.StatusMultiStatus,
			sleep:    1 * time.Millisecond,
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

			// Valid Service using httptest New Server URL
			ns, _ := NewService(ts.URL, "api", false, validToken)

			_, _, err := ns.DeleteTargetSet(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("DeleteTargetSet() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("DeleteTargetSet() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}

}
