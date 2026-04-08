package api

import (
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
)

type CreatureUsecase struct{
	repo repoApi.MemoryCreatureRepository
}

func NewCreatureUseCase(repo repoApi.MemoryCreatureRepository) *CreatureUsecase{
	return &CreatureUsecase{repo: repo}
}