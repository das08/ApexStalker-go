package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
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

func GetPlayers(db *sql.DB) []User {
	var UserList []User
	rows, err := db.Query(`SELECT * FROM user_data`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Uid, &u.Platform, &u.Level, &u.Trio_rank, &u.Arena_rank, &u.Last_update)
		if err != nil {
			fmt.Println(err)
		}
		UserList = append(UserList, u)

	}

	return UserList
}

func UpsertPlayerData(db *sql.DB, u User) {
	query := `REPLACE INTO user_data values (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, u.Uid, u.Platform, u.Level, u.Trio_rank, u.Arena_rank, u.Last_update)
	if err != nil {
		log.Fatal(err)
		return
	}
}
