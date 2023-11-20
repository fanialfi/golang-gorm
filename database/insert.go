package database

import (
	"golang-gorm/model"

	"gorm.io/gorm"
)

func InsertData[V model.User | []*model.User](db *gorm.DB, data V) error {
	// db.Create() juga bisa digunakan untuk insert multiple record
	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func InsertWithSelect(db *gorm.DB, data model.User, items []string) error {
	result := db.Select(items).Create(&data)
	// result := db.Omit(items).Create(&data) // jika menggunakan Omit field / column yang ada di parameter items akan dihilangkan
	if result.Error != nil {
		return result.Error
	}

	return nil
}
