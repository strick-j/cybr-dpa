package types

// ListTargetSetResponse is the struct response provided when listing target sets
type ListTargetSetResponse struct {
	TargetSets          []TargetSets `json:"target_sets,omitempty"`
	B64LastEvaluatedKey string       `json:"b64_last_evaluated_key,omitempty"`
}
type TargetSets struct {
	Name                        string `json:"name,omitempty"`
	Description                 string `json:"description,omitempty"`
	ProvisionFormat             string `json:"provision_format,omitempty"`
	EnableCertificateValidation bool   `json:"enable_certificate_validation,omitempty"`
	SecretType                  string `json:"secret_type,omitempty"`
	SecretID                    string `json:"secret_id,omitempty"`
	Type                        string `json:"type,omitempty"`
}

// DeleteTargetSetResponse is the struct response provided when deleting a target set
type DeleteTargetSetResponse struct {
	Results []Results `json:"results,omitempty"`
}

// TargetSetMapping is the struct format utilized to post a target set to the API
type TargetSetMapping struct {
	StrongAccountID string       `json:"strong_account_id,omitempty"`
	TargetSets      []TargetSets `json:"target_sets,omitempty"`
}

// Results is the struct response provided when adding a target set
type AddTargetSetResponse struct {
	Results []Results `json:"results,omitempty"`
}
type Results struct {
	StrongAccountID string `json:"strong_account_id,omitempty"`
	TargetSetName   string `json:"target_set_name,omitempty"`
	Success         bool   `json:"success,omitempty"`
}
