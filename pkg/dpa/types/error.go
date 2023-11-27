package types

// error response
type ErrorResponse struct {
	Code        string                `json:"code,omitempty"`
	Message     string                `json:"message,omitempty"`
	Description string                `json:"description,omitempty"`
	Errors      []NestedErrorResponse `json:"errors,omitempty"`
}

type NestedErrorResponse struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Field       string `json:"field,omitempty"`
}
