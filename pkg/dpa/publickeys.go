package dpa

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
)

// Local variables for the package
var (
	publicKey       string
	publicKeyScript types.PublicKeyScript
)

// GetPublicKey returns the public key for the DPA Workspace
// Expects an ordered map of the query parameters
// Returns a PublicKey or error if failed
func (s *Service) GetPublicKey(ctx context.Context, query map[string]string) (*types.PublicKey, *types.ErrorResponse, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Validate both query parameters were provided. If not, return error
	if len(query) < 2 {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPublicKey: Missing required parameters")
	}

	// Parse query parameters
	q := url.Values{}
	for a, b := range query {
		q.Add(a, b)
	}

	// Create URL and make request via service client
	path := fmt.Sprintf("/public-keys?%s", q.Encode())
	if err := s.client.Get(ctx, path, &publicKey, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPublicKey: Failed to retrieve public key. %s", err)
	}

	defer cancelCtx()
	return &types.PublicKey{PublicKey: publicKey}, &errorResponse, nil
}

// GetPublicKeyScript returns the public key setup script
// Expects an ordered map of the query parameters
// Returns a PublicKey struct or error if failed
func (s *Service) GetPublicKeyScript(ctx context.Context, query map[string]string) (*types.PublicKeyScript, *types.ErrorResponse, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Validate query parameters were provided. If not, return error
	if len(query) < 2 {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPublicKeyScript: Missing required parameters")
	}

	// Parse query parameters
	q := url.Values{}
	for a, b := range query {
		q.Add(a, b)
	}

	// Create path and make request via service client
	path := fmt.Sprintf("/public-keys/scripts?%s", q.Encode())
	if err := s.client.Get(ctx, path, &publicKeyScript, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPublicKeyScript: Failed to retrieve public key installation script. %s", err)
	}

	defer cancelCtx()
	return &publicKeyScript, &errorResponse, nil
}
