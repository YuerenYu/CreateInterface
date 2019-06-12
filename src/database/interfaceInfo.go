package database

import (
	"CreateInterface/src/utils"
	"database/sql"
	"path/filepath"
	"time"
)

type InterfaceInfo struct {
	Id         int
	Pattern    string
	Content    string
	InsertTime time.Time
}

const database = "src/database/database.db"

func OpenDB() (*sql.DB, error) {
	filePath, _ := filepath.Abs(database)
	return sql.Open("sqlite3", filePath)
}

func InsertInfo(db *sql.DB, info InterfaceInfo) int64 {
	stmt, err := db.Prepare("INSERT INTO interface_info(pattern, content, inserttime) VALUES (?,?,?)")
	utils.CheckError(err)

	res, err := stmt.Exec(info.Pattern, info.Content, time.Now())
	utils.CheckError(err)

	id, err := res.LastInsertId()
	utils.CheckError(err)

	return id
}
