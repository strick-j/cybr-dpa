package dpa

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
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
// Valid query parameter keys are:
//   - b64StartKey - Next page to retrieve if last response returned a value for
//
// b64_last_evaluated_key
//   - name - Target set name to filter with, in wildcard format
//   - strongAccountId - Strong account ID to filter target sets list with
//
// Returns types.ListTargetSetesponse or types.ErrorResponse based on the
// response from the API. An error is returned on request failure
// Example:
//
//	// List Target Sets with query
//	query := map[string]string{"name":"example.com"}
//
//	resp, errResp, err := s.ListTargetSets(context.Background(), query)
//	if err != nil {
//		log.Fatalf("Failed to list target sets. %s", err)
//		return
//	}
func (s *Service) ListTargetSets(ctx context.Context, query interface{}) (*types.ListTargetSetResponse, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	var path string

	// Check if query parameters were passed
	if query != nil {
		// Check to see if query was passed as map[string]string
		v, ok := query.(map[string]string)
		if !ok {
			defer cancelCtx()
			return nil, nil, fmt.Errorf("getPublicKey: Please pass query parameters via map[string]string")
		}

		// Parse query parameters
		q := url.Values{}
		for a, b := range v {
			if len(b) != 0 {
				q.Add(a, b)
			}
		}
		// Create path using query paramters and make request via service client
		path = fmt.Sprintf("/discovery/targetsets?%s", q.Encode())
	} else if query == nil {
		// Create path with no query parameters
		path = "/discovery/targetsets"
	}

	if err := s.client.Get(ctx, path, &listTargetSetResponse, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("getTargetSet: Failed to retrieve Target Sets. %s", err)
	}

	defer cancelCtx()
	return &listTargetSetResponse, &errorResponse, nil
}

// AddTargetSet adds a target set or multiple target sets
// The request body should be a struct containing an array of target sets
// Struct is defined in pkg/cybr/dpa/types/dicovery.go as TargetSetMapping
//
// Returns types.AddTargetSetResonse or types.ErrorResponse based on the
// response from the API. An error is returned on request failure
// Example:
//
//	// TargetSet Struct
//	payload := types.TargetSetMapping{
//		StrongAccountID: "string",
//		TargetSets: []types.TargetSets{
//			Name: "string",
//			Description: "string",
//			ProvisionFormat: "string",
//			EnableCertificateValidation: bool,
//			SecretType: "string",
//			SecretID: "string",
//			Type: "string",
//		},
//	}
//
//	resp, errResp, err := s.ListTargetSets(context.Background(), payload)
//	if err != nil {
//		log.Fatalf("Failed to add target sets. %s", err)
//		return
//	}
func (s *Service) AddTargetSet(ctx context.Context, p interface{}) (*types.AddTargetSetResponse, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Verify interface is of proper type
	val := reflect.ValueOf(p)
	if val.Kind() != reflect.Struct {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("addTargetSet: Invalid type provided. Expected struct of type types.TargetSetMapping")
	}

	// Make request to add policy via service client
	if err := s.client.Post(ctx, "/discovery/targetsets", p, &addTargetSetResponse, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("addTargetSet: Failed to add Target Set. %s", err)
	}

	defer cancelCtx()
	return &addTargetSetResponse, &errorResponse, nil
}

// DeleteTargetSet provides the ability to delete target sets
// The request body should be an array of target set names
//
//	["targetset1", "targetset2"]
//
// Returns types.DeleteTargetSetesponse or types.ErrorResponse based on the
// response from the API. An error is returned on request failure
// Example:
//
//	// Create body for DeleteTargetSet Request
//	payload := []string{"targetsetid1","targetsetid2"}
//
//	// Delete Target Sets using slice
//	resp, errResp, err := s.DeleteTargetSet(context.Background(), payload)
//	if err != nil {
//		log.Fatalf("Failed to delete target sets. %s", err)
//		return
//	}
func (s *Service) DeleteTargetSet(ctx context.Context, p interface{}) (*types.DeleteTargetSetResponse, *types.ErrorResponse, error) {
	// Set a timeout for the request
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)

	// Verify interface is of proper type
	val := reflect.ValueOf(p)
	if val.Kind() != reflect.Slice {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("deleteTargetSet: Invalid type provided. Expected slice of target sets to delete")
	}

	// Make request to delete target set(s) via service client
	path := "/discovery/targetsets/bulk"
	if err := s.client.Delete(ctx, path, p, &deleteTargetSetResponse, &errorResponse); err != nil {
		defer cancelCtx()
		return nil, nil, fmt.Errorf("deleteTargetSet: Failed to delete target set. %s", err)
	}

	defer cancelCtx()
	return &deleteTargetSetResponse, &errorResponse, nil
}
