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

	//_, exists := u.repo.FindByName(spell.Name)
	// if exists {
	// 	return domApi.SpellModel{}, errors.New("Spell Already Exists")
	// }

	if len(spell.Name) <= 2 {
		return domApi.SpellModel{}, errors.New("Invalid Name")
	}

	if len(spell.Description) < 2 {
		return domApi.SpellModel{}, errors.New("Invalid Description")
	}

	if len(spell.Element) < 2 {
		return domApi.SpellModel{}, errors.New("Invalid Element")
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
	if len(s.Name) <= 2 {
		return domApi.SpellModel{}, false
	}

	if len(s.Description) < 2 {
		return domApi.SpellModel{}, false
	}

	if len(s.Element) < 2 {
		return domApi.SpellModel{}, false
	}

	if s.ManaCost <= 0 {
		return domApi.SpellModel{}, false
	}

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

func (u *SpellUsecase) GetPaginated(page, limit int) []domApi.SpellModel {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	return u.repo.FindAllPaginated(limit, offset)
}

func (u *SpellUsecase) GetWithFilters(name string, element string, sort string, order string, page int, limit int) []domApi.SpellModel {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	return u.repo.FindWithFilters(name, element, sort, order, limit, offset)
}
