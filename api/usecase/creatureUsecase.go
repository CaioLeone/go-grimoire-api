package api

import (
	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
)

type CreatureUsecase struct {
	repo repoApi.CreatureRepository
}

func NewCreatureUseCase(repo repoApi.CreatureRepository) *CreatureUsecase {
	return &CreatureUsecase{repo: repo}
}

func (u *CreatureUsecase) CreateCreature(c domApi.CreatureModel) domApi.CreatureModel {
	return u.repo.Insert(c)
}
