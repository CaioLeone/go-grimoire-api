package api

import (
	domainApi "github.com/caioleone/go-user-crud/api/domain/"
	"github.com/google/uuid"
)

type SpellRepository struct {
	data map[uuid.UUID]domainApi.Spell
}

func NewRepositorySpell() *SpellRepository {
	return &SpellRepository{
		data: make(map[uuid.UUID]domainApi.Spell),
	}
}

// FIND
func (s *SpellRepository) FindAll() []domainApi.Spell {
	spells := []domainApi.Spell{}

	for _, spell := range s.data {
		spells = append(spells, spell)
	}

	return spells
}

func (s *SpellRepository) FindById(id uuid.UUID) (domainApi.Spell, bool) {
	spell, exists := s.data[id]
	return spell, exists
}

//INSERT

func (s *SpellRepository) Insert(spell domainApi.Spell) domainApi.Spell {
	spell.ID = uuid.New()
	s.data[spell.ID] = spell
	return spell
}

// UPDATE
func (s *SpellRepository) Update(id uuid.UUID, updated domainApi.Spell) (domainApi.Spell, bool) {
	_, exists := s.data[id]
	if !exists {
		return domainApi.Spell{}, false
	}

	updated.ID = id
	s.data[id] = updated
	return updated, true
}

// DELETE
func (s *SpellRepository) Delete(id uuid.UUID) (domainApi.Spell, bool) {
	spell, exists := s.data[id]
	if !exists {
		return domainApi.Spell{}, false
	}
	delete(s.data, id)
	return spell, true
}
