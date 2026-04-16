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

func (u *CreatureUsecase) CreateCreature(creature domApi.CreatureModel) (domApi.CreatureModel, error) {

	//VALIDACAO
	if len(creature.Name) < 2 {
		return domApi.CreatureModel{}, errors.New("Invalid Name")
	}

	if len(creature.Description) < 2 {
		return domApi.CreatureModel{}, errors.New("Invalid Description")
	}

	if creature.Attack <= 0 {
		return domApi.CreatureModel{}, errors.New("Invalid Attack")
	}

	if creature.Defence <= 0 {
		return domApi.CreatureModel{}, errors.New("Invalid Defence")
	}

	if creature.Hp <= 0 {
		return domApi.CreatureModel{}, errors.New("Invalid Hp")
	}

	return u.creatureRepo.Insert(creature), nil
}

func (u *CreatureUsecase) GetAll() []domApi.CreatureModel {
	return u.creatureRepo.FindAll()
}

func (u *CreatureUsecase) GetById(id uuid.UUID) (domApi.CreatureModel, bool) {
	return u.creatureRepo.FindById(id)
}

func (u *CreatureUsecase) Update(id uuid.UUID, s domApi.CreatureModel) (domApi.CreatureModel, bool) {
	return u.creatureRepo.Update(id, s)
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
