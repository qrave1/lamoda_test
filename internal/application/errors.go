package application

import (
	"errors"
)

var (
	ErrBeginTx = errors.New("error begin transaction")
)

type ApplicationError struct {
	msg string
}

func (ae ApplicationError) Error() string {
	return ae.msg
}

func NewApplicationError(msg string) error {
	return ApplicationError{msg: msg}
}
