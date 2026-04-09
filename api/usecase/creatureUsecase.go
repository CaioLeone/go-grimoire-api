package api

import (
	"errors"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
	"github.com/google/uuid"
)

type CreatureUsecase struct {
	repo repoApi.CreatureRepository
}

func NewCreatureUseCase(repo repoApi.CreatureRepository) *CreatureUsecase {
	return &CreatureUsecase{repo: repo}
}

func (u *CreatureUsecase) CreateCreature(creature domApi.CreatureModel) (domApi.CreatureModel, error) {

	//VALIDACAO
	if len(creature.Name) < 2 {
		return domApi.CreatureModel{}, errors.New("Invalid Name")
	}

	if len(creature.Description) < 2 {
		return domApi.CreatureModel{}, errors.New("Invalid Name")
	}

	if creature.Attack <= 0 {
		return domApi.CreatureModel{}, errors.New("Invalid Name")
	}

	if creature.Defence <= 0 {
		return domApi.CreatureModel{}, errors.New("Invalid Name")
	}

	if creature.Hp <= 0 {
		return domApi.CreatureModel{}, errors.New("Invalid Name")
	}

	return u.repo.Insert(creature), nil
}

func (u *CreatureUsecase) GetAll() []domApi.CreatureModel {
	return u.repo.FindAll()
}

func (u *CreatureUsecase) GetById(id uuid.UUID) (domApi.CreatureModel, bool) {
	return u.repo.FindById(id)
}

func (u *CreatureUsecase) Update(id uuid.UUID, s domApi.CreatureModel) (domApi.CreatureModel, bool) {
	return u.repo.Update(id, s)
}

func (u *CreatureUsecase) Delete(id uuid.UUID) (domApi.CreatureModel, bool) {
	return u.repo.Delete(id)
}
