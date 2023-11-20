package types

// error response
type Error struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}
