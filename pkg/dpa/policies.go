package dpa

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
)

var (
	deletePolicy string

	addPolicy    types.AddPolicy
	listPolicies types.ListPolicies
	getPolicy    types.Policy
	policy       types.Policy
)

// ListPolicies returns all of the currently configured policies
// Returns types.ListPolicies or types.ErrorResponse based on the
// response from the API. An error is returned on request failure
// Example:
//
//	resp, errResp, err := s.ListPolicies(context.Background())
//	if err != nil {
//		log.Fatalf("Failed to list policies. %s", err)
//		return
//	}
func (s *Service) ListPolicies(ctx context.Context) (*types.ListPolicies, *types.ErrorResponse, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)
	if err := s.client.Get(ctx, "/access-policies", &listPolicies, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("listPolicies: Failed to get access policies. %s", err)
	}

	defer cancelCtx()
	return &listPolicies, &errorResponse, nil
}

// GetPolicy returns a specific policy
// Returns types.Policy or types.ErrorResponse based on the
// response from the API. An error is returned on request failure
// Example:
//
//	policyID := "c12f982a-ab1a-12ab-1a31-f221aa31836b"
//	resp, errResp, err := s.ListPolicies(context.Background(), policyID)
//	if err != nil {
//		log.Fatalf("Failed to list policies. %s", err)
//		return
//	}
func (s *Service) GetPolicy(ctx context.Context, i string) (*types.Policy, *types.ErrorResponse, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Check if policy name is empty
	if len(i) == 0 {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPolicy: Policy name cannot be empty")
	}

	// Create path and get policy using policy id
	path := fmt.Sprintf("/access-policies/%s", i)
	if err := s.client.Get(ctx, path, &getPolicy, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("lgetPolicies: Failed to get access policy. %s", err)
	}

	defer cancelCtx()
	return &getPolicy, &errorResponse, nil
}

// Add Policy creates a new policy
// Expects a struct of type types.Policy
// Returns types.AddPolicy or types.ErrorResponse based on the
// response from the API. An error is returned on request failure.
// Example:
//
//	 // Fill out policy Information
//	 var validSamplePolicy = types.Policy{
//			 PolicyName: "Test Policy",
//			 Status:     "Enabled",
//			 ProvidersData: types.ProvidersData{
//				 Aws: types.Aws{
//					 Regions:    []string{"us-east-1"},
//					 Tags:       []types.Tags{},
//					 VpcIds:     []string{},
//					 AccountIds: []string{},
//				 },
//			 },
//			 StartDate: "2024-01-10",
//			 EndDate:   "2025-01-10",
//			 UserAccessRules: []types.UserAccessRules{
//				 {
//					 RuleName: "Example Rule",
//					 UserData: types.UserData{
//						 Roles: []types.Roles{
//							 {
//								 Name: "Example Role",
//							 },
//						 },
//					 },
//					 ConnectionInformation: types.ConnectionInformation{
//						 ConnectAs: types.ConnectAs{
//							 Aws: types.ConnectAsAws{
//								 SSH: "ec2-user",
//							 },
//						 },
//						 GrantAccess: 3,
//						 IdleTime:    10,
//						 DaysOfWeek:  []string{"Mon", "Tue"},
//						 FullDays:    true,
//						 TimeZone:    "Asia/Jerusalem",
//					 },
//				 },
//			 },
//		 }
//
//		 resp, err := s.AddPolicy(context.Background(), validSamplePolicy)
//		 if err != nil {
//			 log.Fatalf("Failed to add policy. %s", err)
//			 return
//		 }
func (s *Service) AddPolicy(ctx context.Context, p interface{}) (*types.AddPolicy, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Validate provided type
	val := reflect.ValueOf(p)
	if val.Kind() != reflect.Struct {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("addPolicy: Invalid type provided. Expected struct of format types.Policy")
	}

	// Make request to add policy via service client
	if err := s.client.Post(ctx, "/access-policies", p, &addPolicy, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("addPolicy: Failed to add policy. %s", err)
	}

	defer cancelCtx()
	return &addPolicy, &errorResponse, nil
}

func (s *Service) UpdatePolicy(ctx context.Context, p interface{}, i string) (*types.Policy, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Check if policy name is empty
	if len(i) == 0 {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("updatePolicy: Policy id cannot be empty")
	}

	// Validate provided type
	val := reflect.ValueOf(p)
	if val.Kind() != reflect.Struct {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("updatePolicy: Invalid type provided. Expected struct of format types.Policy")
	}

	// Create path and get policy using policy id
	path := fmt.Sprintf("/access-policies/%s", i)

	// Make request to add policy via service client
	if err := s.client.Post(ctx, path, p, &policy, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("updatePolicy: Failed to add policy. %s", err)
	}

	defer cancelCtx()
	return &policy, &errorResponse, nil
}

// DeletePolicy deletes a specific policy
// Returns no response if succesfull or types.ErrorResponse based on the
// response from the API. An error is returned on request failure.
// Example:
//
//	 policyID := "c12f982a-ab1a-12ab-1a31-f221aa31836a"
//
//	 resp, err := s.DeletePolicy(context.Background(),policyID)
//	 if err != nil {
//		 log.Fatalf("Failed to delete policy. %s", err)
//		 return
//	 }
func (s *Service) DeletePolicy(ctx context.Context, p string) (*types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Make request to delete policy via service client
	path := fmt.Sprintf("/access-policies/%s", p)
	if err := s.client.Delete(ctx, path, nil, &deletePolicy, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("deletepolicy: Failed to delete policy. %s", err)
	}

	defer cancelCtx()
	return &errorResponse, nil
}
