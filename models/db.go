package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type UserData struct {
	Uid         string `db:"uid"`
	Platform    string `db:"platform"`
	Level       int    `db:"level"`
	Trio_rank   int    `db:"trio_rank"`
	Arena_rank  int    `db:"arena_rank"`
	Last_update int    `db:"last_update"`
}

func Connect() *sql.DB {
	// Create db client
	db, _ := sql.Open("sqlite3", "./apex.db")
	return db
}

func GetPlayerData(db *sql.DB) []UserData {
	var userData []UserData
	rows, err := db.Query(`SELECT * FROM user_data`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var u UserData
		err := rows.Scan(&u.Uid, &u.Platform, &u.Level, &u.Trio_rank, &u.Arena_rank, &u.Last_update)
		if err != nil {
			fmt.Println(err)
		}
		userData = append(userData, u)

	}

	return userData
}

func UpsertPlayerData(db *sql.DB, u UserData) {
	query := `REPLACE INTO user_data values (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, u.Uid, u.Platform, u.Level, u.Trio_rank, u.Arena_rank, u.Last_update)
	if err != nil {
		log.Fatal(err)
		return
	}
}
