package dpa

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"time"

	"github.com/strick-j/cybr-dpa/pkg/dpa/types"
)

var (
	generateScriptResponse types.GenerateScriptResponse
)

// GenerateScript generates a request for a connector setup script
// Body is optional, if nothing is provided a default script will be generated
// The default script will be for a linux connector in AWS
// Returns a GenerateScriptResponse or error if failed
// Example:
//
//	// Create Body for GenerateScript Request
//	generateScriptRequest := struct {
//		ConnectorOS   string `json:"connectorOs,omitempty"`
//		ConnectorType string `json:"connectorType,omitempty"`
//	}{
//		"linux",
//		"AWS"
//	}
//
//	// Generate Script using existing Service and Client
//	apps, err := s.GenerateScript(context.Background(), generateScriptRequest)
//	if err != nil {
//		log.Fatalf("Failed to generate connector script. %s", err)
//		return
//	}
//
//	fmt.Println(apps)
func (s *Service) GenerateScript(ctx context.Context, p interface{}) (*types.GenerateScriptResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	// Validate provided type
	val := reflect.ValueOf(p)
	if val.Kind() != reflect.Struct {
		defer cancelCtx()
		return nil, fmt.Errorf("generateScript: Invalid type provided %s. Must be a struct", val.Kind())
	}

	// Validate provided fields
	if val.Field(0).String() != "ConnectorOS" || val.Field(1).String() != "ConnectorType" {
		defer cancelCtx()
		return nil, fmt.Errorf("generateScript: Invalid field provided %s. Must be ConnectorOS", val.Field(0).String())
	}

	// Validate provided values
	validOS := []string{"linux", "windows", "darwin"}
	validType := []string{"AWS", "AZURE", "GCP", "ON-PREMISE"}

	// Validate provided OS
	containsOs := slices.Contains(validOS, val.FieldByName("ConnectorOS").String())
	if !containsOs {
		defer cancelCtx()
		return nil, fmt.Errorf("generateScript: Invalid OS provided %s. Valid options are linux, windows, darwin", val.FieldByName("ConnectorOS").String())
	}

	// Validate provided Type
	containsType := slices.Contains(validType, val.FieldByName("ConnectorType").String())
	if !containsType {
		defer cancelCtx()
		return nil, fmt.Errorf("generateScript: Invalid Type provided %s. Valid options are AWS, AZURE, GCP, ON-PREMISE", val.FieldByName("ConnectorType").String())
	}

	// Make request for connector setup script via service client
	if err := s.client.Post(ctx, "/connectors/setup-script", p, &generateScriptResponse); err != nil {
		defer cancelCtx()
		return nil, fmt.Errorf("generateScript: Failed to retrieve script. %s", err)
	}

	defer cancelCtx()
	return &generateScriptResponse, nil
}
