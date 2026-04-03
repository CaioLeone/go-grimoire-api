package api

import (
	domainApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

type SpellRepository struct {
	data map[uuid.UUID]domainApi.SpellModel
}

func NewRepositorySpell() *SpellRepository {
	return &SpellRepository{
		data: make(map[uuid.UUID]domainApi.SpellModel),
	}
}

// FIND
func (s *SpellRepository) FindAll() []domainApi.SpellModel {
	spells := []domainApi.SpellModel{}

	for _, spell := range s.data {
		spells = append(spells, spell)
	}

	return spells
}

func (s *SpellRepository) FindById(id uuid.UUID) (domainApi.SpellModel, bool) {
	spell, exists := s.data[id]
	return spell, exists
}

//INSERT

func (s *SpellRepository) Insert(spell domainApi.SpellModel) domainApi.SpellModel {
	spell.ID = uuid.New()
	s.data[spell.ID] = spell
	return spell
}

// UPDATE
func (s *SpellRepository) Update(id uuid.UUID, updated domainApi.SpellModel) (domainApi.SpellModel, bool) {
	_, exists := s.data[id]
	if !exists {
		return domainApi.SpellModel{}, false
	}

	updated.ID = id
	s.data[id] = updated
	return updated, true
}

// DELETE
func (s *SpellRepository) Delete(id uuid.UUID) (domainApi.SpellModel, bool) {
	spell, exists := s.data[id]
	if !exists {
		return domainApi.SpellModel{}, false
	}
	delete(s.data, id)
	return spell, true
}
