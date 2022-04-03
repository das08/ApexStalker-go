package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type UserData struct {
	Uid        string `db:"uid"`
	Platform   string `db:"platform"`
	Stats      UserDataDetail
	LastUpdate int `db:"last_update"`
}

type UserDataDetail struct {
	Level     int `db:"level"`
	TrioRank  int `db:"trio_rank"`
	ArenaRank int `db:"arena_rank"`
}

func Connect() *sql.DB {
	// Create db client
	db, err := sql.Open("sqlite3", "./apex.db")
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
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
		err := rows.Scan(&u.Uid, &u.Platform, &u.Stats.Level, &u.Stats.TrioRank, &u.Stats.ArenaRank, &u.LastUpdate)
		if err != nil {
			fmt.Println(err)
		}
		userData = append(userData, u)

	}

	return userData
}

func UpsertPlayerData(db *sql.DB, u UserData) {
	query := `REPLACE INTO user_data values (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, u.Uid, u.Platform, u.Stats.Level, u.Stats.TrioRank, u.Stats.ArenaRank, u.LastUpdate)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func UpdatePlayerData(db *sql.DB, userID string, ud UserDataDetail) {
	timestamp := time.Now().Unix()
	query := `UPDATE user_data set level=?,trio_rank=?,arena_rank=?,last_update=? WHERE uid=?`
	_, err := db.Exec(query, ud.Level, ud.TrioRank, ud.ArenaRank, timestamp, userID)
	if err != nil {
		log.Fatal(err)
		return
	}
}
