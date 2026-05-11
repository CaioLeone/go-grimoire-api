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
	if f.data == nil {
		f.data = make(map[uuid.UUID]domApi.SpellModel)
	}

	s.ID = uuid.New()
	f.data[s.ID] = s
	return s
}

func (f *fakeSpellRepo) FindAll() []domApi.SpellModel {
	var result []domApi.SpellModel

	for _, s := range f.data {
		result = append(result, s)
	}

	return result
}

func (f *fakeSpellRepo) FindById(id uuid.UUID) (domApi.SpellModel, bool) {
	s, ok := f.data[id]
	return s, ok
}

func (f *fakeSpellRepo) FindByName(name string) (domApi.SpellModel, bool) {
	return domApi.SpellModel{Name: name}, true
}

func (f *fakeSpellRepo) Update(id uuid.UUID, s domApi.SpellModel) (domApi.SpellModel, bool) {
	_, ok := f.data[id]
	if !ok {
		return domApi.SpellModel{}, false
	}

	s.ID = id
	f.data[id] = s
	return s, true
}

func (f *fakeSpellRepo) Delete(id uuid.UUID) (domApi.SpellModel, bool) {
	s, ok := f.data[id]
	if !ok {
		return domApi.SpellModel{}, false
	}
	delete(f.data, id)
	return s, true
}

func (f *fakeSpellRepo) FindAllPaginated(limit, offset int) []domApi.SpellModel {
	return []domApi.SpellModel{}
}

func (f *fakeSpellRepo) FindByElement(element string) []domApi.SpellModel {
	var result []domApi.SpellModel

	for _, s := range f.data {
		if s.Element == element {
			result = append(result, s)
		}
	}
	return result
}

func (f *fakeSpellRepo) FindWithFilters(name, element, sort, order string, limit, offset int) []domApi.SpellModel {
	return []domApi.SpellModel{}
}

func ValidSpell() domApi.SpellModel {
	return domApi.SpellModel{
		Name:        "Raios",
		Description: "Lancas de raios caem do ceu",
		ManaCost:    5,
		Element:     "Raio",
	}
}

// CREATE SPELL VALIDATION
func TestCreateSpell_Validation(t *testing.T) {
	usecase := SpellUsecase{
		repo: newFakeSpellRepo(),
	}

	tests := []struct {
		name     string
		modify   func(s *domApi.SpellModel)
		wantFail bool
	}{
		{
			name:     "Valid Spell",
			modify:   func(s *domApi.SpellModel) {},
			wantFail: false,
		},
		{
			name: "Invalid Name",
			modify: func(s *domApi.SpellModel) {
				s.Name = "A"
			},
			wantFail: true,
		},
		{
			name: "Invalid Description",
			modify: func(s *domApi.SpellModel) {
				s.Description = "L"
			},
			wantFail: true,
		},
		{
			name: "Invalid Mana Cost",
			modify: func(s *domApi.SpellModel) {
				s.ManaCost = -5
			},
			wantFail: true,
		},
		{
			name: "Invalid Element",
			modify: func(s *domApi.SpellModel) {
				s.Element = "a"
			},
			wantFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spell := ValidSpell()
			tt.modify(&spell)

			_, err := usecase.CreateSpell(spell)

			if tt.wantFail && err == nil {
				t.Errorf("Esperava Erro, Mas Veio Nil")
			}

			if !tt.wantFail && err != nil {
				t.Errorf("Nao Esperava Erro, Mas Veio: %v", err)
			}
		})
	}
}

