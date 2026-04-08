package api

import (
	"net/http"

	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type SpellHandler struct {
	repo repoApi.SpellRepository
}

func NewHandlerSpell(repo repoApi.SpellRepository) http.Handler {
	h := &SpellHandler{
		repo: repo,
	}
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	//Spell Route
	r.Route("/api/spell", func(r chi.Router) {
		r.Post("/", h.handleCreateSpell)
		r.Get("/", h.handleGetSpell)
		r.Get("/{id}", h.handleGetByIdSpell)
		r.Put("/{id}", h.handleUpdateSpell)
		r.Delete("{id}", h.handleDeleteSpell)
	})

	return r
}
