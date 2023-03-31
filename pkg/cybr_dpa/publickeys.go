package cybr_dpa

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type TextResponse struct {
	Response string `json:"response,omitempty"`
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
	textResponse    TextResponse
	publicKeyScript PublicKeyScript
)

// GetPublicKey: Returns an SSH CA public key for the specified workspace. Append this CA user key to your existing Trusted CA user keys file on all the target machines.
//
// Example Usage:
//
//	getPublicKey, err := s.GetPublicKey(context.Background, "cb5544d2-678e7-45f0-823e-555dc6f38ea6", "Azure")
func (s *Service) GetPublicKey(ctx context.Context, workspaceID, workspaceType string) (*TextResponse, error) {

	allowedType := []string{"AWS", "Azure"}
	if typeAllowed := contains(allowedType, workspaceType); !typeAllowed {
		return nil, fmt.Errorf("connector type not allowed, valid types are AWS, Azure, ON-PREMISE")
	}

	pathEscapedQuery := url.PathEscape("workspaceId=" + workspaceID + "&workspaceType=" + workspaceType)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "public-keys", pathEscapedQuery), &textResponse); err != nil {

		// Check to see if error response contains ssh-rsa because response is text/plain and the http service/client
		// is expecting JSON
		keyReturned := strings.Contains(fmt.Sprintf("%s", err), "ssh-rsa")
		if keyReturned {
			fmt.Printf("\nResponse contains SSH Key")
			// Split and trim string using ] delimiter
			parts := strings.SplitAfter(fmt.Sprintf("%s", err), "]")
			extractedKey := strings.TrimSpace(parts[1])
			fmt.Printf("\n%s", extractedKey)
			//return extractedKey, nil
		}

		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	return &textResponse, nil
}

// GetPublicKeyScript: Generates an SSH CA public key plus a deployment script for the specified workspace. Use this script to install the CyberArk certificate on target machines that don't have an existing certificate file.
//
// Example Usage:
//
//	getPublicKeyScript, err := s.GetPublicKeyScript(context.Background, "cb5544d2-678e7-45f0-823e-555dc6f38ea6", "Azure")
func (s *Service) GetPublicKeyScript(ctx context.Context, workspaceID, workspaceType string) (*PublicKeyScript, error) {

	allowedType := []string{"AWS", "Azure"}
	if typeAllowed := contains(allowedType, workspaceType); !typeAllowed {
		return nil, fmt.Errorf("connector type not allowed, valid types are AWS, Azure, ON-PREMISE")
	}

	pathEscapedQuery := url.PathEscape("workspaceId=" + workspaceID + "&workspaceType=" + workspaceType)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s?%s", "public-keys", "scripts", pathEscapedQuery), &publicKeyScript); err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	return &publicKeyScript, nil
}
