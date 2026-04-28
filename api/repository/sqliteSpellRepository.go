package api

import (
	"database/sql"
	"strings"

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
func (r *SQLiteSpellRepository) FindById(id uuid.UUID) (domApi.SpellModel, bool) {
	query := `SELECT id, name, description, element, mana_cost FROM spells WHERE id = ?`

	row := r.db.QueryRow(query, id.String())

	var spell domApi.SpellModel
	var idStr string

	err := row.Scan(
		&idStr,
		&spell.Name,
		&spell.Description,
		&spell.Element,
		&spell.ManaCost,
	)
	if err != nil {
		return domApi.SpellModel{}, false
	}
	spell.ID, _ = uuid.Parse(idStr)

	return spell, true
}

func (r *SQLiteSpellRepository) Update(id uuid.UUID, spell domApi.SpellModel) (domApi.SpellModel, bool) {
	query := `UPDATE spells SET name = ?, description = ?, element = ?, mana_cost = ? WHERE id = ?`
	result, err := r.db.Exec(
		query,
		spell.Name,
		spell.Description,
		spell.Element,
		spell.ManaCost,
		id.String(),
	)
	if err != nil {
		return domApi.SpellModel{}, false
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return domApi.SpellModel{}, false
	}

	spell.ID = id

	return spell, true
}

func (r *SQLiteSpellRepository) Delete(id uuid.UUID) (domApi.SpellModel, bool) {
	spell, ok := r.FindById(id)
	if !ok {
		return domApi.SpellModel{}, false
	}

	query := `DELETE FROM spells WHERE id = ?`

	_, err := r.db.Exec(query, id.String())
	if err != nil {
		return domApi.SpellModel{}, false
	}

	return spell, true
}

func (r *SQLiteSpellRepository) FindByElement(element string) []domApi.SpellModel {
	query := `
		SELECT id, name, description, element, mana_cost 
		FROM spells
		WHERE LOWER(element) = LOWER(?)
	`
	rows, err := r.db.Query(query, element)
	if err != nil {
		return []domApi.SpellModel{}
	}

	defer rows.Close()

	var spells []domApi.SpellModel

	for rows.Next() {
		var s domApi.SpellModel
		var id string

		err := rows.Scan(
			&id,
			&s.Name,
			&s.Description,
			&s.Element,
			&s.ManaCost,
		)

		if err != nil {
			continue
		}

		s.ID, _ = uuid.Parse(id)
		spells = append(spells, s)
	}

	return spells
}

func (r *SQLiteSpellRepository) FindAllPaginated(limit, offset int) []domApi.SpellModel {
	query := `
		SELECT id, name, description, element, mana_cost 
		FROM spells
		LIMIT ? OFFSET ?
	`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return []domApi.SpellModel{}
	}

	defer rows.Close()

	var spells []domApi.SpellModel

	for rows.Next() {
		var s domApi.SpellModel
		var id string

		err := rows.Scan(
			&id,
			&s.Name,
			&s.Description,
			&s.Element,
			&s.ManaCost,
		)

		if err != nil {
			continue
		}

		s.ID, _ = uuid.Parse(id)
		spells = append(spells, s)
	}

	return spells
}

func (r *SQLiteSpellRepository) FindWithFilters(name, element, sort, order string, limit, offset int) []domApi.SpellModel {
	query := `
		SELECT id, name, description, element, mana_cost 
		FROM spells
		WHERE 1=1
	`

	args := []interface{}{}

	//Filtro Name
	if name != "" {
		query += "AND LOWER(name) LIKE LOWER(?)"
		args = append(args, "%"+strings.ToLower(name)+"%")
	}

	//Filtro element
	if element != "" {
		query += "AND LOWER(element) = LOWER(?)"
		args = append(args, element)
	}

	//Ordenacao segura
	validSort := map[string]bool{
		"name":       true,
		"mana_costa": true,
	}

	if validSort[sort] {
		if order != "desc" {
			order = "asc"
		}
		query += " ORDER BY " + sort + " " + order
	}

	//Paginacao
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return []domApi.SpellModel{}
	}
	defer rows.Close()

	var spells []domApi.SpellModel

	for rows.Next() {
		var s domApi.SpellModel
		var id string
		err := rows.Scan(
			&id,
			s.Name,
			&s.Description,
			&s.Element,
			&s.ManaCost,
		)
		if err != nil {
			continue
		}

		s.ID, _ = uuid.Parse(id)
		spells = append(spells, s)
	}

	return spells
}

func (r *SQLiteSpellRepository) FindByName(name string) (domApi.SpellModel, bool) {
	query := `
		SELECT id, name, description, element, mana_cost 
		FROM spells
		WHERE Lower(name) = LOWER(?)
	`
	row := r.db.QueryRow(query, name)

	var spell domApi.SpellModel
	var id string

	err := row.Scan(
		&id,
		&spell.Name,
		&spell.Description,
		&spell.Element,
		&spell.ManaCost,
	)
	if err != nil {
		return domApi.SpellModel{}, false
	}

	spell.ID, _ = uuid.Parse(id)
	return spell, true
}
