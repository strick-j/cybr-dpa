package dpa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestListPolicies(t *testing.T) {
	var tests = []struct {
		name     string
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:    "Invalid Timeout",
			sleep:   6 * time.Second,
			wantErr: true,
		},
		{
			name:    "Invalid Too Many Requests",
			header:  http.StatusTooManyRequests,
			sleep:   1 * time.Millisecond,
			wantErr: true,
		},
		{
			name: "Valid Response",
			response: `{
				"items": [
					{
						"policyId": "c12f982a-ab1a-12ab-1a31-f221aa31836b",
						"status": "Enabled",
						"policyName": "Example Policy 2",
						"description": "",
						"updatedOn": "2023-11-20T14:16:42.161149",
						"ruleNames": [
							"AzureSSHAccess"
						],
						"platforms": [
							"Azure"
						]
					},
					{
						"policyId": "c12f982a-ab1a-12ab-1a31-f221aa31836b",
						"status": "Enabled",
						"policyName": "Example Policy 2",
						"description": "",
						"updatedOn": "2023-11-07T21:45:12.303319",
						"ruleNames": [
							"EL EC2",
							"Ubuntu EC2"
						],
						"platforms": [
							"AWS"
							]
						}
					],
					"totalCount": 2
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

			_, _, err := ns.ListPolicies(context.Background())
			if tt.wantErr {
				if err == nil {
					t.Errorf("ListPolicies() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("ListPolicies() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestGetPolicy(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:    "Invalid Timeout",
			input:   "c12f982a-ab1a-12ab-1a31-f221aa31836a",
			header:  http.StatusOK,
			sleep:   6 * time.Second,
			wantErr: true,
		},
		{
			name:   "Invalid Status Not Found",
			input:  "c12f982a-ab1a-12ab-1a31-f221aa31836a",
			header: http.StatusNotFound,
			response: `{
				"code": "DPA_CRUD_ACTION_FAILED",
				"message": "Unable to update an Authorization Policy.",
				"description": "Unable to update an Authorization Policy. Error(s): Policy with id 01a4f891-1591-4acb-ae3f-f27e56d45491 was not found",
				"doc": null,
				"steps": null
			}`,
			sleep:   1 * time.Millisecond,
			wantErr: false,
		},
		{
			name:  "Valid Response",
			input: "01a4f891-1591-4acb-ae3f-f27e56d45499",
			response: `{
				"policyId": "01a4f891-1591-4acb-ae3f-f27e56d45499",
				"policyName": "Production System Access",
				"status": "Draft",
				"description": "",
				"providersData": {
					"OnPrem": {
						"fqdnRulesConjunction": "OR",
						"fqdnRules": [
							{
								"operator": "CONTAINS",
								"computernamePattern": "prod",
								"domain": "example.local"
							},
							{
								"operator": "CONTAINS",
								"computernamePattern": "prd",
								"domain": "example.local"
							}
						],
						"logicalNames": null
					}
				},
				"startDate": null,
				"endDate": null,
				"userAccessRules": [
					{
						"ruleName": "StorageTower",
						"userData": {
							"roles": [
								{
									"name": "StorageTower",
									"source": null
								}
							],
							"groups": [],
							"users": []
						},
						"connectionInformation": {
							"connectAs": {
								"OnPrem": {
									"rdp": {
										"localEphemeralUser": {
											"assignGroups": [
												"Remote Desktop Users",
												"Administrators"
											]
										}
									}
								}
							},
							"grantAccess": 2,
							"idleTime": 10,
							"daysOfWeek": [
								"Fri",
								"Mon",
								"Sat",
								"Sun",
								"Thu",
								"Tue",
								"Wed"
							],
							"fullDays": false,
							"hoursFrom": "08:00",
							"hoursTo": "18:00",
							"timeZone": "America/New_York"
						}
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

			_, _, err := ns.GetPolicy(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetPolicy() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("GetPolicy() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestDeletePolicy(t *testing.T) {

}
