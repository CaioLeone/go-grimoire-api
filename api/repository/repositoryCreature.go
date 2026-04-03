package api

import (
	domainApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

type CreatureRepository struct {
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
