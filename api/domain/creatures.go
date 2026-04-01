package api

import "github.com/google/uuid"

type CreatureModel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name_creature"`
	Description string    `json:"description_creature"`
	Hp          int       `json:"hp_creature"`
	Attack      int       `json:"attack_creature"`
	Defence     int       `json:"defence_creature"`
}
