package api

import "github.com/google/uuid"

type SpellModel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ManaCost    int       `json:"mana_cost"`
	Element     string    `json:"element"`
	Type        string    `json:"type"`
	Power       int       `json:"power"`
}
