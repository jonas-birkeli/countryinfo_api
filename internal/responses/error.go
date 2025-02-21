package responses

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// isResponse represents the response for errors
func (r ErrorResponse) isResponse() {}
