package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/thalissonfelipe/banking/pkg/domain/entities"
	"github.com/thalissonfelipe/banking/pkg/gateways/http/responses"
	"github.com/thalissonfelipe/banking/pkg/services/auth"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		responses.SendError(w, http.StatusBadRequest, errInvalidJSON.Error())
		return
	}

	if err := body.isValid(); err != nil {
		responses.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	input := auth.AuthenticateInput{
		CPF:    body.CPF,
		Secret: body.Secret,
	}
	token, err := h.authService.Autheticate(r.Context(), input)
	if err != nil {
		if errors.Is(err, entities.ErrAccountDoesNotExist) {
			responses.SendError(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, auth.ErrSecretDoesNotMatch) {
			responses.SendError(w, http.StatusBadRequest, err.Error())
			return
		}

		responses.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := responseBody{Token: token}
	responses.SendJSON(w, http.StatusOK, response)
}