package repository

import "errors"

var (
	ErrNoRowsAffected = errors.New("no rows affected")
	ErrNoRowsFound    = errors.New("no rows")
)
