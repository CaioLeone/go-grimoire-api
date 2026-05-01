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
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		creature.ID.String(),
		creature.Name,
		creature.Description,
		creature.Attack,
		creature.Defence,
		creature.Hp,
	)
	if err != nil {
		fmt.Println("INSERT ERROR: ", err)
	}

	return creature
}

func (r *SQLiteCreatureRepository) FindAll() []domApi.CreatureModel {
	rows, err := r.db.Query("SELECT id, name, description, attack, defence, hp FROM creatures")
	if err != nil {
		return []domApi.CreatureModel{}
	}
	defer rows.Close()

	var creatures []domApi.CreatureModel

	for rows.Next() {
		var c domApi.CreatureModel
		var id string

		rows.Scan(&id, &c.Name, &c.Description, &c.Attack, &c.Defence, &c.Hp)
		c.ID, _ = uuid.Parse(id)

		creatures = append(creatures, c)
		fmt.Println("CREATURE: ", c.Name)
	}
	fmt.Println(">>> BUSCANDO TODAS CRIATURAS <<< ")
	return creatures
}

func (r *SQLiteCreatureRepository) FindById(id uuid.UUID) (domApi.CreatureModel, bool) {
	queryCreature := `
		SELECT id,name, description, attack, defence, hp
		FROM creatures
		WHERE id = ?
	`
	var creature domApi.CreatureModel
	var creatureID string

	err := r.db.QueryRow(queryCreature, id.String()).Scan(
		&creatureID,
		&creature.Name,
		&creature.Description,
		&creature.Attack,
		&creature.Defence,
		&creature.Hp,
	)

	if err != nil {
		return domApi.CreatureModel{}, false
	}
	fmt.Println(">>> CRIATURA ENCONTRADA <<<")
	creature.ID, _ = uuid.Parse(creatureID)

	querySpells := `
		SELECT s.id, s.name, s.description, s.element, s.mana_cost
		FROM spells s
		INNER JOIN creature_spells cs ON s.id = cs.spell_id
		WHERE cs.creature_id = ?
	`
	//RETORNA CRIATURA SEM SPELLS
	rows, err := r.db.Query(querySpells, id.String())
	if err != nil {
		return creature, false
	}
	defer rows.Close()

	//spellsMap := make(map[string]domApi.SpellModel)

	for rows.Next() {

		var spell domApi.SpellModel
		var spellID string

		err := rows.Scan(
			&spellID,
			&spell.Name,
			&spell.Description,
			&spell.Element,
			&spell.ManaCost,
		)

		if err != nil {
			continue
		}

		spell.ID, _ = uuid.Parse(spellID)
		creature.Spells = append(creature.Spells, spell)
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

	result, err := r.db.Exec(query, id.String())
	if err != nil {
		return domApi.CreatureModel{}, false
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
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

func (r *SQLiteCreatureRepository) FindAllPaginated(limit, offset int) []domApi.CreatureModel {
	query := `
		SELECT id, name, description, attack, defence, hp
		FROM creatures
		LIMIT ? OFFSET ?
	`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return []domApi.CreatureModel{}
	}

	var creatures []domApi.CreatureModel

	for rows.Next() {
		var c domApi.CreatureModel
		var id string

		err := rows.Scan(
			&id,
			&c.Name,
			&c.Description,
			&c.Attack,
			&c.Defence,
			&c.Hp,
		)

		if err != nil {
			continue
		}

		c.ID, _ = uuid.Parse(id)
		creatures = append(creatures, c)
	}
	return creatures
}

func (r *SQLiteCreatureRepository) FindWithFilters(name string, attack int, defence int, sort string, order string, limit int, offset int) []domApi.CreatureModel {
	query := `
		SELECT id, name, description, attack, defence, hp
		FROM creatures
		WHERE 1=1
	`
	args := []interface{}{}

	//BUSCA POR NOME
	if name != "" {
		query += " AND LOWER(name) LIKE LOWER(?)"
		args = append(args, "%"+name+"%")
	}

	//BUSCA POR ATAQUE
	if attack != 0 {
		query += " AND attack = ?"
		args = append(args, attack)
	}

	//BUSCA POR ATAQUE
	if defence != 0 {
		query += " AND defence = ?"
		args = append(args, defence)
	}

	validSort := map[string]bool{
		"name":    true,
		"attack":  true,
		"defence": true,
		"hp":      true,
	}

	//ORDENCAO SEGURA
	if validSort[sort] {
		if order != "desc" {
			order = "asc"
		}
		query += " ORDER BY " + sort + " " + order
	}

	//PAGINACAO
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		fmt.Println("SQL ERROR: ", err)
		return []domApi.CreatureModel{}
	}

	defer rows.Close()

	var creatures []domApi.CreatureModel

	for rows.Next() {
		var c domApi.CreatureModel
		var id string

		err := rows.Scan(
			&id,
			&c.Name,
			&c.Description,
			&c.Attack,
			&c.Defence,
			&c.Hp,
		)

		if err != nil {
			continue
		}

		c.ID, _ = uuid.Parse(id)
		creatures = append(creatures, c)
	}

	return creatures
}
