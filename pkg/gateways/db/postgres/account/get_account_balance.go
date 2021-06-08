package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
	"github.com/thalissonfelipe/banking/pkg/domain/vos"
)

const getBalanceQuery = `
select balance from accounts where id=$1
`

func (r Repository) GetBalanceByID(ctx context.Context, id vos.AccountID) (int, error) {
	var balance int

	err := r.db.QueryRow(ctx, getBalanceQuery, id).Scan(&balance)
	if err == nil {
		return balance, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return 0, entities.ErrAccountDoesNotExist
	}

	return 0, fmt.Errorf("unexpected error occurred on get balance query: %w", err)
}
