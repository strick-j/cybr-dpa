package cybr_dpa

import (
	"context"
	"fmt"
	"net/url"
)

type PublicKey struct {
	PublicKey string `json:"publickey,omitempty"`
}

type PublicKeyScript struct {
	Base64Cmd string `json:"base64_cmd,omitempty"`
}

type PubKeyError struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Errors      []struct {
		Code        string `json:"code,omitempty"`
		Message     string `json:"message,omitempty"`
		Description string `json:"description,omitempty"`
		Field       string `json:"field,omitempty"`
	} `json:"errors,omitempty"`
}

var (
	publicKey       PublicKey
	publicKeyScript PublicKeyScript
)

// GetPublicKey: Returns an SSH CA public key for the specified workspace. Append this CA user key to your existing Trusted CA user keys file on all the target machines.
//
// Example Usage:
//
//	getPublicKey, err := s.GetPublicKey(context.Background, "cb5544d2-678e7-45f0-823e-555dc6f38ea6", "Azure")
func (s *Service) GetPublicKey(ctx context.Context, workspaceID, workspaceType string) (*PublicKey, error) {

	allowedType := []string{"AWS", "Azure"}
	if typeAllowed := contains(allowedType, workspaceType); !typeAllowed {
		return nil, fmt.Errorf("connector type not allowed, valid types are AWS, AZURE, ON-PREMISE")
	}

	pathEscapedQuery := url.PathEscape("workpaceId=" + workspaceID + "&workspaceType=" + workspaceType)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "public-keys", pathEscapedQuery), &publicKey); err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	return &publicKey, nil
}

// GetPublicKeyScript: Generates an SSH CA public key plus a deployment script for the specified workspace. Use this script to install the CyberArk certificate on target machines that don't have an existing certificate file.
//
// Example Usage:
//
//	getPublicKeyScript, err := s.GetPublicKeyScript(context.Background, "cb5544d2-678e7-45f0-823e-555dc6f38ea6", "Azure")
func (s *Service) GetPublicKeyScript(ctx context.Context, workspaceID, workspaceType string) (*PublicKeyScript, error) {

	allowedType := []string{"AWS", "Azure"}
	if typeAllowed := contains(allowedType, workspaceType); !typeAllowed {
		return nil, fmt.Errorf("connector type not allowed, valid types are AWS, AZURE, ON-PREMISE")
	}

	pathEscapedQuery := url.PathEscape("workpaceId=" + workspaceID + "&workspaceType=" + workspaceType)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s?%s", "public-keys", "scripts", pathEscapedQuery), &publicKeyScript); err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	return &publicKeyScript, nil
}
