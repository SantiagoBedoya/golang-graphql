package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/hackernews")
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	DB = db
}

func CloseDB() error {
	return DB.Close()
}

func Migrate() {
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(DB, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
