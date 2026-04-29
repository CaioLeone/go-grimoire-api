package api

import "github.com/google/uuid"

type CreatureModel struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Hp          int          `json:"hp"`
	Attack      int          `json:"attack"`
	Defence     int          `json:"defence"`
	Spells      []SpellModel `json:"teach_spell,omitempty"`
}
