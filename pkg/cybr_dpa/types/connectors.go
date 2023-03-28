package types

type ConnectorScript struct {
	ScriptUrl   string `json:"script_url"`
	BashCommand string `json:"bash_cmd"`
}

type ConnectorRequest struct {
	ConnectorType string `json:"connector_type"`
	ConnectorOS   string `json:"connector_os"`
}
