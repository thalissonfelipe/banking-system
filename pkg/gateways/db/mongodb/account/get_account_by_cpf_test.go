package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
	"github.com/thalissonfelipe/banking/pkg/tests/dockertest"
	"github.com/thalissonfelipe/banking/pkg/tests/testdata"
)

func TestRepostory_GetAccountByCPF(t *testing.T) {
	r := NewRepository(db)

	defer dockertest.DropCollection(t, db.Collection("accounts"))

	account, err := r.GetAccountByCPF(context.Background(), testdata.GetValidCPF())
	assert.Nil(t, account)
	assert.ErrorIs(t, err, entities.ErrAccountDoesNotExist)

	acc := dockertest.CreateAccount(t, db.Collection("accounts"), 100)

	account, err = r.GetAccountByCPF(context.Background(), acc.CPF)
	assert.NoError(t, err)

	assert.Equal(t, acc.ID, account.ID)
	assert.Equal(t, acc.Name, account.Name)
	assert.Equal(t, acc.CPF, account.CPF)
	assert.Equal(t, acc.Balance, account.Balance)
	assert.NotEmpty(t, account.Secret)
}
