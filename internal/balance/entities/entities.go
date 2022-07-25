package entities

import (
	"fmt"
)

type Id int
type Money int

type TransactionStatus int

const (
	tsDone TransactionStatus = iota
	tsProcessing
	tsFailed
)

var (
	ts2string = map[TransactionStatus]string{
		tsDone:       "Done",
		tsProcessing: "Processing",
		tsFailed:     "Failed",
	}
)

type Transaction struct {
	FromId  Id     `json:"from_id"`
	ToId    Id     `json:"to_id"`
	Amount  Money  `json:"amount"`
	Comment string `json:"comment"`

	Status TransactionStatus `json:"status"`
}

type SystemTransaction struct {
	Status  TransactionStatus `json:"status"`
	Id      Id                `json:"id"`
	Amount  Money             `json:"amount"`
	Comment string            `json:"comment"`
}

type Account struct {
	Id      Id    `json:"id"`
	Balance Money `json:"balance"`
}

func (m Money) String() string {
	return fmt.Sprintf("%d.%d", m/100, m%100)
}

func (t TransactionStatus) MarshalJSON() ([]byte, error) {
	return []byte("\"" + ts2string[t] + "\""), nil
}

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}
