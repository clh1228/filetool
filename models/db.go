package models

import (
	database "filetool/database"

	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return database.GetDataBase()
}
