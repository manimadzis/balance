package store

import "errors"

var (
	ErrUnknownId      error = errors.New("Unknown id")
	ErrNonexistentId  error = errors.New("Account with given id doesn't exist")
	ErrNotEnoughMoney error = errors.New("Account hasn't enough money")
)
