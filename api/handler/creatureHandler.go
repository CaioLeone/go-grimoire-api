package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
)

// CreateCreature godoc
// @Summary Cria uma criatura
// @Description Cria uma nova criatura
// @Tags creatures
// @Accept json
// @Produce json
// @Param creature body domApi.CreatureModel true "Dados da criatura"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /api/creature [post]
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

// GetAllCreatures godoc
// @Summary Lista todas as criaturas
// @Description Lista todas as crituras com filtros opcionais e paginação
// @Tags creatures
// @Produce json
// @Param name query string false "nome da criatura"
// @Param attack query int false "Ataque"
// @Param defence query int false "Defesa"
// @Param sort query string false "Campo de ordenação (name, attack)"
// @Param order query string false "asc ou desc"
// @Param page query int false "Pagina"
// @Param limit query int false "limite"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /api/creature [get]
func (h *CreatureHandler) handleGetAllCreature(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	name := query.Get("name")
	attack, _ := strconv.Atoi(query.Get("attack"))
	defence, _ := strconv.Atoi(query.Get("defence"))
	sort := query.Get("sort")
	order := query.Get("order")

	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))

	creatures := h.useCase.GetWithFilters(name, attack, defence, sort, order, page, limit)

	//creatures := h.useCase.GetAll()
	SendJson(w, Response{Data: creatures}, http.StatusOK)

}

// GetCreatureByID godoc
// @Summary Busca criatura por ID
// @Description Retorna uma criaturas especifica
// @Tags creatures
// @Produce json
// @Param id path string false "ID da criatura"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /api/creature/{id} [get]
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
		return
	}
	SendJson(w, Response{Data: creature}, http.StatusOK)

}

// UpdateCreature godoc
// @Summary Atualiza criatura
// @Description Atualiza dados da criatura
// @Tags creatures
// @Accept json
// @Produce json
// @Param id path string false "ID da criatura"
// @Param creature body domApi.CreatureModel true "Dados Atualizados"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /api/creature/{id} [put]
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

// DeleteCreature godoc
// @Summary Deleta criatura
// @Description Deleta uma criatura do grimorio
// @Tags creatures
// @Produce json
// @Param id path string false "ID da criatura"
// @Success 201 {object} Response
// @Failure 404 {object} Response
// @Router /api/creature/{id} [delete]
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

// TeachSpell godoc
// @Summary Criatura ensina uma nova magia
// @Description Associa uma magina a uma Criatura
// @Tags creatures
// @Produce json
// @Param id path string false "ID da criatura"
// @Param spellId path string false "ID da magia"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /api/creature/{id}/teach/{spellId} [post]
func (h *CreatureHandler) handleTeachSpell(w http.ResponseWriter, r *http.Request) {
	creatureIdStr := chi.URLParam(r, "id")
	spellIdStr := chi.URLParam(r, "spellId")

	creatureId, err := uuid.Parse(creatureIdStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Creature ID"}, http.StatusBadRequest)
		return
	}

	spellId, err := uuid.Parse(spellIdStr)
	if err != nil {
		SendJson(w, Response{Error: "Invalid Spell ID"}, http.StatusBadRequest)
		return
	}

	err = h.useCase.TeachSpell(creatureId, spellId)
	if err != nil {
		SendJson(w, Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	SendJson(w, Response{Data: "Spell taught Succesfully"}, http.StatusOK)
}
