package server

import (
	. "balance/internal/balance/entities"
	"encoding/json"
	"fmt"
	"io"
)

func parseAccount(stream io.Reader, account *Account) error {
	jsonDecoder := json.NewDecoder(stream)
	if err := jsonDecoder.Decode(account); err != nil {
		return err
	}

	if err := account.Validate(); err != nil {
		return err
	}

	return nil
}

func parseTransaction(stream io.Reader, transaction *Transaction) error {
	jsonDecoder := json.NewDecoder(stream)

	if err := jsonDecoder.Decode(&transaction); err != nil {
		return err
	}

	if err := transaction.Validate(); err != nil {
		return err
	}

	if transaction.Comment == "" {
		transaction.Comment = fmt.Sprintf("Перевод со счета %d на счет %d %s рублей",
			transaction.FromId,
			transaction.ToId,
			transaction.Amount.String())
	}

	return nil
}

func parseSystemTransaction(stream io.Reader, transaction *SystemTransaction) error {
	jsonDecoder := json.NewDecoder(stream)

	if err := jsonDecoder.Decode(&transaction); err != nil {
		if err != io.EOF {
			return err
		}
	}

	if err := transaction.Validate(); err != nil {
		return err
	}

	if transaction.Comment == "" {
		if transaction.Amount < 0 {
			transaction.Comment = fmt.Sprintf("Списание с баланса %s рублей", transaction.Amount.String())
		} else {
			transaction.Comment = fmt.Sprintf("Пополнение баланса на %s рублей", transaction.Amount.String())
		}
	}

	return nil
}
