package cybr_dpa

import (
	"context"
	"fmt"
)

// GetConnectors: Generate a signed URL to a connector installation script for the platform you select. The script contains a secret token that is valid for 15 minutes from the time it is generated.
//
// Example Usage:
//	getConnectorScript, err := s.GetConnectorScript(context.Background, AWS, linux)

func (s *Service) GetConnectorScript(ctx context.Context, connectorType, connectorOS string) (*types.Connectors, error) {
	allowedType := []string{"AWS", "AZURE", "ON-PREMISE"}

	if typeAllowed := contains(allowedType, connectorType); typeAllowed != true {
		return nil, fmt.Errorf("connector type not allowed, valid types are AWS, AZURE, ON-PREMISE")
	}

	if err := s.client.Post(ctx, fmt.Sprintf("/%s/%s/%s", "api", "Connectors", "setup-script"), &Connectors); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}
}
