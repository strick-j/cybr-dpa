package types

// GenerateScriptResponse response from generating a script
type GenerateScriptResponse struct {
	ScriptURL string `json:"script_url,omitempty"`
	BashCmd   string `json:"bash_cmd,omitempty"`
}
