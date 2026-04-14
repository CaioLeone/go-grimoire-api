package api

import (
	"database/sql"

	domApi "github.com/caioleone/go-grimoire-api/api/domain"
	"github.com/google/uuid"
)

type SQLiteCreatureRepository struct {
	db *sql.DB
}

func NewSQLiteCreatureRepository(db *sql.DB) *SQLiteCreatureRepository {
	return &SQLiteCreatureRepository{db: db}
}

func (r *SQLiteCreatureRepository) Insert(creature domApi.CreatureModel) domApi.CreatureModel {
	creature.ID = uuid.New()

	query := `
		INSERT INTO spells(id, name, description, attack, defence, hp)
		VALUES (?, ?, ?, ?, ?)
	`

	r.db.Exec(query,
		creature.ID.String(),
		creature.Name,
		creature.Description,
		creature.Attack,
		creature.Defence,
		creature.Hp,
	)

	return creature
}

func (r *SQLiteCreatureRepository) FindAll() []domApi.CreatureModel {
	rows, _ := r.db.Query("SELECT id, name, description, attack, defence, hp FROM creatures")
	defer rows.Close()

	var creatures []domApi.CreatureModel

	for rows.Next() {
		var c domApi.CreatureModel
		var id string

		rows.Scan(&id, &c.Name, &c.Description, &c.Attack, &c.Defence, &c.Hp)
		c.ID, _ = uuid.Parse(id)

		creatures = append(creatures, c)
	}

	return creatures
}

func (r *SQLiteCreatureRepository) FindById(id uuid.UUID) (domApi.CreatureModel, bool) {
	query := `SELECT id, name, description, attack, defence, hp FROM creatures WHERE id = ?`

	row := r.db.QueryRow(query, id.String())

	var creature domApi.CreatureModel
	var idStr string

	err := row.Scan(
		&idStr,
		&creature.Name,
		&creature.Description,
		&creature.Attack,
		&creature.Defence,
		&creature.Hp,
	)

	if err != nil {
		return domApi.CreatureModel{}, false
	}

	creature.ID, _ = uuid.Parse(idStr)
	return creature, true

}
func (r *SQLiteCreatureRepository) Update(id uuid.UUID, creature domApi.CreatureModel) (domApi.CreatureModel, bool) {
	query := `UPDATE creatyres SET name = ?, description = ?, attack = ?, defence = ?, hp = ? WHERE id = ?`

	result, err := r.db.Exec(
		query,
		creature.Name,
		creature.Description,
		creature.Attack,
		creature.Defence,
		creature.Hp,
		id.String(),
	)

	if err != nil {
		return domApi.CreatureModel{}, false
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return domApi.CreatureModel{}, false
	}

	creature.ID = id
	return creature, true
}

func (r *SQLiteCreatureRepository) Delete(id uuid.UUID) (domApi.CreatureModel, bool) {
	creature, ok := r.FindById(id)
	if !ok {
		return domApi.CreatureModel{}, false
	}

	query := `DELETE FROM creatures WHERE id = ?`

	_, err := r.db.Exec(query, id.String())
	if err != nil {
		return domApi.CreatureModel{}, false
	}

	return creature, true
}
