package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
)

/*
// func NewHandlerSpell(repoSpell repoApi.SpellRepository) http.Handler {
// 	r := chi.NewMux()

// 	r.Use(middleware.Recoverer)
// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.Logger)

// 	//Spell Route
// 	r.Route("/api/spell", func(r chi.Router) {
// 		//POST
// 		r.Post("/", handleCreateSpell(repoSpell))

// 		//GET
// 		r.Get("/", handleGetSpell(repoSpell))
// 		r.Get("/{id}", handleGetByIdSpell(repoSpell))

// 		//PUT
// 		r.Put("/{id}", handleUpdateSpell(repoSpell))

// 		r.Delete("{id}", handleDeleteSpell(repoSpell))
// 	})

// 	return r
// }
*/

func (h *SpellHandler) handleCreateSpell(w http.ResponseWriter, r *http.Request) {
	var spell domApi.SpellModel

	if err := json.NewDecoder(r.Body).Decode(&spell); err != nil {
		SendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
		return
	}

	if len(spell.Name) < 2 || len(spell.Description) < 2 || len(spell.Element) < 2 || spell.ManaCost <= 0 {
		SendJson(w, Response{Error: "Invalid Fields"}, http.StatusBadRequest)
		return
	}

	created := h.repo.Insert(spell)
	SendJson(w, Response{Data: created}, http.StatusCreated)

}

func (h *SpellHandler) handleGetSpell(w http.ResponseWriter, r *http.Request) {

	spell := h.repo.FindAll()
	SendJson(w, Response{Data: spell}, http.StatusOK)

}

func (h *SpellHandler) handleGetByIdSpell(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
		return
	}

	spell, ok := h.repo.FindById(id)
	if !ok {
		SendJson(w, Response{Error: "Spell Not Found"}, http.StatusBadRequest)
	}

	SendJson(w, Response{Data: spell}, http.StatusOK)

}

func (h *SpellHandler) handleUpdateSpell(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
		return
	}

	var updated domApi.SpellModel
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		SendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
		return
	}

	result, ok := h.repo.Update(id, updated)
	if !ok {
		SendJson(w, Response{Error: "Spell Not Found"}, http.StatusNotFound)
		return
	}

	SendJson(w, Response{Data: result}, http.StatusOK)

}

func (h *SpellHandler) handleDeleteSpell(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
		return
	}

	deleted, ok := h.repo.Delete(id)
	if !ok {
		SendJson(w, Response{Error: "User Not Found"}, http.StatusNotFound)
		return
	}

	SendJson(w, Response{Data: deleted}, http.StatusOK)

}
