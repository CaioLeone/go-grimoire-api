package main

import (
	"net/http"
	"time"

	newApi "github.com/caioleone/go-user-crud/api"
)

func main() {

}

func run() error {
	db := make(map[string]string)
	handler := newApi.NewHandler(db)

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
