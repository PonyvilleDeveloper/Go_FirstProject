package storageSQL

import (
	"app/service"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB { //Открытие БД
	db, err := sql.Open("sqlite3", "resources/storage.db")
	if err != nil {
		service.Log("\t\t[STORAGE]: Error to open db-file %v\n", err)
	}
	return db
}

func Dispose(db *sql.DB) { //Освобождение БД
	db.Close()
}
