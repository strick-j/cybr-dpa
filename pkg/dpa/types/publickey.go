package types

// PublicKey response from retrieving a public key
type PublicKey struct {
	PublicKey string `json:"publicKey"`
}

// PublicKeyScript response from retrieving a public key install script
type PublicKeyScript struct {
	Base64Cmd string `json:"base64_cmd,omitempty"`
}
