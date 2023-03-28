package cybr_dpa

import (
	"context"
	"fmt"

	"github.com/strick-j/cybr-dpa/pkg/cybr_dpa/types"
)

var (
	Policies types.Policies
)

// GetPolicies: Returns all authorization policies
//
// Example Usage:
//
//	getPolicies, err := s.GetPolicies(context.Background)
func (s *Service) GetPolicies(ctx context.Context) (*types.Policies, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "access-policies"), &Policies); err != nil {
		return nil, fmt.Errorf("failed to get policies: %w", err)
	}

	return &Policies, nil
}

// GetPolicyById: Retrieves authorization policy for the given ID
//
// Example Usage:
//
//	getPolicies, err := s.GetPolicyById(context.Background, "{policy_id}")
func (s *Service) GetPolicyById(ctx context.Context, policyId string) (*types.Policies, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "access-policies", policyId), &Policies); err != nil {
		return nil, fmt.Errorf("failed to get policy: %w", err)
	}

	return &Policies, nil
}