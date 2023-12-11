package db_helper

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//var schema = `
//CREATE TABLE IF NOT EXISTS Users (
//    user_name text NOT NULL PRIMARY KEY,
//    password text,
//    balance text,
//    activity int,
//    bank_country_code text,
//    bank_name text
//);
//CREATE TABLE IF NOT EXISTS Transactions(
//    id serial NOT NULL PRIMARY KEY ,
//    sender_user_name text,
//    sender_balance text,
//    sender_result_balance text,
//    recipient_user_name text,
//    recipient_balance text,
//    recipient_result_balance text,
//    amount text
//
//);`

type DBHelper struct {
	dbConnection *sqlx.DB
}
