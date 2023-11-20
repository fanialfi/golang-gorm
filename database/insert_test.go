package database

import (
	"golang-gorm/model"
	"testing"
)

func TestInsert(t *testing.T) {
	db := Connect()
	db.AutoMigrate(model.User{})

	data := model.User{
		Name:  "Cek Aja",
		Email: "mail@example.com",
		Age:   20,
	}

	err := InsertWithSelect(db, data, []string{"Name", "Email", "Age"})
	if err != nil {
		t.Logf(err.Error())
		t.Failed()
	}
}
