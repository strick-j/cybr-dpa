package dpa

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
)

var (
	settings       types.Settings
	featureSetting types.FeatureSetting
)

// ListSettings provides all settings as a response.
// Returns a types.Settings response or types.ErrorResponse based on the
// response from the API. An error is returned on request failure
// Example:
//
//	// List Settings
//	resp, errResp, err := s.ListSettings(context.Background())
//	if err != nil {
//		log.Fatalf("Failed to retrieve settings. %s", err)
//		return
//	}
func (s *Service) ListSettings(ctx context.Context) (*types.Settings, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Make request for settings via service client
	if err := s.client.Get(ctx, "/settings", &settings, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getSettings: Failed to retrieve settings. %s", err)
	}
	defer cancelCtx()
	return &settings, &errorResponse, nil
}

// ListSettingsFeature provides a specific setting reponse.
// Valid input strings are:
// 'MFA_CACHING', 'STANDING_ACCESS', 'SSH_COMMAND_AUDIT', 'RDP_FILE_TRANSFER', 'CERTIFICATE_VALIDATION'
// Returns a types.Settings response or types.ErrorResponse based on the
// response from the API. An error is returned on request failure
// Example:
//
//	// List Settings Feature
//	resp, errResp, err := s.ListSettingsFeature(context.Background(), "MFA_CACHING")
//	if err != nil {
//		log.Fatalf("Failed to retrieve setting. %s", err)
//		return
//	}
func (s *Service) ListSettingsFeature(ctx context.Context, f string) (*types.FeatureSetting, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Make request for specific setting via service client
	if err := s.client.Get(ctx, fmt.Sprintf("%s/%s", "/settings", f), &featureSetting, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getSettings: Failed to retrieve settings. %s", err)
	}

	defer cancelCtx()
	return &featureSetting, &errorResponse, nil
}

// UpdateSettings updates the settings for the DPA instance
// Expects a struct of type types.Settings
// Returns a types.Settings response or types.ErrorResponse based on the
// response from the API. An error is returned on request failure
// Example:
//
//		// Create Body for UpdateSettings Request
//	 	updateSettingsRequest := struct {
//			IsMfaCachingEnabled  bool `json:"isMfaCachingEnabled,omitempty"`
//			KeyExpirationTimeSec int  `json:"keyExpirationTimeSec,omitempty"`
//		}{
//			true,
//			3600
//		}
//
//		// Update settings using created struct
//		resp, errResp, err := s.UpdateSettings(context.Background(), updateSettingsRequest)
//		if err != nil {
//			log.Fatalf("Failed to update settings. %s", err)
//			return
//		}
func (s *Service) UpdateSettings(ctx context.Context, p interface{}) (*types.Settings, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Validate provided type
	val := reflect.ValueOf(p)
	if val.Kind() != reflect.Struct {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("updateSettings: Invalid type provided. Expected struct of format types.Settings")
	}

	// Make request to update settings via service client
	if err := s.client.Patch(ctx, "/settings", p, &settings, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("updateSettings: Failed to update settings. %s", err)
	}

	defer cancelCtx()
	return &settings, &errorResponse, nil
}
