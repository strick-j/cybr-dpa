package types

type ConnectorScript struct {
	ScriptUrl   *string `json:"script_url,omitempty"`
	BashCommand *string `json:"bash_cmd,omitempty"`
	Code        *string `json:"code,omitempty"`
	Message     *string `json:"message,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ConnectorRequest struct {
	ConnectorType string `json:"connector_type"`
	ConnectorOS   string `json:"connector_os"`
}
