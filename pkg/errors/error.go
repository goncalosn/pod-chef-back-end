package errors

import (
	"encoding/json"
	"fmt"
)

type HTTPError struct {
	Cause   error  `json:"-"`
	Message string `json:"Message"`
	Status  int    `json:"-"`
}

func (e *HTTPError) Error() string {
	return e.Cause.Error()
}

func NewHTTPError(err error, status int, message string) error {
	return &HTTPError{
		Cause:   err,
		Message: message,
		Status:  status,
	}
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

// GetStatus returns the error status
func (e *HTTPError) GetStatus() int {
	return e.Status
}

// GetMessage returns the error message
func (e *HTTPError) GetMessage() string {
	return e.Message
}
