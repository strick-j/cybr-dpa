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
	errorResponse          types.ErrorResponse
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
func (s *Service) GenerateScript(ctx context.Context, p interface{}) (*types.GenerateScriptResponse, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 10000*time.Millisecond)

	if err := parameterValidation(p); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("generateScript: Parameter validation failed. %s", err)
	}

	// Make request for connector setup script via service client
	if err := s.client.Post(ctx, "/connectors/setup-script", p, &generateScriptResponse, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("generateScript: Failed to retrieve script. %s", err)
	}

	defer cancelCtx()
	return &generateScriptResponse, &errorResponse, nil
}

// Validates proper parameters were passed for the GenerateScript API endpoint
func parameterValidation(p interface{}) error {
	// Validate provided type
	val := reflect.ValueOf(p)
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("parameterValidation: Invalid type provided %s. Must be a struct", val.Kind())
	}

	// Validate provided fields
	field1 := val.FieldByName("ConnectorOS")
	field2 := val.FieldByName("ConnectorType")
	if !field1.IsValid() || !field2.IsValid() {
		return fmt.Errorf("parameterValidation: Invalid fields provided %s, %s. Must be ConnectorOS and ConnectorType", val.Field(0).String(), val.Field(1).String())
	}

	// Validate provided values
	validOS := []string{"linux", "windows", "darwin"}
	validType := []string{"AWS", "AZURE", "GCP", "ON-PREMISE"}

	// Validate provided  Connector OS
	containsOs := slices.Contains(validOS, val.FieldByName("ConnectorOS").String())
	if !containsOs {
		return fmt.Errorf("parameterValidation: Invalid Connector OS provided %s. Valid options are linux, windows, darwin", val.FieldByName("ConnectorOS").String())
	}

	// Validate provided Connector Type
	containsType := slices.Contains(validType, val.FieldByName("ConnectorType").String())
	if !containsType {
		return fmt.Errorf("parameterValidation: Invalid Connector Type provided %s. Valid options are AWS, AZURE, GCP, ON-PREMISE", val.FieldByName("ConnectorType").String())
	}

	return nil
}
