package kerrors

import "fmt"

// BaseError defines a common structure for all custom errors
type JoinedError struct {
	Msg          string
	WrappedError error
}

type ValueError struct{ JoinedError }
type SystemError struct{ JoinedError }
type NetworkError struct{ JoinedError }

func (joinedErrorInstance JoinedError) Error() string {
	if joinedErrorInstance.WrappedError != nil {
		return fmt.Sprintf("%s: %v", joinedErrorInstance.Msg, joinedErrorInstance.WrappedError)
	}
	return joinedErrorInstance.Msg
}

func (joinedErrorInstance *JoinedError) Unwrap() error {
	return joinedErrorInstance.WrappedError
}

// Constructors
func NewValueError(msg string, err error) error {
	return &ValueError{
		JoinedError{
			Msg:          msg,
			WrappedError: err,
		},
	}
}

func NewSystemError(msg string, err error) error {
	return &SystemError{
		JoinedError{
			Msg:          msg,
			WrappedError: err,
		},
	}
}

func NewNetworkError(msg string, err error) error {
	return &NetworkError{
		JoinedError{
			Msg:          msg,
			WrappedError: err,
		},
	}
}
