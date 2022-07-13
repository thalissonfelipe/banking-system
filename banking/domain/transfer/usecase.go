package transfer

import (
	"context"

	"github.com/thalissonfelipe/banking/banking/domain/entities"
	"github.com/thalissonfelipe/banking/banking/domain/vos"
)

type Usecase interface {
	ListTransfers(ctx context.Context, accountID vos.AccountID) ([]entities.Transfer, error)
	CreateTransfer(ctx context.Context, input CreateTransferInput) error
}

type CreateTransferInput struct {
	AccountOriginID      vos.AccountID
	AccountDestinationID vos.AccountID
	Amount               int
}

func NewTransferInput(accOriginID, accDestID vos.AccountID, amount int) CreateTransferInput {
	return CreateTransferInput{
		AccountOriginID:      accOriginID,
		AccountDestinationID: accDestID,
		Amount:               amount,
	}
}
