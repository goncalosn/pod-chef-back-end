package errors

import "net/http"

// Error defines a standard application error.
type Error struct {
	// Machine-readable error code.
	Code int `json:"-"`

	// Human-readable message.
	Message string `json:"Message"`

	// Logical operation and nested error.
	Op  string `json:"-"`
	Err error  `json:"-"`
}

func (err *Error) Error() string {
	return err.Err.Error()
}

// ErrorCode returns the code of the root error, if available. Otherwise returns EINTERNAL.
func ErrorCode(err error) int {
	if err == nil {
		return http.StatusOK
	} else if e, ok := err.(*Error); ok && e.Code != http.StatusOK {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return http.StatusInternalServerError
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred."
}
