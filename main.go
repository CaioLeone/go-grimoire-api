package main

import (
	"log/slog"
	"net/http"
	"time"

	handlerApi "github.com/caioleone/go-grimoire-api/api/handler"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
	useApi "github.com/caioleone/go-grimoire-api/api/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}
	slog.Info("All Systems Offline")
}

func run() error {

	//REPOS
	repoSpell := repoApi.NewMemoryRepositorySpell()
	repoCreature := repoApi.NewMemoryRepositoryCreature()

	//USECASES
	spellUsecase := useApi.NewSpellUseCase(repoSpell)
	creatureUsecase := useApi.NewCreatureUseCase(repoCreature, repoSpell)

	//HANDLERS
	spellHandler := handlerApi.NewHandlerSpell(spellUsecase)
	creatureHandler := handlerApi.NewHandlerCreature(creatureUsecase)

	r := chi.NewMux()

	r.Mount("/", spellHandler)
	r.Mount("/", creatureHandler)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      r,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
