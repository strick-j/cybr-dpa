package dpa

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type clientCredsTokenSource struct {
	ctx  context.Context
	conf *clientcredentials.Config
}

// OauthCredClient returns a validated Oauth2 Authentication Token based on the following provided information:
//
//	clientID - Username for the Application (e.g. "identity-privilege-integration-user$@example.com")
//	clientSecret - Password for the Application Service User
//	clientAppID - Application ID for the Oauth2 Application
//	clientURL - URL for the Application (e.g. "example.cyberark.cloud")
//	scope - Scope for the application (e.g. "dpa")
//
// Returns an oauth2.Token or error
func OauthCredClient(clientID, clientSecret, clientAppID, clientURL string, scope []string) (*oauth2.Token, error) {
	// Establish oauth2/clientcredentials config with user provided data
	var credentialConfig = clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://" + clientURL + "/oauth2/token/" + clientAppID,
		AuthStyle:    0,
		Scopes:       scope,
	}

	// Create tokenSource with provided configuration info
	ts := &clientCredsTokenSource{
		ctx:  context.Background(),
		conf: &credentialConfig,
	}

	// Request new token from server using Client Credentials
	authToken, err := ts.conf.Token(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to obtain Oauth2 Token %w", err)
	}

	return authToken, nil
}

// OauthPlatformToken returns a validated Oauth2 Authentication Token based on the following provided information:
//
//	clientID - Username for the Application (e.g. "identity-privilege-integration-user$@example.com")
//	clientSecret - Password for the Application Service User
//	clientURL - URL for the Application (e.g. "example.cyberark.cloud")
//
// Returns an oauth2.Token or error
func OauthPlatformToken(clientID, clientSecret, clientURL string) (*oauth2.Token, error) {
	// Establish oauth2/clientcredentials config with user provided data
	var credentialConfig = clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://" + clientURL + "/oauth2/platformtoken",
		AuthStyle:    0,
	}

	// Create tokenSource with provided configuration info
	ts := &clientCredsTokenSource{
		ctx:  context.Background(),
		conf: &credentialConfig,
	}

	// Request new token from server using Client Credentials
	authToken, err := ts.conf.Token(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to obtain Oauth2 Token %w", err)
	}

	return authToken, nil
}
