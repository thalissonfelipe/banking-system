package transfer

import (
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/thalissonfelipe/banking/banking/domain/usecases"
	"github.com/thalissonfelipe/banking/banking/gateway/http/middlewares"
)

//go:generate moq -pkg transfer -skip-ensure -out usecase_mock.gen.go ../../../domain/usecases Transfer:UsecaseMock

type Handler struct {
	usecase usecases.Transfer
}

func NewHandler(r chi.Router, usecase usecases.Transfer) *Handler {
	handler := Handler{usecase: usecase}

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthorizeMiddleware)
		r.Route("/transfers", func(r chi.Router) {
			r.Get("/", handler.ListTransfers)
			r.Post("/", handler.PerformTransfer)
		})
	})

	return &handler
}

func getTokenFromHeader(authHeader string) string {
	return strings.Split(authHeader, "Bearer ")[1]
}
