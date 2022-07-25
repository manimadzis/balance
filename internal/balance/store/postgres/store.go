package postgres

import (
	. "balance/internal/balance/entities"
	"balance/internal/balance/store"
	"context"
	"database/sql"
	"log"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) store.AccountStore {
	store := &Store{
		db: db,
	}

	if err := store.initDB(); err != nil {
		log.Fatal(err)
		panic("Cannot init DB")
	}

	return store
}

func (s *Store) initDB() error {
	if _, err := s.db.Exec(
		`CREATE TABLE IF NOT EXISTS account(
			id BIGINT PRIMARY KEY,
			balance BIGINT NOT NULL);
	`); err != nil {
		return err
	}

	if _, err := s.db.Exec(
		`CREATE TABLE IF NOT EXISTS transaction(
			id BIGSERIAL,
			date TIMESTAMP NOT NULL,
			from_id BIGINT,
			to_id BIGINT NOT NULL,
			amount BIGINT NOT NULL,
			comment VARCHAR(60)
		)
	`); err != nil {
		return err
	}

	return nil
}

func (s *Store) applyTransaction(transaction Transaction) error {
	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if _, err := tx.Exec(
		`UPDATE account
		SET balance = balance + $2
		WHERE id = $1`,
		transaction.ToId,
		transaction.Amount,
	); err != nil {
		return err
	}

	if _, err = tx.Exec(
		`UPDATE account
		SET balance = balance - $2
		WHERE id = $1`,
		transaction.FromId,
		transaction.Amount,
	); err != nil {
		return nil
	}

	if _, err := tx.Exec(
		`INSERT INTO transaction(
			date, 
			from_id,
			to_id, 
			amount,
			comment
		)
		VALUES (
			CURRENT_TIMESTAMP,
			$1, $2, $3, $4
		)`,
		transaction.FromId,
		transaction.ToId,
		transaction.Amount,
		transaction.Comment,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Store) getBalance(id Id) (Money, error) {
	result := s.db.QueryRow(
		`SELECT balance
		FROM account
		WHERE id = $1`,
		id,
	)

	var money int
	if err := result.Scan(&money); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrNonexistentId
		}
		return 0, err
	}

	return Money(money), nil
}

func (s *Store) SaveTransaction(transaction *Transaction) error {
	return s.applyTransaction(*transaction)
}

func (s *Store) GetBalance(account *Account) error {
	balance, err := s.getBalance(account.Id)
	if err != nil {
		return err
	}

	account.Balance = balance

	return nil
}

func (s *Store) applySystemTransaction(transaction SystemTransaction) error {
	if transaction.Amount < 0 {
		balance, err := s.getBalance(transaction.Id)

		if err == sql.ErrNoRows {
			return store.ErrUnknownId
		} else if err != nil {
			return err
		}

		if balance+Money(transaction.Amount) < 0 {
			return store.ErrNotEnoughMoney
		}
	}

	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		`UPDATE account
		SET balance = balance + $2
		WHERE id = $1`,
		transaction.Id,
		transaction.Amount,
	)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		if _, err := s.db.Exec(
			`INSERT INTO account
			VALUES ($1, $2);`,
			transaction.Id,
			transaction.Amount,
		); err != nil {
			return err
		}
	}

	if _, err := tx.Exec(
		`INSERT INTO transaction(
			date, 
			to_id, 
			amount,
			comment
		)
		VALUES (
			CURRENT_TIMESTAMP,
			$1, $2, $3
		)`,
		transaction.Id,
		transaction.Amount,
		transaction.Comment,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Store) SaveSystemTransaction(transaction *SystemTransaction) error {
	return s.applySystemTransaction(*transaction)
}
