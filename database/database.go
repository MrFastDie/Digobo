package database

import (
	"Digobo/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

var db *sqlx.DB

func Init() {
	var err error

	db, err = connect()
	if err != nil {
		log.Fatal(err)
	}
}

const NO_ROWS = "sql: no rows in result set"
const PQ_DUPLICATES = "pq: duplicate key value violates unique constraint"

func connect() (*sqlx.DB, error) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.Name, config.Config.Database.Host, config.Config.Database.Port)
	localDb, err := sqlx.Connect("postgres", connString)

	if err != nil {
		return nil, err
	}

	return localDb, nil
}

func Migrate() error {
	db, err := connect()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}
