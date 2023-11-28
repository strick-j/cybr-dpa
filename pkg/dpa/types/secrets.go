package types

// Secrets response from retrieving secrets.
type Secrets []struct {
	SecretID      string        `json:"secret_id,omitempty"`
	TenantID      string        `json:"tenant_id,omitempty"`
	SecretType    string        `json:"secret_type,omitempty"`
	SecretName    any           `json:"secret_name,omitempty"`
	SecretDetails SecretDetails `json:"secret_details,omitempty"`
	IsActive      bool          `json:"is_active,omitempty"`
}
type SecretDetails struct {
	CertFileName  string `json:"certFileName,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Domains       []any  `json:"domains,omitempty"`
	AccountDomain string `json:"account_domain,omitempty"`
}

// SingleSecret response when calling the secrets endpoint with
// a secret ID.
type SingleSecret struct {
	SecretID      string        `json:"secret_id,omitempty"`
	TenantID      string        `json:"tenant_id,omitempty"`
	Secret        Secret        `json:"secret,omitempty"`
	SecretType    string        `json:"secret_type,omitempty"`
	SecretDetails SecretDetails `json:"secret_details,omitempty"`
	IsActive      bool          `json:"is_active,omitempty"`
	IsRotatable   bool          `json:"is_rotatable,omitempty"`
	CreationTime  string        `json:"creation_time,omitempty"`
	LastModified  string        `json:"last_modified,omitempty"`
	SecretName    string        `json:"secret_name,omitempty"`
}
type Secret struct {
	SecretData      string `json:"secret_data,omitempty"`
	TenantEncrypted bool   `json:"tenant_encrypted,omitempty"`
}
