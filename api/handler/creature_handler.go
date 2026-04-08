package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
)

func NewHandlerCreature(repoCreature repoApi.CreatureRepository) http.Handler {
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

func handleCreateCreature(repoCreature repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creature domApi.CreatureModel

		if err := json.NewDecoder(r.Body).Decode(&creature); err != nil {
			SendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
			return
		}

		//VALIDACAO
		if len(creature.Name) < 2 || len(creature.Description) < 2 || creature.Attack <= 0 || creature.Defence <= 0 || creature.Hp <= 0 {
			SendJson(w, Response{Error: "Invalid Fields"}, http.StatusBadRequest)
			return
		}

		created := repoCreature.Insert(creature)
		SendJson(w, Response{Data: created}, http.StatusCreated)
	}
}

func handleGetAllCreature(repoCreature repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		creature := repoCreature.FindAll()
		SendJson(w, Response{Data: creature}, http.StatusOK)
	}
}

func handleGetByIdCreature(repoCreature repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
			return
		}

		creature, ok := repoCreature.FindById(id)
		if !ok {
			SendJson(w, Response{Error: "Creature Not Found"}, http.StatusNotFound)
		}
		SendJson(w, Response{Data: creature}, http.StatusOK)
	}
}

func handleUpdateCreature(repoCreature repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
			return
		}

		var updated domApi.CreatureModel
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			SendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
			return
		}

		result, ok := repoCreature.Update(id, updated)
		if !ok {
			SendJson(w, Response{Error: "Creature Not Found"}, http.StatusNotFound)
			return
		}

		SendJson(w, Response{Data: result}, http.StatusOK)
	}
}

func handleDeleteCreature(repoCreature repoApi.CreatureRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
			return
		}

		deleted, ok := repoCreature.Delete(id)
		if !ok {
			SendJson(w, Response{Error: "Creature Not Found"}, http.StatusNotFound)
			return
		}

		SendJson(w, Response{Data: deleted}, http.StatusOK)
	}
}
