package api

import "github.com/google/uuid"

type SpeelModel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ManaCost    int64     `json:"manaCost"`
	Element     string    `json:"element"`
}
