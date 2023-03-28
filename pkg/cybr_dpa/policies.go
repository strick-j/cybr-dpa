package cybr_dpa

import (
	"context"
	"fmt"

	"github.com/strick-j/cybr-dpa/pkg/cybr_dpa/types"
)

var (
	Policies types.Policies
)

// GetPolcies: Returns all authorization policies
//
// Example Usage:
//
//	getPolicies, err := s.GetPolicies(context.Background)
func (s *Service) GetPolicies(ctx context.Context) (*types.Policies, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "access-polcies"), &Policies); err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	return &Policies, nil
}
