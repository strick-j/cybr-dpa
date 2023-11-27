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

func (s *Service) ListPolicies(ctx context.Context) (*types.ListPolicies, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 20000*time.Millisecond)
	if err := s.client.Get(ctx, "/access-policies", &listPolicies); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("listPolicies: Failed to get access policies. %s", err)
	}

	defer cancelCtx()
	return &listPolicies, nil
}

func (s *Service) GetPolicy(ctx context.Context, p string) (*types.Policy, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 20000*time.Millisecond)

	// Check if policy name is empty
	if len(p) == 0 {
		defer cancelCtx()
		return nil, fmt.Errorf("getPolicy: Policy name cannot be empty")
	}

	// Create path and get policy using policy id
	path := fmt.Sprintf("/access-policies/%s", p)
	if err := s.client.Get(ctx, path, &getPolicy); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("lgetPolicies: Failed to get access policy. %s", err)
	}

	defer cancelCtx()
	return &getPolicy, nil
}

// TODO: Test this function
func (s *Service) AddPolicy(ctx context.Context, p interface{}) (*types.AddPolicy, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Make request to add policy via service client
	if err := s.client.Post(ctx, "/access-policies", p, &addPolicy); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("deleteTargetSet: Failed to add policy. %s", err)
	}

	defer cancelCtx()
	return &addPolicy, nil
}

// TODO: Test this function
func (s *Service) DeletePolicy(ctx context.Context, p string) (*string, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Make request to delete policy via service client
	path := fmt.Sprintf("/access-policies/%s", p)
	if err := s.client.Delete(ctx, path, nil, &deletePolicy); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("deleteTargetSet: Failed to delete policy. %s", err)
	}

	defer cancelCtx()
	return &deletePolicy, nil
}