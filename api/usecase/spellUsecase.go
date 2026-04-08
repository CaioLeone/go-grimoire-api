package api

import (
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
)

type SpellUsecase struct {
	repo repoApi.MemorySpellRepository
}

func NewSpellUseCase(repo repoApi.MemorySpellRepository) *SpellUsecase {
	return &SpellUsecase{repo: repo}
}
