package api

import (
	domainApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

type MemoryCreatureRepository struct {
	data map[uuid.UUID]domainApi.CreatureModel
}

func NewRepositoryCreature() *CreatureRepository {
	return &CreatureRepository{
		data: make(map[uuid.UUID]domainApi.CreatureModel),
	}
}

// FIND
func (c *CreatureRepository) FindAll() []domainApi.CreatureModel {
	creatures := []domainApi.CreatureModel{}

	for _, creature := range c.data {
		creatures = append(creatures, creature)
	}
	return creatures
}

func (c *CreatureRepository) FindById(id uuid.UUID) (domainApi.CreatureModel, bool) {
	creature, exists := c.data[id]
	return creature, exists
}

// INSERT
func (c *CreatureRepository) Insert(creature domainApi.CreatureModel) domainApi.CreatureModel {
	creature.ID = uuid.New()
	c.data[creature.ID] = creature
	return creature
}

// UPDATE
func (c *CreatureRepository) Update(id uuid.UUID, updated domainApi.CreatureModel) (domainApi.CreatureModel, bool) {
	updated, exists := c.data[id]
	if !exists {
		return domainApi.CreatureModel{}, false
	}

	updated.ID = id
	c.data[id] = updated
	return updated, true
}

// DELETE
func (c *CreatureRepository) Delete(id uuid.UUID) (domainApi.CreatureModel, bool) {
	deleted, exists := c.data[id]
	if !exists {
		return domainApi.CreatureModel{}, false
	}

	delete(c.data, id)
	return deleted, true
}
