package api

import (
	"errors"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
)

type SpellUsecase struct {
	repo repoApi.SpellRepository
}

func NewSpellUseCase(repo repoApi.SpellRepository) *SpellUsecase {
	return &SpellUsecase{repo: repo}
}

func (u *SpellUsecase) CreateSpell(spell domApi.SpellModel) (domApi.SpellModel, error) {

	if len(spell.Name) < 2 {
		return domApi.SpellModel{}, errors.New("Invalid Name")
	}

	if spell.ManaCost <= 0 {
		return domApi.SpellModel{}, errors.New("Invalid Mana Cost")
	}

	return u.repo.Insert(spell), nil
}

func (u *SpellUsecase) GetAll() []domApi.SpellModel {
	return u.repo.FindAll()
}
