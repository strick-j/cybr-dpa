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
// Returns a PublicKey, DPA Error Response, or generic error if failed.
//
// Example:
//
//	// Create query for GetPublicKey
//	query := map[string]string{"workspaceId":"12347578363","workspaceType":"AWS"}
//
//	// Call GetPublicKey wtih query
//	apps, dpaerr, err := s.GetPublicKey(context.Background(), query)
//	if err != nil {
//		log.Fatalf("Failed to retrieve public key. %s", err)
//		return
//	}
func (s *Service) GetPublicKey(ctx context.Context, query interface{}) (*types.PublicKey, *types.ErrorResponse, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Check to see if query was passed as map[string]string
	v, ok := query.(map[string]string)
	if !ok {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPublicKey: Please pass query parameters via map[string]string")
	}

	// Validate both query parameters were provided. If not, return error
	if len(v) != 2 {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPublicKey: Missing required parameters")
	}

	// Parse query parameters
	q := url.Values{}
	for a, b := range v {
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

// GetPublicKeyScript returns the public key script for the DPA Workspace
// Expects an ordered map of the query parameters
// Returns a PublicKeyScript, DPA Error Response, or generic error if failed.
//
// Example:
//
//	// Create query for GetPublicKeyScript
//	query := map[string]string{"workspaceId":"12347578363","workspaceType":"AWS"}
//
//	// Call GetPublicKeyScript wtih query
//	apps, dpaerr, err := s.GetPublicKeyScript(context.Background(), query)
//	if err != nil {
//		log.Fatalf("Failed to generate public key script. %s", err)
//		return
//	}
func (s *Service) GetPublicKeyScript(ctx context.Context, query interface{}) (*types.PublicKeyScript, *types.ErrorResponse, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Check to see if query was passed as map[string]string
	v, ok := query.(map[string]string)
	if !ok {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPublicKey: Please pass query parameters via map[string]string")
	}

	// Validate query parameters were provided. If not, return error
	if len(v) != 2 {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getPublicKeyScript: Missing required parameters")
	}

	// Parse query parameters
	q := url.Values{}
	for a, b := range v {
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
