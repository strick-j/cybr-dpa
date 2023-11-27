package dpa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

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

			// Valid Token
			token := &oauth2.Token{
				AccessToken: "123",
				TokenType:   "bearer",
				Expiry:      time.Now().Add(5 * time.Hour),
			}
			// Valid Service using httptest New Server URL
			ns, _ := NewService(ts.URL, "api", false, token)

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

}

func TestDeleteTargetSet(t *testing.T) {

}
