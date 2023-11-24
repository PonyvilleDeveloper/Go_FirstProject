package storageSQL

import (
	"app/entity"
	"app/service"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var driver string

var file string

func init() {
	driver = "sqlite3"
	file = "resources/storage.db"
}

func AddCart(item entity.Cart) { //Экв. CreateCart в пакете service
	db, err := sql.Open(driver, file) //Открытие БД
	if err != nil {
		service.Log("\t\t[STORAGE]: Error to open db-file %v\n", err)
	}

	defer db.Close()

	result, err := db.Exec("insert into items (itemId, article, name, price, creator) values")
	if err != nil {
		panic(err)
	}
}

func ChangeCart(id uint32, updts entity.Cart) { //Экв. UpdateCart в пакете service
	db, err := sql.Open(driver, file) //Открытие БД
	if err != nil {
		service.Log("\t\t[STORAGE]: Error to open db-file %v\n", err)
	}
	defer db.Close()
}

func DeleteCart(id uint32) { //Экв. DeleteCart в пакете service
	db, err := sql.Open(driver, file) //Открытие БД
	if err != nil {
		service.Log("\t\t[STORAGE]: Error to open db-file %v\n", err)
	}
	defer db.Close()
}

func GetCartById(id uint32) entity.Cart { //Экв. GetCartById в пакете service
	db, err := sql.Open(driver, file) //Открытие БД
	if err != nil {
		service.Log("\t\t[STORAGE]: Error to open db-file %v\n", err)
	}
	defer db.Close()
}

func GetCartAll() []entity.Cart { //Экв. GetCartAll в пакете service
	db, err := sql.Open(driver, file) //Открытие БД
	if err != nil {
		service.Log("\t\t[STORAGE]: Error to open db-file %v\n", err)
	}
	defer db.Close()
}
