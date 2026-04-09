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

	created, err := h.useCase.CreateCreature(creature)
	if err != nil {
		SendJson(w, Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	SendJson(w, Response{Data: created}, http.StatusCreated)

}

func (h *CreatureHandler) handleGetAllCreature(w http.ResponseWriter, r *http.Request) {
	creature := h.useCase.GetAll()
	SendJson(w, Response{Data: creature}, http.StatusOK)

}

func (h *CreatureHandler) handleGetByIdCreature(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
		return
	}

	creature, ok := h.useCase.GetById(id)
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

	result, ok := h.useCase.Update(id, updated)
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

	deleted, ok := h.useCase.Delete(id)
	if !ok {
		SendJson(w, Response{Error: "Creature Not Found"}, http.StatusNotFound)
		return
	}

	SendJson(w, Response{Data: deleted}, http.StatusOK)

}
