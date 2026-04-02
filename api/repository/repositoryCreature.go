package api

import (
	domainApi "github.com/caioleone/go-user-crud/api/domain/"
	"github.com/google/uuid"
)

type CreatureRepository struct {
	data map[uuid.UUID]domainApi.Creature
}
