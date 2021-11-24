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
)

func connectToDB() {

	// Capture connection properties.
	/*cfg := m.Config{
		User:   "root",
		Passwd: "mysql",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: database,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Cannot connect to database!")
	}

	fmt.Println("Connected!")
	return db
	*/
}

func Run() error {
	// Capture connection properties.
	/*cfg := m.Config{
		User:   "root",
		Passwd: "mysql",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: database,
	}*/

	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/golendar?multiStatements=true")
	if err != nil {
		panic(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Cannot connect to database!")
	}

	fmt.Println("Connected!")

	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
	}

	m.Steps(2)
	return err
}
