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
