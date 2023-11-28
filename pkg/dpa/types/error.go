package types

// error response
type ErrorResponse struct {
	Code        string                `json:"code,omitempty"`
	Message     string                `json:"message,omitempty"`
	Description string                `json:"description,omitempty"`
	Steps       string                `json:"steps,omitempty"`
	Doc         string                `json:"doc,omitempty"`
	Errors      []NestedErrorResponse `json:"errors,omitempty"`
}

type NestedErrorResponse struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Field       string `json:"field,omitempty"`
}
