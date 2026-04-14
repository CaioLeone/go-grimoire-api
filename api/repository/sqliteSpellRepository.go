package api

import (
	"database/sql"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

type SQLiteSpellRepository struct {
	db *sql.DB
}

func NewSQLiteSpellRepository(db *sql.DB) *SQLiteSpellRepository {
	return &SQLiteSpellRepository{db: db}
}

func (r *SQLiteSpellRepository) Insert(spell domApi.SpellModel) domApi.SpellModel {
	spell.ID = uuid.New()

	query := `
		INSERT INTO spells(id, name, description, element, mana_cost)
		VALUES (?, ?, ?, ?, ?)
	`

	r.db.Exec(query,
		spell.ID.String(),
		spell.Name,
		spell.Description,
		spell.Element,
		spell.ManaCost,
	)

	return spell
}

func (r *SQLiteSpellRepository) FindAll() []domApi.SpellModel {
	rows, _ := r.db.Query("SELECT id, name, description, element, mana_cost FROM spells")
	defer rows.Close()

	var spells []domApi.SpellModel

	for rows.Next() {
		var s domApi.SpellModel
		var id string

		rows.Scan(&id, &s.Name, &s.Description, &s.Element, &s.ManaCost)
		s.ID, _ = uuid.Parse(id)

		spells = append(spells, s)
	}

	return spells
}
