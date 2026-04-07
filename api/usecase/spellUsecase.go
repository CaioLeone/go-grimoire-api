package api

type SpellUsecase struct {
	repo SpellRepository
}

func NewSpellUseCase(repo SpellRepository) *CreatureUsecase {
	return &CreatureUsecase{repo: repo}
}
