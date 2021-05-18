package account

import (
	"context"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
)

func (r Repository) GetAccounts(ctx context.Context) ([]entities.Account, error) {
	query := "SELECT id, name, cpf, balance, created_at FROM account"

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]entities.Account, 0)

	for rows.Next() {
		var account entities.Account
		err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.CPF,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
