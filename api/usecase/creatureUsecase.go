package api

import (
	"errors"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
	"github.com/google/uuid"
)

type CreatureUsecase struct {
	creatureRepo repoApi.CreatureRepository
	spellRepo    repoApi.SpellRepository
}

func NewCreatureUseCase(creatureRepo repoApi.CreatureRepository, spellRepo repoApi.SpellRepository) *CreatureUsecase {
	return &CreatureUsecase{
		creatureRepo: creatureRepo,
		spellRepo:    spellRepo,
	}
}

func validateCreature(creature domApi.CreatureModel) error {
	if len(creature.Name) < 2 {
		return errors.New("Invalid Name")
	}

	if len(creature.Description) < 2 {
		return errors.New("Invalid Description")
	}

	if creature.Attack <= 0 {
		return errors.New("Invalid Attack")
	}

	if creature.Defence <= 0 {
		return errors.New("Invalid Defence")
	}

	if creature.Hp <= 0 {
		return errors.New("Invalid Hp")
	}

	return nil
}

func (u *CreatureUsecase) CreateCreature(creature domApi.CreatureModel) (domApi.CreatureModel, error) {

	if err := validateCreature(creature); err != nil {
		return domApi.CreatureModel{}, err
	}

	return u.creatureRepo.Insert(creature), nil
}

func (u *CreatureUsecase) GetAll() []domApi.CreatureModel {
	return u.creatureRepo.FindAll()
}

func (u *CreatureUsecase) GetById(id uuid.UUID) (domApi.CreatureModel, bool) {
	return u.creatureRepo.FindById(id)
}

func (u *CreatureUsecase) Update(id uuid.UUID, c domApi.CreatureModel) (domApi.CreatureModel, bool) {
	if err := validateCreature(c); err != nil {
		return domApi.CreatureModel{}, false
	}

	return u.creatureRepo.Update(id, c)
}

func (u *CreatureUsecase) Delete(id uuid.UUID) (domApi.CreatureModel, bool) {
	return u.creatureRepo.Delete(id)
}

// Teach Spell
func (u *CreatureUsecase) TeachSpell(creatureID, spellID uuid.UUID) error {
	//VALIDA CREATURE
	_, ok := u.creatureRepo.FindById(creatureID)
	if !ok {
		return errors.New("Creature Not Found")
	}

	//VALIDA SPELL
	_, ok = u.spellRepo.FindById(spellID)
	if !ok {
		return errors.New("Spell Not Found")
	}

	return u.creatureRepo.TeachSpell(creatureID, spellID)
}

func (u *CreatureUsecase) GetPaginated(page, limit int) []domApi.CreatureModel {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	return u.creatureRepo.FindAllPaginated(limit, offset)
}

func (u *CreatureUsecase) GetWithFilters(name string, attack int, defence int, sort string, order string, page int, limit int) []domApi.CreatureModel {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	return u.creatureRepo.FindWithFilters(name, attack, defence, sort, order, limit, offset)
}
