package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
)

func (h *SpellHandler) handleCreateSpell(w http.ResponseWriter, r *http.Request) {
	var spell domApi.SpellModel

	if err := json.NewDecoder(r.Body).Decode(&spell); err != nil {
		SendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
		return
	}

	created, err := h.useCase.CreateSpell(spell)
	if err != nil {
		SendJson(w, Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	SendJson(w, Response{Data: created}, http.StatusCreated)

}

func (h *SpellHandler) handleGetSpell(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	element := query.Get("element")
	sort := query.Get("sort")
	order := query.Get("order")
	//pageStr := query.Get("page")
	//limitStr := query.Get("limit")

	// page, _ := strconv.Atoi(pageStr)
	// limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))

	spells := h.useCase.GetWithFilters(element, sort, order, page, limit)

	//spells := h.useCase.GetAll()
	SendJson(w, Response{Data: spells}, http.StatusOK)

}

func (h *SpellHandler) handleGetByIdSpell(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Id"}, http.StatusBadRequest)
		return
	}

	spell, ok := h.useCase.GetById(id)
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

	result, ok := h.useCase.Update(id, updated)
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

	deleted, ok := h.useCase.Delete(id)
	if !ok {
		SendJson(w, Response{Error: "User Not Found"}, http.StatusNotFound)
		return
	}

	SendJson(w, Response{Data: deleted}, http.StatusOK)

}

func (h *SpellHandler) handleGetByElement(w http.ResponseWriter, r *http.Request) {
	element := chi.URLParam(r, "element")

	spells, err := h.useCase.GetByElement(element)
	if err != nil {
		SendJson(w, Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	SendJson(w, Response{Data: spells}, http.StatusOK)
}
