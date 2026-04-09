package api

import (
	"net/http"

	useApi "github.com/caioleone/go-grimoire-api/api/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CreatureHandler struct {
	useCase *useApi.CreatureUsecase
}

func NewHandlerCreature(u *useApi.CreatureUsecase) http.Handler {
	h := &CreatureHandler{
		useCase: u,
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
