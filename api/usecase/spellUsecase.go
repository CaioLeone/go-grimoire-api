package api

import (
	"errors"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
	"github.com/google/uuid"
)

type SpellUsecase struct {
	repo repoApi.SpellRepository
}

func NewSpellUseCase(repo repoApi.SpellRepository) *SpellUsecase {
	return &SpellUsecase{repo: repo}
}

func (u *SpellUsecase) CreateSpell(spell domApi.SpellModel) (domApi.SpellModel, error) {

	if len(spell.Name) <= 2 {
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

func (u *SpellUsecase) GetById(id uuid.UUID) (domApi.SpellModel, bool) {
	return u.repo.FindById(id)
}

func (u *SpellUsecase) Update(id uuid.UUID, s domApi.SpellModel) (domApi.SpellModel, bool) {
	return u.repo.Update(id, s)
}

func (u *SpellUsecase) Delete(id uuid.UUID) (domApi.SpellModel, bool) {
	return u.repo.Delete(id)
}

func (u *SpellUsecase) GetByElement(element string) ([]domApi.SpellModel, error) {
	if len(element) < 2 {
		return nil, errors.New("Invalid Element")
	}

	spells := u.repo.FindByElement(element)

	if len(spells) == 0 {
		return nil, errors.New("No Spells Found For This Element")
	}

	return spells, nil
}
