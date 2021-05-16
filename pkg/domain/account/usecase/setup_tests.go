package usecase

import (
	"context"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
)

type StubRepository struct {
	accounts []entities.Account
	err      error
}

func (s StubRepository) GetAccounts(ctx context.Context) ([]entities.Account, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.accounts, nil
}

func (s StubRepository) GetBalanceByID(ctx context.Context, id string) (int, error) {
	for _, account := range s.accounts {
		if account.ID == id {
			return account.Balance, nil
		}
	}
	return 0, entities.ErrAccountDoesNotExist
}