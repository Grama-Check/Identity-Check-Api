package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func newStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Exec takes in a function and persons the transactions specified making sure that its rolled back if theres an error
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return err

	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err:%v,rb err:%v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// Checks if person in db

func (store *Store) CheckPersonTx(ctx context.Context, id string) (bool, error) {
	err := store.execTx(ctx, func(q *Queries) error {
		_, err := q.GetPerson(ctx, id)

		if err != nil {
			return err

		}
		return err

	})
	if err == nil {
		return true, err
	}
	return false, err
}
