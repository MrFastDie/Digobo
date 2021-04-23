package database

import (
	"Digobo/config"
	"database/sql"
	"fmt"
	"github.com/Femaref/dbx"
	"github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

var db dbx.DBAccess

func Init() {
	db, _ = connect(config.Config.Database.Name, config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.Host, config.Config.Database.Port)
}

func TestDatabase(name, user, password, host string, port int) bool {
	_, err := connect(name, user, password, host, port)

	if err != nil {
		return false
	}

	return true
}

func connect(name, user, password, host string, port int) (dbx.DBAccess, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.Host, config.Config.Database.Port, config.Config.Database.Password)
	dbx.Configure("postgres", connString)
	dbx.QuoteIdentifier = pq.QuoteIdentifier

	db := dbx.MustConnect()
	return db, db.Ping()
}

func Migrate() error {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.Host, config.Config.Database.Port, config.Config.Database.Password))
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
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