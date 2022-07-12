package usecase

import (
	"github.com/thalissonfelipe/banking/banking/domain/account"
	"github.com/thalissonfelipe/banking/banking/domain/transfer"
)

type Transfer struct {
	repository     transfer.Repository
	accountUsecase account.Usecase
}

func NewTransferUsecase(repo transfer.Repository, accUsecase account.Usecase) *Transfer {
	return &Transfer{
		repository:     repo,
		accountUsecase: accUsecase,
	}
}