package aurora

import "fmt"

// APIError represents an error returned by the Aurora API.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("aurora: api error (status %d): %s", e.StatusCode, e.Message)
}
