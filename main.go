package main

import (
	"net/http"
	"time"

	handlerApi "github.com/caioleone/go-grimoire-api/api/handler"
	repoApi "github.com/caioleone/go-grimoire-api/api/repository"
)

func main() {

}

func run() error {
	repoSpell := repoApi.NewMemoryRepositorySpell()
	handler := handlerApi.NewHandlerSpell(repoSpell)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
