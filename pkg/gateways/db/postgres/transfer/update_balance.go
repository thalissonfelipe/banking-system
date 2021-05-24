package transfer

import (
	"context"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
	"github.com/thalissonfelipe/banking/pkg/domain/vos"
)

func (r Repository) UpdateBalance(ctx context.Context, transfer *entities.Transfer) error {
	// First experience with rollback.
	// Tutorial: https://www.sohamkamani.com/golang/sql-transactions/
	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.WithError(err).Error("unexpected error ocurred on begin starting transaction")
		return err
	}

	defer tx.Rollback(ctx)

	err = r.updateBalance(ctx, tx, -transfer.Amount, transfer.AccountOriginID)
	if err != nil {
		log.WithError(err).Error("unexpected error ocurred while updating account origin balance: doing a rollback")
		return err
	}

	err = r.updateBalance(ctx, tx, transfer.Amount, transfer.AccountDestinationID)
	log.WithError(err).Error("unexpected error ocurred while updating account destination balance: doing a rollback")
	if err != nil {
		return err
	}

	err = r.saveTransfer(ctx, tx, transfer)
	log.WithError(err).Error("unable to save transfer: doing a rollback")
	if err != nil {
		return err
	}

	tx.Commit(ctx)

	return nil
}

func (r Repository) updateBalance(ctx context.Context, tx pgx.Tx, balance int, id vos.ID) error {
	const query = `
		UPDATE accounts
		SET balance=balance+$1
		WHERE id=$2
	`

	_, err := tx.Exec(ctx, query, balance, id)
	return err
}

func (r Repository) saveTransfer(ctx context.Context, tx pgx.Tx, transfer *entities.Transfer) error {
	const query = `
		INSERT INTO transfers (
			id,
			account_origin_id,
			account_destination_id,
			amount
		) VALUES (
			$1, $2, $3, $4
		) RETURNING created_at
	`

	err := tx.QueryRow(ctx, query,
		transfer.ID,
		transfer.AccountOriginID,
		transfer.AccountDestinationID,
		transfer.Amount,
	).Scan(
		&transfer.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
