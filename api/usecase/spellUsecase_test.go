package api

import (
	"testing"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

type fakeSpellRepo struct {
	data map[uuid.UUID]domApi.SpellModel
}

func newFakeSpellRepo() *fakeSpellRepo {
	return &fakeSpellRepo{
		data: make(map[uuid.UUID]domApi.SpellModel),
	}
}

func (f *fakeSpellRepo) Insert(s domApi.SpellModel) domApi.SpellModel {
	s.ID = uuid.New()
	return s
}

func (f *fakeSpellRepo) FindAll() []domApi.SpellModel {
	return []domApi.SpellModel{}
}

func (f *fakeSpellRepo) FindById(id uuid.UUID) (domApi.SpellModel, bool) {
	return domApi.SpellModel{ID: id}, true
}

func (f *fakeSpellRepo) FindByName(name string) (domApi.SpellModel, bool) {
	return domApi.SpellModel{Name: name}, true
}

func (f *fakeSpellRepo) Update(id uuid.UUID, s domApi.SpellModel) (domApi.SpellModel, bool) {
	s.ID = id
	return s, true
}

func (f *fakeSpellRepo) Delete(id uuid.UUID) (domApi.SpellModel, bool) {
	return domApi.SpellModel{ID: id}, true
}

func (f *fakeSpellRepo) FindAllPaginated(limit, offset int) []domApi.SpellModel {
	return []domApi.SpellModel{}
}

func (f *fakeSpellRepo) FindByElement(element string) []domApi.SpellModel {
	return []domApi.SpellModel{}
}

func (f *fakeSpellRepo) FindWithFilters(name, element, sort, order string, limit, offset int) []domApi.SpellModel {
	return []domApi.SpellModel{}
}

func TestCreateSpell_Success(t *testing.T) {
	repo := newFakeSpellRepo()

	usecase := SpellUsecase{
		spellRepo: repo,
	}

	spell := domApi.SpellModel{
		Name:        "Fireball",
		Description: "Explosão de fogo",
		Element:     "Fire",
		ManaCost:    10,
	}

	result, err := usecase.CreateSpell(spell)

	if err != nil {
		t.Fatalf("Esperava sucesso, mas deu erro: %v", err)
	}

	if result.ID == uuid.Nil {
		t.Errorf("ID Nao Foi Gerado")
	}

	if result.Name != "Fireball" {
		t.Errorf("Nome Incorreto")
	}
}

func TestCreateSpell_InvalidName(t *testing.T) {
	usecase := SpellUsecase{}
	spell := domApi.SpellModel{
		Name: "A",
	}

	_, err := usecase.CreateSpell(spell)

	if err == nil {
		t.Errorf("Esperava Erro Para Nome Invalido")
	}
}

func TestGetSpellById_Success(t *testing.T) {
	repo := &fakeSpellRepo{}

	usecase := SpellUsecase{
		spellRepo: repo,
	}

	created := repo.Insert(domApi.SpellModel{
		Name:        "Ice Spike",
		Description: "Espinho de gelo",
		Element:     "Ice",
		ManaCost:    5,
	})

	result, ok := usecase.GetById(created.ID)

	if !ok {
		t.Fatalf("Esperaeva Encontrar Spell")
	}
	if result.Name != "Ice Spike" {
		t.Errorf("Nome Incorreto")
	}
}

func TestUpdateSpell_Success(t *testing.T) {
	repo := &fakeSpellRepo{}

	usecase := SpellUsecase{
		spellRepo: repo,
	}

	created := repo.Insert(domApi.SpellModel{
		Name:        "Wind",
		Description: "Vento leve",
		Element:     "Air",
		ManaCost:    3,
	})

	updated := domApi.SpellModel{
		Name:        "Storm",
		Description: "Tempestade",
		Element:     "Air",
		ManaCost:    8,
	}

	result, ok := usecase.Update(created.ID, updated)
	if !ok {
		t.Fatalf("Update Falhou")
	}
	if result.Name != "Storm" {
		t.Errorf("Update Não Aplicado")
	}
}

func TestDeleteSpell_Success(t *testing.T) {
	repo := &fakeSpellRepo{}

	usecase := SpellUsecase{
		spellRepo: repo,
	}

	created := repo.Insert(domApi.SpellModel{
		Name:        "Lightning",
		Description: "Raio",
		Element:     "Electric",
		ManaCost:    12,
	})

	_, ok := usecase.Delete(created.ID)

	if !ok {
		t.Fatalf("Delete Falhou")
	}

	_, stillExists := repo.FindById(created.ID)

	if stillExists {
		t.Errorf("Spell Nao Foi Deletado")
	}
}

func TestGetAllSpells(t *testing.T) {
	repo := &fakeSpellRepo{}

	usecase := SpellUsecase{
		spellRepo: repo,
	}

	repo.Insert(domApi.SpellModel{Name: "Fire"})
	repo.Insert(domApi.SpellModel{Name: "Ice"})

	result := usecase.GetAll()

	if len(result) != 2 {
		t.Errorf("Esperava 2 Spells, veioi %d", len(result))
	}
}

func TestFindSpellByElement(t *testing.T) {
	repo := &fakeSpellRepo{}

	usecase := SpellUsecase{
		spellRepo: repo,
	}

	result := usecase.GetByElement("Fire")

	if result == nil {
		t.Errorf("Esperava Lista de Spells")
	}
}
