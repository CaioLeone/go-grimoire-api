package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
)

// CreateSpell godoc
// @Summary Cria uma nova magia
// @Description Cria uma magia com nome, descrição, elemento e custo de mana
// @Tags spells
// @Accept json
// @Produce json
// @Param spell body domApi.SpellModel true "Dados da magia"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Router /api/spell [post]
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

// GetSpells godoc
// @Summary Lista magias com filtros
// @Tags spells
// @Produce json
// @Param name query string false "Nome da magia"
// @Param element query string false "elemento da magia"
// @Param page query int false "Página"
// @Param limit query int false "Limite"
// @Success 200 {object} Response
// @Router /api/spell [get]
func (h *SpellHandler) handleGetSpell(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	name := query.Get("name")
	element := query.Get("element")
	sort := query.Get("sort")
	order := query.Get("order")

	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))

	spells := h.useCase.GetWithFilters(name, element, sort, order, page, limit)

	//spells := h.useCase.GetAll()
	SendJson(w, Response{Data: spells}, http.StatusOK)

}

// GetSpellsById godoc
// @Summary Lista magias com filtros por ID
// @Tags spells
// @Produce json
// @Param id query int false "ID da magia"
// @Success 200 {object} Response
// @Router /api/spell/{id} [get]
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

// UpdateSpells godoc
// @Summary Atualiza magias
// @Tags spells
// @Produce json
// @Param id query string false "ID da magia"
// @Success 200 {object} Response
// @Router /api/spell/{id} [put]
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

// DeleteSpells godoc
// @Summary Deleta magia
// @Tags spells
// @Produce json
// @Param id query string false "ID da magia"
// @Success 200 {object} Response
// @Router /api/spell/{id} [delete]
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

// GetSpellsByElement godoc
// @Summary Lista magias por elemento
// @Tags spells
// @Produce json
// @Param element query string false "Elemento da magia"
// @Success 200 {object} Response
// @Router /api/spell/element/{element} [get]
func (h *SpellHandler) handleGetByElement(w http.ResponseWriter, r *http.Request) {
	element := chi.URLParam(r, "element")

	spells, err := h.useCase.GetByElement(element)
	if err != nil {
		SendJson(w, Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	SendJson(w, Response{Data: spells}, http.StatusOK)
}
