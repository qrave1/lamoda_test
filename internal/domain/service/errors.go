package service

import (
	"errors"
)

var (
	ErrBeginTx = errors.New("error begin transaction")
)

type ApplicationError struct {
	msg    string
	status int
}

func (ae ApplicationError) Error() string {
	return ae.msg
}

func (ae ApplicationError) Status() int {
	return ae.status
}

func NewApplicationError(msg string, status int) error {
	return ApplicationError{msg: msg, status: status}
}
