package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	uid        string
	platform   string
	level      int
	trio_rank  int
	arena_rank int
}

var DbConnection *sql.DB

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./apex.db")
	defer DbConnection.Close()

}
