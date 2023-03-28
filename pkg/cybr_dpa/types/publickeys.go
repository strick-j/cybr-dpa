package types

type PublicKey struct {
	PublicKey   string `json:"publickey,omitempty"`
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}
