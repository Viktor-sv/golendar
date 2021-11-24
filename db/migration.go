package db

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	database string = "golendar"
	IP       string = "127.0.0.1"
	PORT     string = "3306"
)

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/golendar?multiStatements=true")
	if err != nil {
		fmt.Println("sql.Open() fail", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect to database!")
		return nil, err
	}

	fmt.Println("Connected!")
	return db, nil
}

func Run() error {
	db, err := connectToDB()
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
	}

	//err = m.Steps(2)
	err = m.Up()
	if err != nil {
		fmt.Println("m.Step(2):", err.Error())
	}

	return err
}
