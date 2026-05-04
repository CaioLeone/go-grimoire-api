package api

import (
	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

type fakeSpellRepo struct{}

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
