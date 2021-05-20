package transfer

import (
	"context"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
)

func (r Repository) GetTransfers(ctx context.Context, id string) ([]entities.Transfer, error) {
	const query = `
		SELECT
			id,
			account_origin_id,
			account_destination_id,
			amount,
			created_at
		FROM transfer
		WHERE account_origin_id=$1
	`

	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := make([]entities.Transfer, 0)

	for rows.Next() {
		var account entities.Transfer
		err := rows.Scan(
			&account.ID,
			&account.AccountOriginID,
			&account.AccountDestinationID,
			&account.Amount,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, account)
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}
