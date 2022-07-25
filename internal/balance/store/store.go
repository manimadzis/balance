package store

import . "balance/internal/balance/entities"

type AccountStore interface {
	SaveTransaction(*Transaction) error
	SaveSystemTransaction(*SystemTransaction) error
	GetBalance(*Account) error
}
