package api

import (
	"fmt"
	"testing"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

// Fake Repository (mock simples)
type fakeCreatureRepo struct {
	data map[uuid.UUID]domApi.CreatureModel
}

func newFakeCreatureRepo() *fakeCreatureRepo {
	return &fakeCreatureRepo{
		data: make(map[uuid.UUID]domApi.CreatureModel),
	}
}

//type fakeSpellRepo struct{}

func (f *fakeCreatureRepo) Insert(c domApi.CreatureModel) domApi.CreatureModel {
	c.ID = uuid.New()
	f.data[c.ID] = c
	return c
}

func (f *fakeCreatureRepo) FindAll() []domApi.CreatureModel {
	var result []domApi.CreatureModel

	for _, c := range f.data {
		result = append(result, c)
	}

	return result
}

func (f *fakeCreatureRepo) FindById(id uuid.UUID) (domApi.CreatureModel, bool) {
	c, ok := f.data[id]
	return c, ok
}

func (f *fakeCreatureRepo) Update(id uuid.UUID, c domApi.CreatureModel) (domApi.CreatureModel, bool) {
	_, ok := f.data[id]
	if !ok {
		return domApi.CreatureModel{}, false
	}
	c.ID = id
	f.data[id] = c
	return c, true
}

func (f *fakeCreatureRepo) Delete(id uuid.UUID) (domApi.CreatureModel, bool) {
	c, ok := f.data[id]
	if !ok {
		return domApi.CreatureModel{}, false
	}
	delete(f.data, id)
	return c, true
}

func (f *fakeCreatureRepo) TeachSpell(creatureID, spellID uuid.UUID) error {
	c, ok := f.data[creatureID]
	if !ok {
		return fmt.Errorf("Creature Not Found")
	}
	c.Spells = append(c.Spells, domApi.SpellModel{ID: spellID})
	f.data[creatureID] = c
	return nil
}

func (f *fakeCreatureRepo) FindAllPaginated(limit, offset int) []domApi.CreatureModel {
	return f.FindAll()
}

func (f *fakeCreatureRepo) FindWithFilters(name string, attack int, defence int, sort string, order string, limit int, offset int) []domApi.CreatureModel {
	return f.FindAll()
}

func ValidCreature() domApi.CreatureModel {
	return domApi.CreatureModel{
		Name:        "Orc",
		Description: "Bruto forte",
		Attack:      5,
		Defence:     3,
		Hp:          20,
	}
}

// Implementa so o que precisa
func TestCreateCreature_Success(t *testing.T) {
	repo := newFakeCreatureRepo()

	usecase := CreatureUsecase{
		creatureRepo: repo,
		spellRepo:    &fakeSpellRepo{},
	}

	creature := domApi.CreatureModel{
		Name:        "Goblin",
		Description: "Pequeno e indisciplinado",
		Attack:      3,
		Defence:     2,
		Hp:          10,
	}

	result, err := usecase.CreateCreature(creature)

	if err != nil {
		t.Fatalf("Esperava sucesso, mas deu erro: %v", err)
	}

	if result.ID == uuid.Nil {
		t.Errorf("ID nao foi gerado")
	}

	if result.Name != "Goblin" {
		t.Errorf("Nome Incorreto")
	}
}

func TestCreateCreature_Validation(t *testing.T) {
	usecase := CreatureUsecase{
		creatureRepo: newFakeCreatureRepo(),
	}

	tests := []struct {
		name     string
		modify   func(c *domApi.CreatureModel)
		wantFail bool
	}{
		{
			name:     "Valid Creature",
			modify:   func(c *domApi.CreatureModel) {},
			wantFail: false,
		},
		{
			name: "Invalid Name",
			modify: func(c *domApi.CreatureModel) {
				c.Name = "A"
			},
			wantFail: true,
		},
		{
			name: "Invalid Description",
			modify: func(c *domApi.CreatureModel) {
				c.Description = "L"
			},
			wantFail: true,
		},
		{
			name: "Invalid Attack",
			modify: func(c *domApi.CreatureModel) {
				c.Attack = -5
			},
			wantFail: true,
		},
		{
			name: "Invalid Defence",
			modify: func(c *domApi.CreatureModel) {
				c.Defence = -4
			},
			wantFail: true,
		},
		{
			name: "Invalid HP",
			modify: func(c *domApi.CreatureModel) {
				c.Hp = -50
			},
			wantFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creature := ValidCreature()
			tt.modify(&creature)

			_, err := usecase.CreateCreature(creature)

			if tt.wantFail && err == nil {
				t.Errorf("Esperava Erro, Mas Veio Nil")
			}

			if !tt.wantFail && err != nil {
				t.Errorf("Nao Esperava Erro, Mas Veio: %v", err)
			}
		})
	}
}

