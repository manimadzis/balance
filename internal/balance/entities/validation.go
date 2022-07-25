package entities

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	mimCommentLength        = 1
	maxCommentLength        = 60
	minAmountForTransaction = Money(1)
)

func (a *Account) Validate() error {
	return validation.ValidateStruct(
		validation.Field(&a.Id, validation.Required),
	)
}

func (t *SystemTransaction) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Id, validation.Required),
		validation.Field(&t.Amount, validation.Required),
		validation.Field(&t.Comment, validation.Length(mimCommentLength, maxCommentLength)),
	)
}

func (t *Transaction) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.FromId, validation.Required),
		validation.Field(&t.ToId, validation.Required),
		validation.Field(&t.Amount, validation.Required, validation.Min(minAmountForTransaction)),
		validation.Field(&t.Comment, validation.Length(mimCommentLength, maxCommentLength)),
	)
}
