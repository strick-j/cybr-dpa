package cybr_dpa

import (
	"context"
	"fmt"
	"net/url"
	//"github.com/strick-j/cybr-dpa/pkg/cybr_dpa/types"
)

var (
	PublicKey string
)

// GetConnectors: Returns an SSH CA public key for the specified workspace. Append this CA user key to your existing Trusted CA user keys file on all the target machines.
//
// Example Usage:
//
//	getPublicKeys, err := s.GetPublicKeys(context.Background, "cb5544d2-678e7-45f0-823e-555dc6f38ea6", "Azure")
func (s *Service) GetPublicKeys(ctx context.Context, workspaceID, workspaceType string) (*string, error) {

	allowedType := []string{"AWS", "Azure"}
	if typeAllowed := contains(allowedType, workspaceType); !typeAllowed {
		return nil, fmt.Errorf("connector type not allowed, valid types are AWS, AZURE, ON-PREMISE")
	}

	pathEscapedQuery := url.PathEscape("workpaceid=" + workspaceID + "&workspacetype" + workspaceType)
	if err := s.client.Get(ctx, fmt.Sprintf("/%s?%s", "public-keys", pathEscapedQuery), &PublicKey); err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	return &PublicKey, nil
}
