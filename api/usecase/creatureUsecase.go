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
	//busca creature que ensina
	creature, ok := u.creatureRepo.FindById(creatureID)
	if !ok {
		return errors.New("Creature not found")
	}

	//Valida se spell existe
	_, ok = u.spellRepo.FindById(spellID)
	if !ok {
		return errors.New("Spell Not Found")
	}

	//Evita duplicacao
	for _, s := range creature.Spells {
		if s == spellID {
			return errors.New("Spell already taught by this creature")
		}
	}

	//Adiciona spell
	creature.Spells = append(creature.Spells, spellID)

	if len(creature.Spells) > 3 {
		return errors.New("Creature can only teach 3 spells")
	}

	//Salva
	_, _ = u.creatureRepo.Update(creatureID, creature)

	return nil
}
