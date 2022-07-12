package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/thalissonfelipe/banking/banking/domain/entities"
	"github.com/thalissonfelipe/banking/banking/domain/vos"
	"github.com/thalissonfelipe/banking/banking/tests/dockertest"
	"github.com/thalissonfelipe/banking/banking/tests/testdata"
)

func TestAccountRepository_GetBalanceByID(t *testing.T) {
	t.Run("should get balance successfully", func(t *testing.T) {
		db := pgDocker.DB
		r := NewRepository(db)
		ctx := context.Background()

		defer dockertest.TruncateTables(ctx, db)

		wantBalance := 100

		acc := entities.NewAccount("name", testdata.GetValidCPF(), testdata.GetValidSecret())
		acc.Balance = wantBalance

		err := r.CreateAccount(ctx, &acc)
		require.NoError(t, err)

		balance, err := r.GetBalanceByID(ctx, acc.ID)
		require.NoError(t, err)
		assert.Equal(t, wantBalance, balance)
	})

	t.Run("should return an error if account does not exist", func(t *testing.T) {
		db := pgDocker.DB
		r := NewRepository(db)
		ctx := context.Background()

		balance, err := r.GetBalanceByID(ctx, vos.NewAccountID())
		assert.ErrorIs(t, err, entities.ErrAccountDoesNotExist)
		assert.Zero(t, balance)
	})
}