package mocks

import (
	"context"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
)

type StubTransferRepository struct {
	Transfers []entities.Transfer
	Err       error
}

func (s StubTransferRepository) GetTransfers(ctx context.Context, id string) ([]entities.Transfer, error) {
	if s.Err != nil {
		return nil, entities.ErrInternalError
	}

	var transfers []entities.Transfer
	for _, tr := range s.Transfers {
		if tr.AccountOriginID == id {
			transfers = append(transfers, tr)
		}
	}

	return transfers, nil
}

func (s *StubTransferRepository) UpdateBalance(ctx context.Context, transfer entities.Transfer) error {
	s.Transfers = append(s.Transfers, transfer)
	return nil
}