func TestCreateSpell_Success(t *testing.T) {
	repo := newFakeSpellRepo()

	usecase := SpellUsecase{
		repo: repo,
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

// GETBYID VALIDATION
func TestSpellGetById_Validation(t *testing.T) {
	repoSpell := newFakeSpellRepo()

	usecase := SpellUsecase{
		repo: repoSpell,
	}

	//Cria Criatura valida
	created := repoSpell.Insert(ValidSpell())

	tests := []struct {
		name      string
		id        uuid.UUID
		wantFound bool
	}{
		{
			name:      "Spell Found",
			id:        created.ID,
			wantFound: true,
		},
		{
			name:      "Spell Not Found",
			id:        uuid.New(),
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, ok := usecase.GetById(tt.id)

			if ok != tt.wantFound {
				t.Errorf("Esperava %v, Mas Veio: %v", tt.wantFound, ok)
			}
		})
	}
}
func TestGetSpellById_Success(t *testing.T) {
	repo := newFakeSpellRepo()

	usecase := SpellUsecase{
		repo: repo,
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

// UPDATE SPELL VALIDATION
func TestUpdateSpell_Validation(t *testing.T) {
	tests := []struct {
		name    string
		payload domApi.SpellModel
		wantOK  bool
	}{
		{
			name: "Success Update",
			payload: domApi.SpellModel{
				Name:        "Torre de Chamas",
				Description: "Torre de chamas surge e queima tudo",
				ManaCost:    10,
				Element:     "Fogo",
			},
			wantOK: true,
		},
		{
			name: "Fail - Invalid Name",
			payload: func() domApi.SpellModel {
				s := ValidSpell()
				s.Name = "A"
				return s
			}(),
			wantOK: false,
		},
		{
			name: "Fail - Invalid Description",
			payload: func() domApi.SpellModel {
				s := ValidSpell()
				s.Description = "D"
				return s
			}(),
			wantOK: false,
		},
		{
			name: "Fail - Invalid Mana Cost",
			payload: func() domApi.SpellModel {
				s := ValidSpell()
				s.ManaCost = -1
				return s
			}(),
			wantOK: false,
		},
		{
			name: "Fail - Invalid Element",
			payload: func() domApi.SpellModel {
				s := ValidSpell()
				s.Element = "A"
				return s
			}(),
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newFakeSpellRepo()

			usecase := SpellUsecase{
				repo: repo,
			}

			created := repo.Insert(ValidSpell())

			result, ok := usecase.Update(created.ID, tt.payload)

			if ok != tt.wantOK {
				t.Errorf("Esperava %v, Mas Veio %v", tt.wantOK, ok)
			}

			if ok {
				if result.Name != tt.payload.Name {
					t.Errorf("Update Nao Aplicado Corretamente")
				}
			}
		})
	}
}

func TestUpdateSpell_Success(t *testing.T) {
	repo := newFakeSpellRepo()

	usecase := SpellUsecase{
		repo: repo,
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

// DELETE SPELL VALIDATION
func TestDeleteSpell_Validation(t *testing.T) {

	tests := []struct {
		name       string
		setup      func(repo *fakeSpellRepo) uuid.UUID
		wantDelete bool
	}{
		{
			name: "Spell Deleted",
			setup: func(repo *fakeSpellRepo) uuid.UUID {
				created := repo.Insert(ValidSpell())
				return created.ID
			},
			wantDelete: true,
		},
		{
			name: "Spell Not Found",
			setup: func(repo *fakeSpellRepo) uuid.UUID {
				return uuid.New()
			},
			wantDelete: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := newFakeSpellRepo()

			usecase := SpellUsecase{
				repo: repo,
			}

			id := tt.setup(repo)

			_, ok := usecase.Delete(id)

			if ok != tt.wantDelete {
				t.Errorf("Esperava %v, Mas Veio %v", tt.wantDelete, ok)
			}
		})
	}
}
func TestDeleteSpell_Success(t *testing.T) {
	repo := newFakeSpellRepo()

	usecase := SpellUsecase{
		repo: repo,
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

// GET ALL SPELLS VALIDATION
func TestGetAllSpells_Validation(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(repo *fakeSpellRepo)
		wantCount int
	}{
		{
			name: "Empty List",
			setup: func(repo *fakeSpellRepo) {

			},
			wantCount: 0,
		},
		{
			name: "One Spell",
			setup: func(repo *fakeSpellRepo) {
				repo.Insert(domApi.SpellModel{
					Name:        "Fireboll",
					Description: "Bola de fogo",
					Element:     "Fire",
					ManaCost:    10,
				})
			},
			wantCount: 1,
		},
		{
			name: "Multiple Spells",
			setup: func(repo *fakeSpellRepo) {
				repo.Insert(domApi.SpellModel{
					Name:        "Fireball",
					Description: "Bola de fogo",
					Element:     "Fire",
					ManaCost:    10,
				})
				repo.Insert(domApi.SpellModel{
					Name:        "Lighting Spear",
					Description: "Lancas de raio",
					Element:     "Lighting",
					ManaCost:    15,
				})
				repo.Insert(domApi.SpellModel{
					Name:        "Thunder Hammer",
					Description: "Marreta de Raio",
					Element:     "Electric",
					ManaCost:    12,
				})
			},
			wantCount: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newFakeSpellRepo()

			usecase := SpellUsecase{
				repo: repo,
			}

			tt.setup(repo)

			result := usecase.GetAll()

			if len(result) != tt.wantCount {
				t.Errorf("Esperava %d spells, mas veio %d", tt.wantCount, len(result))
			}
		})
	}
}

func TestGetAllSpells(t *testing.T) {
	repo := newFakeSpellRepo()

	usecase := SpellUsecase{
		repo: repo,
	}

	repo.Insert(domApi.SpellModel{Name: "Fire"})
	repo.Insert(domApi.SpellModel{Name: "Ice"})

	result := usecase.GetAll()

	if len(result) != 2 {
		t.Errorf("Esperava 2 Spells, veioi %d", len(result))
	}
}

// SPELL BY ELEMENT VALIDATION
func TestFindSpellByElement_Validation(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(repo *fakeSpellRepo)
		element   string
		wantCount int
		wantError bool
	}{
		{
			name: "Success - One Fire Spell",
			setup: func(repo *fakeSpellRepo) {
				fire := ValidSpell()
				fire.Name = "Fireball"
				fire.Element = "Fire"

				ice := ValidSpell()
				ice.Name = "Ice Spike"
				ice.Element = "ice"

				repo.Insert(fire)
				repo.Insert(ice)
			},
			element:   "Fire",
			wantCount: 1,
			wantError: false,
		},
		{
			name: "Success - Multiple Fire Spells",
			setup: func(repo *fakeSpellRepo) {
				fire1 := ValidSpell()
				fire1.Name = "Fireball"
				fire1.Element = "Fire"

				fire2 := ValidSpell()
				fire2.Name = "Flame Tornado"
				fire2.Element = "Fire"

				repo.Insert(fire1)
				repo.Insert(fire2)
			},
			element:   "Fire",
			wantCount: 2,
			wantError: false,
		},
		{
			name: "Fail - Element Not Found",
			setup: func(repo *fakeSpellRepo) {
				ice := ValidSpell()
				ice.Name = "Ice Spike"
				ice.Element = "Ice"

				repo.Insert(ice)
			},
			element:   "Fire",
			wantCount: 0,
			wantError: true,
		},
		{
			name: "Fail - Empty Repository",
			setup: func(repo *fakeSpellRepo) {

			},
			element:   "Fire",
			wantCount: 0,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newFakeSpellRepo()

			usecase := SpellUsecase{
				repo: repo,
			}

			tt.setup(repo)

			result, err := usecase.GetByElement(tt.element)

			//valida erro
			if tt.wantError && err == nil {
				t.Errorf("Esperava Erro, Mais Veio Nil")
			}

			if !tt.wantError && err != nil {
				t.Errorf("Nao Esperava Erro, Mas Veio: %v", err)
			}

			//Valida Quantidade
			if len(result) != tt.wantCount {
				t.Errorf("Esperava %d Spells, Mas Veio %d", tt.wantCount, len(result))
			}

			//Valida Se Todos Os Elementos Retornados Sao Corretos
			for _, spell := range result {
				if spell.Element != tt.element {
					t.Errorf("Esperava Element %s, Mas Veio %s", tt.element, spell.Element)
				}
			}
		})
	}
}

func TestFindSpellByElement(t *testing.T) {
	repo := newFakeSpellRepo()

	usecase := SpellUsecase{
		repo: repo,
	}

	repo.Insert(domApi.SpellModel{
		Name:    "Fireball",
		Element: "Fire",
	})

	repo.Insert(domApi.SpellModel{
		Name:    "Ice Spike",
		Element: "Ice",
	})

	result, err := usecase.GetByElement("Fire")

	if err != nil {
		t.Errorf("Nao Esperava Erro, Mas Veio: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Esperava 1 Spell, veio %d", len(result))
	}

	if result[0].Element != "Fire" {
		t.Errorf("Filtro Nao Funciona")
	}
}
