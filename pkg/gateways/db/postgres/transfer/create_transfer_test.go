package transfer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
	"github.com/thalissonfelipe/banking/pkg/domain/vos"
	"github.com/thalissonfelipe/banking/pkg/gateways/db/postgres/account"
	"github.com/thalissonfelipe/banking/pkg/tests"
)

func TestRepository_CreateTransfer(t *testing.T) {
	db := pgDocker.DB
	accRepository := account.NewRepository(db)
	r := NewRepository(db)
	ctx := context.Background()

	acc1 := entities.NewAccount("Maria", tests.TestCPF1, vos.NewSecret("12345678"))
	acc1.Balance = 0
	acc2 := entities.NewAccount("Pedro", tests.TestCPF2, vos.NewSecret("12345678"))

	err := accRepository.CreateAccount(ctx, &acc1)
	assert.NoError(t, err)

	err = accRepository.CreateAccount(ctx, &acc2)
	assert.NoError(t, err)

	transfer := entities.NewTransfer(acc1.ID, acc2.ID, 50)
	assert.Empty(t, transfer.CreatedAt)

	err = r.CreateTransfer(ctx, &transfer)
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer.CreatedAt)

	transfers, err := r.GetTransfers(ctx, acc1.ID)
	assert.NoError(t, err)
	assert.Len(t, transfers, 1)

	assert.Equal(t, transfer.ID, transfers[0].ID)
	assert.Equal(t, transfer.AccountOriginID, transfers[0].AccountOriginID)
	assert.Equal(t, transfer.AccountDestinationID, transfers[0].AccountDestinationID)
	assert.Equal(t, transfer.Amount, transfers[0].Amount)
}
