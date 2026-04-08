package api

import (
	"net/http"

	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CreatureHandler struct {
	repo repoApi.CreatureRepository
}

func NewHandlerCreature(repo repoApi.CreatureRepository) http.Handler {
	h := &CreatureHandler{
		repo: repo,
	}

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	//Creature Route
	r.Route("/api/creature", func(r chi.Router) {
		r.Post("/", h.handleCreateCreature)
		r.Get("/", h.handleGetAllCreature)
		r.Get("/{id}", h.handleGetByIdCreature)
		r.Put("{id}", h.handleUpdateCreature)
		r.Delete("/{id}", h.handleDeleteCreature)
	})
	return r
}