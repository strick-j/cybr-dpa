package dpa

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
)

var (
	addTargetSetResponse    types.AddTargetSetResponse
	listTargetSetResponse   types.ListTargetSetResponse
	deleteTargetSetResponse types.DeleteTargetSetResponse
)

// ListTargetSets returns a list of target sets
// Query parameters can be used to filter the results and are optional
func (s *Service) ListTargetSets(ctx context.Context, query map[string]string) (*types.ListTargetSetResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Parse query parameters
	q := url.Values{}
	for a, b := range query {
		if len(b) != 0 {
			q.Add(a, b)
		}
	}

	// Create path using query paramters and make request via service client
	path := fmt.Sprintf("/discovery/targetsets?%s", q.Encode())
	log.Printf("path: %s", path)
	if err := s.client.Get(ctx, path, &listTargetSetResponse); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("getTargetSet: Failed to retrieve Target Sets. %s", err)
	}

	defer cancelCtx()
	return &listTargetSetResponse, nil
}

// AddTargetSet adds a target set or multiple target sets
// The request body should be a struct containing an array of target sets
// Struct is defined in pkg/cybr/dpa/types/dicovery.go as TargetSetMapping
func (s *Service) AddTargetSet(ctx context.Context, p interface{}) (*types.AddTargetSetResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Make request to add policy via service client
	if err := s.client.Post(ctx, "/discovery/targetsets", p, &addTargetSetResponse); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("addTargetSet: Failed to add Target Set. %s", err)
	}

	defer cancelCtx()
	return &addTargetSetResponse, nil
}

// Provides the ability to delete target sets
// The request body should be an array of target set names
// e.g. ["targetset1", "targetset2"]
func (s *Service) DeleteTargetSet(ctx context.Context, n []string) (*types.DeleteTargetSetResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Check if target set name(s) are empty
	if len(n) < 1 {
		defer cancelCtx()
		return nil, fmt.Errorf("deleteTargetSet: Target Set name(s) cannot be empty")
	}

	// Make request to delete target set(s) via service client
	path := "/discovery/targetsets/bulk"
	if err := s.client.Delete(ctx, path, n, &deleteTargetSetResponse); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("deleteTargetSet: Failed to delete target set. %s", err)
	}

	defer cancelCtx()
	return &deleteTargetSetResponse, nil
}
