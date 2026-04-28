package api

import (
	"net/http"

	useApi "github.com/caioleone/go-grimoire-api/api/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type SpellHandler struct {
	useCase *useApi.SpellUsecase
}

func NewHandlerSpell(u *useApi.SpellUsecase) http.Handler {
	h := &SpellHandler{
		useCase: u,
	}

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	//Spell Route
	r.Route("/", func(r chi.Router) {
		r.Post("/", h.handleCreateSpell)
		r.Get("/", h.handleGetSpell)
		r.Get("/{id}", h.handleGetByIdSpell)
		//r.Get("/element/{element}", h.handleGetByElement)
		r.Put("/{id}", h.handleUpdateSpell)
		r.Delete("/{id}", h.handleDeleteSpell)
	})

	return r
}
