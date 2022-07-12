package account

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/thalissonfelipe/banking/banking/domain/vos"
	"github.com/thalissonfelipe/banking/banking/gateway/http/account/schemes"
	"github.com/thalissonfelipe/banking/banking/gateway/http/rest"
)

// GetAccountBalance returns a balance by accountID
// @Tags accounts
// @Summary Get account balance
// @Description Get account balance by accountID, if exists.
// @Accept json
// @Produce json
// @Success 200 {object} balanceResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /accounts/{accountID}/balance [GET].
func (h Handler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	accountID := vos.ConvertStringToAccountID(chi.URLParam(r, "accountID"))

	balance, err := h.usecase.GetAccountBalanceByID(r.Context(), accountID)
	if err != nil {
		rest.HandleError(w, err)

		return
	}

	response := schemes.BalanceResponse{Balance: balance}
	rest.SendJSON(w, http.StatusOK, response)
}