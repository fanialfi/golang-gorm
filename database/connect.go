package database

import (
	"fmt"
	"golang-gorm/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=golang_gorm port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error occurd on open database : %s\n", err.Error()))
	}

	db.AutoMigrate(model.User{})
}
