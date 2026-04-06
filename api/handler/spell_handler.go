package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
)

type PostBodySpell struct {
	URL string `json:url`
}

type ResponseSpell struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandlerSpell(repoSpell *repoApi.SpellRepository) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	//Spell Route
	r.Route("/api/spell", func(r chi.Router) {
		//POST
		r.Post("/", handleCreateSpell(repoSpell))

		//GET
		r.Get("/", handleGetSpell(repoSpell))
		r.Get("/{id}", handleGetByIdSpell(repoSpell))

		//PUT
		r.Put("/{id}", handleUpdateSpell(repoSpell))

		r.Delete("{id}", handleDeleteSpell(repoSpell))
	})

	return r
}

func handleCreateSpell(repo *repoApi.SpellRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var spell domApi.SpellModel

		if err := json.NewDecoder(r.Body).Decode(&spell); err != nil {
			return
		}
	}
}

func handleGetSpell(repo *repoApi.SpellRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func handleGetByIdSpell(repo *repoApi.SpellRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func handleUpdateSpell(repo *repoApi.SpellRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func handleDeleteSpell(repo *repoApi.SpellRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
