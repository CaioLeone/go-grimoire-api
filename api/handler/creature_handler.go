package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
)

type PostBodyCreature struct {
	URL string `json:url`
}

type ResponseCreature struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandlerCreature(repoCreature *repoApi.CreatureRepository) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	//Creature Route
	r.Route("/api/creature", func(r chi.Router) {
		r.Post("/", handleCreateCreature(repoCreature))
		r.Get("/", handleGetAllCreature(repoCreature))
		r.Get("/{id}", handleGetByIdCreature(repoCreature))
		r.Put("{id}", handleUpdateCreature(repoCreature))
		r.Delete("/{id}", handleDeleteCreature(repoCreature))
	})
	return r
}

func handleCreateCreature(repoCreature *repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func handleGetAllCreature(repoCreature *repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func handleGetByIdCreature(repoCreature *repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func handleUpdateCreature(repoCreature *repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func handleDeleteCreature(repoCreature *repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
