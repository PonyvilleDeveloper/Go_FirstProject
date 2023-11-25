package db

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB
var migrate = make([]func(), 0)

func Add(mF func()) {
	migrate = append(migrate, mF)
}
func DB() *gorm.DB {
	return database
}

func Migrate() {
	for _, f := range migrate {
		f()
	}
}
