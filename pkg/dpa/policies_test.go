package dpa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
)

// Create sample policy used in post and put tests
var validSamplePolicy = types.Policy{
	PolicyName: "Test Policy",
	Status:     "Enabled",
	ProvidersData: types.ProvidersData{
		Aws: types.Aws{
			Regions:    []string{"us-east-1"},
			Tags:       []types.Tags{},
			VpcIds:     []string{},
			AccountIds: []string{},
		},
	},
	StartDate: "2024-01-10",
	EndDate:   "2025-01-10",
	UserAccessRules: []types.UserAccessRules{
		{
			RuleName: "Example Rule",
			UserData: types.UserData{
				Roles: []types.Roles{
					{
						Name: "Dev Test Role",
					},
				},
			},
			ConnectionInformation: types.ConnectionInformation{
				ConnectAs: types.ConnectAs{
					Aws: types.ConnectAsAws{
						SSH: "ec2-user",
					},
				},
				GrantAccess: 3,
				IdleTime:    10,
				DaysOfWeek:  []string{"Mon", "Tue"},
				FullDays:    true,
				TimeZone:    "Asia/Jerusalem",
			},
		},
	},
}

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
			header:  http.StatusOK,
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
			name:    "Empty Policy ID",
			input:   "",
			header:  http.StatusOK,
			sleep:   1 * time.Millisecond,
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

func TestAddPolicy(t *testing.T) {
	// Modify valid sample policy to have invalid date
	var invalidSamplePolicy = validSamplePolicy
	invalidSamplePolicy.StartDate = "2010-10-10"

	var tests = []struct {
		name     string
		input    interface{}
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:    "Invalid Timeout",
			input:   validSamplePolicy,
			header:  http.StatusOK,
			sleep:   6 * time.Second,
			wantErr: true,
		},
		{
			name:    "Invalid Type",
			input:   "Example String",
			wantErr: true,
		},
		{
			name:   "Invalid Date In the Past",
			input:  invalidSamplePolicy,
			header: http.StatusBadRequest,
			response: `{
				"code": "DPA_CRUD_ACTION_FAILED",
				"message": "Unable to create an Authorization Policy.",
				"description": "Unable to create an Authorization Policy. Error(s): Policy start date must not be in the past (field: start_date)",
				"doc": null,
				"steps": null
			}`,
			sleep:   1 * time.Millisecond,
			wantErr: false,
		},
		{
			name:  "Valid Policy Add Response",
			input: validSamplePolicy,
			response: `{
				"policyId": "07280193-cfd3-4155-b3dd-232a77a6c72a"
			}`,
			header:  http.StatusCreated,
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

			_, _, err := ns.AddPolicy(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("AddPolicy() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("AddPolicy() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}

}

func TestUpdatePolicy(t *testing.T) {
	// Add Policy ID to validSamplePolicy Struct
	validSamplePolicy.PolicyID = "c12f982a-ab1a-12ab-1a31-f221aa31836a"
	validSamplePolicy.Status = "Disabled"
	var tests = []struct {
		name     string
		input    interface{}
		id       string
		header   http.ConnState
		sleep    time.Duration
		response string
		wantErr  bool
	}{
		{
			name:    "Invalid Timeout",
			input:   validSamplePolicy,
			id:      "c12f982a-ab1a-12ab-1a31-f221aa31836a",
			header:  http.StatusOK,
			sleep:   6 * time.Second,
			wantErr: true,
		},
		{
			name:    "Invalid Payload Type",
			input:   "Example String",
			id:      "c12f982a-ab1a-12ab-1a31-f221aa31836a",
			wantErr: true,
		},
		{
			name:    "Invalid Empty Policy ID",
			input:   validSamplePolicy,
			id:      "",
			wantErr: true,
		},
		{
			name:   "Invalid Policy ID Mismatch",
			input:  validSamplePolicy,
			id:     "c12f322a-ab1a-12ab-1a31-f221aa31836b",
			header: http.StatusBadRequest,
			response: `{
				"code": "DPA_CRUD_ACTION_FAILED",
				"message": "Unable to update an Authorization Policy.",
				"description": "Unable to update an Authorization Policy. Error(s): Policy ID in the path parameter is not the same as the policy ID in the request body",
				"doc": null,
				"steps": null
			}`,
			sleep:   1 * time.Millisecond,
			wantErr: false,
		},
		{
			name:  "Valid Policy Update Response",
			id:    "76321c1a-32ff-488e-af40-b7ae5eefb808",
			input: validSamplePolicy,
			response: `{
				"policyId": "76321c1a-32ff-488e-af40-b7ae5eefb808",
				"policyName": "Test Policy 4",
				"status": "Disabled",
				"description": "",
				"providersData": {
					"AWS": {
						"regions": [
							"us-east-1",
                			"us-east-2"
						],
						"tags": [
							{
								"Key": "env",
								"Value": [
									"Prod",
									"Dev"
								]
							}
						],
						"vpcIds": [],
						"accountIds": []
					}
				},
				"startDate": "2024-01-10",
				"endDate": "2025-01-10",
				"userAccessRules": [
					{
						"ruleName": "dev-team-access",
						"userData": {
							"roles": [
								{
									"name":"DEV_TEAM_ROLE"
								}
							],
							"groups": [],
							"users": []
						},
						"connectionInformation": {
							"connectAs": {
								"AWS": {
									"ssh": "ec2-user"
								}
							},
							"grantAccess": 3,
							"idleTime": 10,
							"EffectiveSessionDuration": null,
							"daysOfWeek": [
								"Mon",
								"Tue"
							],
							"fullDays": true,
							"hoursFrom": null,
							"hoursTo": null,
							"timeZone": "Asia/Jerusalem"
						}
					}
				]
			}`,
			header:  http.StatusCreated,
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

			_, _, err := ns.UpdatePolicy(context.Background(), tt.input, tt.id)
			if tt.wantErr {
				if err == nil {
					t.Errorf("UpdatePolicy() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("UpdatePolicy() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}

}

func TestDeletePolicy(t *testing.T) {
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
				"message": "Unable to delete an Authorization Policy.",
				"description": "Unable to delete an Authorization Policy. Error(s): badly formed hexadecimal UUID string",
				"doc": null,
				"steps": null
			}`,
			sleep:   1 * time.Millisecond,
			wantErr: false,
		},
		{
			name:    "Valid Delete Response",
			input:   "c12f982a-ab1a-12ab-1a31-f221aa31836a",
			header:  http.StatusNoContent,
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

			_, err := ns.DeletePolicy(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("DeletePolicy() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("DeletePolicy() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
