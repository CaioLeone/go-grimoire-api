package api

import (
	domainApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

type CreatureRepository interface {
	FindAll() []domainApi.CreatureModel
	FindById(id uuid.UUID) (domainApi.CreatureModel, bool)
	Insert(creature domainApi.CreatureModel) domainApi.CreatureModel
	Update(id uuid.UUID, creature domainApi.CreatureModel) (domainApi.CreatureModel, bool)
	Delete(id uuid.UUID) (domainApi.CreatureModel, bool)

	TeachSpell(creatureID, spellID uuid.UUID) error

	//Pagination
	FindAllPaginated(limit, offset int) []domainApi.CreatureModel
}

type SpellRepository interface {
	FindAll() []domainApi.SpellModel
	FindById(id uuid.UUID) (domainApi.SpellModel, bool)
	Insert(spell domainApi.SpellModel) domainApi.SpellModel
	Update(id uuid.UUID, spell domainApi.SpellModel) (domainApi.SpellModel, bool)
	Delete(id uuid.UUID) (domainApi.SpellModel, bool)

	FindByElement(element string) []domainApi.SpellModel

	//Pagination
	FindAllPaginated(limit, offset int) []domainApi.SpellModel
}
