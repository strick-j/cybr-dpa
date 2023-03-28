package cybr_dpa

import (
	"context"
	"fmt"

	"github.com/strick-j/cybr-dpa/pkg/cybr_dpa/types"
)

var (
	ConnectorScript types.ConnectorScript
)

// GetConnectors: Generate a signed URL to a connector installation script for the platform you select. The script contains a secret token that is valid for 15 minutes from the time it is generated.
//
// Example Usage:
//
//	connectorRequest := types.ConnectorRequest {
//		ConnectorType: "AWS",
//		ConnectorOS: "linux"
//	}
//	getConnectorScript, err := s.GetConnectorScript(context.Background, connectorRequest)
func (s *Service) GetConnectorScript(ctx context.Context, connectorRequest types.ConnectorRequest) (*types.ConnectorScript, error) {
	//allowedType := []string{"AWS", "AZURE", "ON-PREMISE"}

	//if typeAllowed := contains(allowedType, connectorRequest.ConnectorType); !typeAllowed {
	//	return nil, fmt.Errorf("connector type not allowed, valid types are AWS, AZURE, ON-PREMISE")
	//}

	if err := s.client.Post(ctx, fmt.Sprintf("/%s/%s", "connectors", "setup-script"), connectorRequest, &ConnectorScript); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}

	return &ConnectorScript, nil
}
