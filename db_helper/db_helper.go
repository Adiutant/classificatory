package db_helper

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	Username string `db:"user_name"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func Check(username string, password string) (bool, error) {
	connStr := "host=localhost port=5432 user=ts password=pass dbname=test sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return false, err
	}
	var u User
	err = db.Get(&u, "SELECT * FROM Users WHERE user_name=$1 AND password=$2", username, password)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil

}
