package api

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./grimoire.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB(db *sql.DB) error {
	querySpell := `
		CREATE TABLE IF NOT EXISTS spells (
			id TEXT PRIMARY KEY,
			name TEXT,
			description TEXT,
			element TEXT,
			mana_cost INTEGER,
			type TEXT,
			power INTEGER
		);`

	queryCreature := `
		CREATE TABLE IF NOT EXISTS creatures (
			id TEXT PRIMARY KEY,
			name TEXT,
			description TEXT,
			attack INTEGER,
			defence INTEGER,
			hp INTEGER
		);`

	queryRelation := `
		CREATE TABLE IF NOT EXISTS creature_spells (
			creature_id TEXT,
			spell_id TEXT
		);`

	_, err := db.Exec(querySpell)
	if err != nil {
		return err
	}

	_, err = db.Exec(queryCreature)
	if err != nil {
		return err
	}

	_, err = db.Exec(queryRelation)
	if err != nil {
		return err
	}

	return nil
}
