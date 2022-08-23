package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:dbpass@(0.0.0.0:3306)/hackernews")
	if err != nil {
		fmt.Println("MASUK1")
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("MASUK2")
		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
		fmt.Println("MASUK4")
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		
		fmt.Println("MASUK3")
		log.Fatal(err)
	}

}