func TestCreateCreature_InvalidName(t *testing.T) {
	usecase := CreatureUsecase{}

	creature := domApi.CreatureModel{
		Name: "A", //invalido
	}

	_, err := usecase.CreateCreature(creature)
	if err == nil {
		t.Errorf("Esperava erro para nome invalido")
	}
}

func TestGetById_Success(t *testing.T) {
	repo := newFakeCreatureRepo()

	usecase := CreatureUsecase{
		creatureRepo: repo,
	}

	created := repo.Insert(domApi.CreatureModel{
		Name:        "Orc",
		Description: "Bruto bem treinado",
		Attack:      5,
		Defence:     3,
		Hp:          20,
	})

	result, ok := usecase.GetById(created.ID)

	if !ok {
		t.Fatalf("Esperava Encontrar Criatura")
	}
	if result.Name != "Orc" {
		t.Errorf("Nome Incorreto")
	}
}

func TestUpdateCreature_Success(t *testing.T) {
	repo := newFakeCreatureRepo()

	usecase := CreatureUsecase{
		creatureRepo: repo,
	}

	created := repo.Insert(domApi.CreatureModel{

		Name:        "Slime",
		Description: "Gelatinoso",
		Attack:      1,
		Defence:     1,
		Hp:          5,
	})

	updated := domApi.CreatureModel{
		Name:        "Mega Slime",
		Description: "Maior",
		Attack:      2,
		Defence:     2,
		Hp:          10,
	}

	result, ok := usecase.Update(created.ID, updated)

	if !ok {
		t.Fatalf("Update Falhou")
	}
	if result.Name != "Mega Slime" {
		t.Errorf("Update Nao Aplicado")
	}
}

func TestDeleteCreature_Success(t *testing.T) {
	repo := newFakeCreatureRepo()

	usecase := CreatureUsecase{
		creatureRepo: repo,
	}

	created := repo.Insert(domApi.CreatureModel{
		Name:        "Dragon",
		Description: "Feroz",
		Attack:      10,
		Defence:     8,
		Hp:          100,
	})

	_, ok := usecase.Delete(created.ID)

	if !ok {
		t.Fatalf("Delete Falhou")
	}

	_, stillExists := repo.FindById(created.ID)

	if stillExists {
		t.Errorf("Criatura Nao Foi Deletada")
	}
}

func TestTeachSpell_Success(t *testing.T) {
	repo := newFakeCreatureRepo()
	spellRepo := newFakeSpellRepo()

	usecase := CreatureUsecase{
		creatureRepo: repo,
		spellRepo:    spellRepo,
	}

	created := repo.Insert(domApi.CreatureModel{
		Name:        "Mage",
		Description: "Sabio",
		Attack:      2,
		Defence:     2,
		Hp:          10,
	})

	spell := spellRepo.Insert(domApi.SpellModel{
		Name:    "Fireball",
		Element: "Fire",
	})

	err := usecase.TeachSpell(created.ID, spell.ID)

	if err != nil {
		t.Fatalf("Erro ao ensinar magia: %v", err)
	}

	c, _ := repo.FindById(created.ID)

	if len(c.Spells) == 0 {
		t.Errorf("Magia Nao Foi Adicionado")
	}
}

func TestTeachSpell_SpellNotFound(t *testing.T) {
	repo := newFakeCreatureRepo()
	spellRepo := newFakeSpellRepo()

	usecase := CreatureUsecase{
		creatureRepo: repo,
		spellRepo:    spellRepo,
	}

	created := repo.Insert(domApi.CreatureModel{
		Name: "Feiticeira",
	})

	err := usecase.TeachSpell(created.ID, uuid.New())

	if err == nil {
		t.Errorf("Esperava erro de spell nao encontrado")
	}
}
