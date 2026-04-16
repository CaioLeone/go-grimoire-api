package api

import (
	"database/sql"
	"fmt"

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
		INSERT INTO creatures(id, name, description, attack, defence, hp)
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
	query := `
	SELECT 
		c.id, c.name, c.description, c.attack, c.defence, c.hp,
		s.id, s.name, s.description, s.element, s.mana_cost 
	FROM creatures 
	LEFT JOIN creature_spells cs ON c.id = cs.creature_id
	LEFT JOIN spells s ON cs.spell_id = s.id
	WHERE c.id = ?
	`

	rows, err := r.db.Query(query, id.String())
	if err != nil {
		return domApi.CreatureModel{}, false
	}
	defer rows.Close()

	var creature domApi.CreatureModel
	var creatureID string

	spellsMap := make(map[string]domApi.SpellModel)

	for rows.Next() {
		var spell domApi.SpellModel
		var spellID sql.NullString

		err := rows.Scan(
			&creatureID,
			&creature.Name,
			&creature.Description,
			&creature.Attack,
			&creature.Defence,
			&creature.Hp,
			&spellID,
			&spell.Name,
			&spell.Description,
			&spell.Element,
			&spell.ManaCost,
		)

		if err != nil {
			return domApi.CreatureModel{}, false
		}

		creature.ID, _ = uuid.Parse(creatureID)

		//EVITA NIL SPELL
		if spellID.Valid {
			spell.ID, _ = uuid.Parse(spellID.String)

			if _, exists := spellsMap[spell.ID.String()]; !exists {
				spellsMap[spell.ID.String()] = spell
			}
		}
	}

	//TRANSFORMA MAP -> SLICE
	for _, s := range spellsMap {
		creature.Spells = append(creature.Spells, s)
	}

	return creature, true

}
func (r *SQLiteCreatureRepository) Update(id uuid.UUID, creature domApi.CreatureModel) (domApi.CreatureModel, bool) {
	query := `UPDATE creatures SET name = ?, description = ?, attack = ?, defence = ?, hp = ? WHERE id = ?`

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
func (r *SQLiteCreatureRepository) TeachSpell(creatureID, spellID uuid.UUID) error {
	// VERIFICA DUPLICACAO
	queryCheck := `
		SELECT 1 
		FROM creature_spells 
		WHERE creature_id = ? AND spell_id = ?
	`
	row := r.db.QueryRow(queryCheck, creatureID.String(), spellID.String())

	var exists int
	err := row.Scan(&exists)

	if err == nil {
		return fmt.Errorf("Spell Already Taught By This Creature")
	}

	//INSERE RELACAO
	queryInsert := `
		INSERT INTO creature_spells(creature_id, spell_id) 
		VALUES (?, ?)
	`
	_, err = r.db.Exec(queryInsert, creatureID.String(), spellID.String())
	if err != nil {
		return err
	}

	return nil
}
