package api

type CreatureUsecase struct{
	repo CreatureRepository
}

func NewCreatureUseCase(repo CreatureRepository) *CreatureUsecase{
	return &CreatureUsecase{repo: repo}
}