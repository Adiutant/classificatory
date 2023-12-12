package db_helper

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	Username string `db:"user_name"`
	Password string `db:"password"`
	Role     string `db:"role"`
}
type Connector struct {
	dbConn *sqlx.DB
}

func Init() (*Connector, error) {
	connStr := "host=postgres port=5432 user=ts password=pass dbname=test sslmode=disable"
	dbMigrate, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(dbMigrate, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return nil, err
	}
	m.Up()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(dbMigrate)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Connector{dbConn: db}, nil
}
func (c *Connector) Check(username string, password string) (bool, error) {

	var u User
	err := c.dbConn.Get(&u, "SELECT * FROM Users WHERE user_name=$1 AND password=$2", username, password)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil

}
