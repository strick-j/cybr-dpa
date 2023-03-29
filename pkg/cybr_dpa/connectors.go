package cybr_dpa

import (
	"context"
	"fmt"
)

type ConnectorScript struct {
	ScriptUrl   string `json:"script_url,omitempty"`
	BashCommand string `json:"bash_cmd,omitempty"`
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

type ConnectorRequest struct {
	ConnectorType string `json:"connector_type"`
	ConnectorOS   string `json:"connector_os"`
}

var (
	connectorScript ConnectorScript
)

// GetConnectors: Generate a signed URL to a connector installation script for the platform you select. The script contains a secret token that is valid for 15 minutes from the time it is generated.
//
// Example Usage:
//
//	connectorRequest := ConnectorRequest {
//	  ConnectorType: "AWS",
//	  ConnectorOS: "linux"
//	}
//
//	getConnectorScript, err := s.GetConnectorScript(context.Background, connectorRequest)
func (s *Service) GetConnectorScript(ctx context.Context, connectorRequest ConnectorRequest) (*ConnectorScript, error) {
	allowedType := []string{"AWS", "AZURE", "ON-PREMISE"}

	if typeAllowed := contains(allowedType, connectorRequest.ConnectorType); !typeAllowed {
		return nil, fmt.Errorf("connector type not allowed, valid types are AWS, AZURE, ON-PREMISE")
	}

	if err := s.client.Post(ctx, fmt.Sprintf("/%s/%s", "connectors", "setup-script"), connectorRequest, &connectorScript); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}

	return &connectorScript, nil
}
