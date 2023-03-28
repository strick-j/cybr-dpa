package cybr_dpa

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type tokenSource struct {
	ctx  context.Context
	conf *clientcredentials.Config
}

// OauthCredClient returns a validated Oauth2 Authentication Token based on the following provided information:
//
//	clientID - Username for the Application (e.g. "identity-privilege-integration-user$@example.com")
//	clientSecret - Password for the Application
//	clientAppID - ID for the Application
//	clientURL - URL for the Application (e.g. "example.my.idaptive.app")
//	clientScope - Scopes for Application
func OauthCredClient(clientID, clientSecret, clientAppID, clientURL string, clientScope []string) (*oauth2.Token, error) {
	// Establish oauth2/clientcredentials config with user provided data
	var credentialConfig = clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://" + clientURL + "/oauth2/token/" + clientAppID,
		AuthStyle:    0,
		Scopes:       clientScope,
	}

	// Create tokenSource with provided configuration info
	ts := &tokenSource{
		ctx:  context.Background(),
		conf: &credentialConfig,
	}

	// Request new token from SCIM server using Client Credentials
	authToken, err := ts.conf.Token(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to obtain SCIM Oauth2 Token %w", err)
	}

	return authToken, nil
}

// OauthResourceOwner returns a validated Oauth2 Authentication Token with Refresh Token based on the following provided information:
//
//	clientID - Username for the Application (e.g. "identity-privilege-integration-user$@example.com")
//	clientSecret - Password for the Application
//	clientAppID - ID for the Application
//	clientURL - URL for the Application (e.g. "example.my.idaptive.app")
//	clientScope - Scopes for Application
//	resourceUsername - Username for the Resource Owner
//	resourcePassword - Password for the Resource Owner
func OauthResourceOwner(clientID, clientSecret, clientAppID, clientURL, resourceUsername, resourcePassword string, clientScope []string) (*oauth2.Token, error) {
	endpoint := oauth2.Endpoint{
		AuthURL:   "https://" + clientURL + "/oauth2/authorize/" + clientAppID,
		TokenURL:  "https://" + clientURL + "/oauth2/token/" + clientAppID,
		AuthStyle: 0,
	}

	// Establish oauth2/clientcredentials config with user provided data
	var resourceOwnerConfig = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		Scopes:       clientScope,
	}

	ctx := context.Background()

	authToken, err := resourceOwnerConfig.PasswordCredentialsToken(ctx, resourceUsername, resourcePassword)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain SCIM Oauth2 Token %w", err)
	}

	return authToken, nil
}
