package server

import "errors"

var (
	ErrInvalidId     error = errors.New("Invalid 'id'")
	ErrInvalidFromId error = errors.New("Invalid 'from_id'")
	ErrInvalidToId   error = errors.New("Invalid 'to_id'")
	ErrInvalidAmount error = errors.New("Invalid 'amount'")

	ErrEmptyId     error = errors.New("Empty 'id'")
	ErrEmptyFromId error = errors.New("Empty 'from_id'")
	ErrEmptyToId   error = errors.New("Empty 'to_id'")
	ErrEmptyAmount error = errors.New("Empty 'amount'")

	ErrDbFault error = errors.New("Error happend while DB work")
)
