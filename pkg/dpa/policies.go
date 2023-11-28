package dpa

import (
	"context"
	"fmt"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
)

var (
	deletePolicy string

	addPolicy    types.AddPolicy
	listPolicies types.ListPolicies
	getPolicy    types.Policy
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
func (s *Service) GetPolicy(ctx context.Context, p string) (*types.Policy, *types.ErrorResponse, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Check if policy name is empty
	if len(p) == 0 {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPolicy: Policy name cannot be empty")
	}

	// Create path and get policy using policy id
	path := fmt.Sprintf("/access-policies/%s", p)
	if err := s.client.Get(ctx, path, &getPolicy, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("lgetPolicies: Failed to get access policy. %s", err)
	}

	defer cancelCtx()
	return &getPolicy, &errorResponse, nil
}

// TODO: Test this function
func (s *Service) AddPolicy(ctx context.Context, p interface{}) (*types.AddPolicy, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Make request to add policy via service client
	if err := s.client.Post(ctx, "/access-policies", p, &addPolicy, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("deleteTargetSet: Failed to add policy. %s", err)
	}

	defer cancelCtx()
	return &addPolicy, &errorResponse, nil
}

// TODO: Test this function
func (s *Service) DeletePolicy(ctx context.Context, p string) (*string, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Make request to delete policy via service client
	path := fmt.Sprintf("/access-policies/%s", p)
	if err := s.client.Delete(ctx, path, nil, &deletePolicy, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("deleteTargetSet: Failed to delete policy. %s", err)
	}

	defer cancelCtx()
	return &deletePolicy, &errorResponse, nil
}
