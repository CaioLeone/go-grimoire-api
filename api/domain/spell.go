package api

import "github.com/google/uuid"

type SpellModel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name_spell"`
	Description string    `json:"description_spell"`
	ManaCost    int64     `json:"manaCost"`
	Element     string    `json:"element"`
}
