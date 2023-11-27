package dpa

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
)

var (
	settings       types.Settings
	featureSetting types.FeatureSetting
)

// ListSettings retrieves the settings for the DPA instance
// Returns a Settings struct or error if failed
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

func (s *Service) ListSettingsFeature(ctx context.Context, f string) (*types.FeatureSetting, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Validate provided feature name
	if strings.ToLower(f) == "mfacaching" {
		f = "mfaCaching"
	} else if strings.ToLower(f) == "sshcommandaudit" {
		f = "sshCommandAudit"
	} else {
		err := fmt.Errorf("getSettings: Invalid Feature provided %s. Valid options are mfaCaching and sshCommandAudit", f)
		defer cancelCtx()
		return nil, nil, err
	}

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
