package clickup

import "errors"

type Error struct {
	StatusCode int
	Message    string
}

func (e Error) Error() string {
	return e.Message
}

func NewError(statusCode int, message string) error {
	return Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

// ErrInvalidRequestSignature is an error that should be returned if a Webhook event request signature is not valid
var ErrInvalidRequestSignature = errors.New("invalid request signature")
