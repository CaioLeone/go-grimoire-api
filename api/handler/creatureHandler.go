package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
)


func (h *CreatureHandler) handleCreateCreature(w http.ResponseWriter, r *http.Request) {

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

	created := h.repo.Insert(creature)
	SendJson(w, Response{Data: created}, http.StatusCreated)

}

func (h *CreatureHandler) handleGetAllCreature(w http.ResponseWriter, r *http.Request) {
	creature := h.repo.FindAll()
	SendJson(w, Response{Data: creature}, http.StatusOK)

}

func (h *CreatureHandler) handleGetByIdCreature(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
		return
	}

	creature, ok := h.repo.FindById(id)
	if !ok {
		SendJson(w, Response{Error: "Creature Not Found"}, http.StatusNotFound)
	}
	SendJson(w, Response{Data: creature}, http.StatusOK)

}

func (h *CreatureHandler) handleUpdateCreature(w http.ResponseWriter, r *http.Request) {
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

	result, ok := h.repo.Update(id, updated)
	if !ok {
		SendJson(w, Response{Error: "Creature Not Found"}, http.StatusNotFound)
		return
	}

	SendJson(w, Response{Data: result}, http.StatusOK)

}

func (h *CreatureHandler) handleDeleteCreature(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
		return
	}

	deleted, ok := h.repo.Delete(id)
	if !ok {
		SendJson(w, Response{Error: "Creature Not Found"}, http.StatusNotFound)
		return
	}

	SendJson(w, Response{Data: deleted}, http.StatusOK)

}
