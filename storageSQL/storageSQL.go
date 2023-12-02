package storageSQL

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB
var migrations = make([]func(), 0)

func AddMigration(mF func()) {
	migrations = append(migrations, mF)
}

func Migrate() {
	for _, f := range migrations {
		f()
	}
}
