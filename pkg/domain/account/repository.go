package account

import (
	"context"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
)

type Repository interface {
	GetAccounts(ctx context.Context) ([]entities.Account, error)
	GetBalanceByID(ctx context.Context, id string) (int, error)
}