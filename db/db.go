package db

import (
	"database/sql"
	_ "modernc.org/sqlite"

)

var DB *sql.DB

func InitDB(){
	var err error
	DB, err = sql.Open("sqlite","api.db")
	if err != nil{
		panic("Could not connect to database.")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables(){

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil{
		panic("Could not create users table")
	}

	createMealsTable := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES user(id)
	)
	`
	_, err = DB.Exec(createMealsTable)
	if err != nil{
		panic("Could not create events table.")
	} 

	createConfirmTable := `
	CREATE TABLE IF NOT EXISTS meal_confirmations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		meal_id INTEGER,
		status TEXT,
		UNIQUE(user_id, meal_id)
	);`
	_, err = DB.Exec(createConfirmTable)
	if err != nil {
		panic("Could not create meal confirmations table")
	}
}