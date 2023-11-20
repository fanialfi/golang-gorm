package handler

import (
	"encoding/json"
	"fmt"
	"golang-gorm/database"
	"golang-gorm/lib"
	"golang-gorm/model"
	"log"
	"net/http"
)

func HandlePostOne(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Accept POST method", http.StatusBadRequest)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("error on decoding data from request")
		return
	}

	db := database.Connect()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("error on run migrate model")
		return
	}

	err = database.InsertData[model.User](db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("error on insert data to database")
		return
	}

	lib.OutputJSON(w, user)
}

func HandlePostMultiple(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Accept POST method", http.StatusBadRequest)
		return
	}

	var user []*model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("error on decoding data from request")
		return
	}

	db := database.Connect()
	err = db.AutoMigrate(model.User{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("error on run migrate model")
		return
	}

	// db.Create() juga bisa digunakan untuk insert multiple record
	err = database.InsertData[[]*model.User](db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("error on insert data to database")
		return
	}

	lib.OutputJSON(w, user)
}

func HandlePostInsertSelect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Accept POST method", http.StatusBadRequest)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("%#v\n", user)

	db := database.Connect()
	db.Debug()
	err = db.AutoMigrate(model.User{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = database.InsertWithSelect(db, user, []string{"Name", "Email", "Age"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lib.OutputJSON(w, user)
}

func HandlePostAssociation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only accept POST method", http.StatusBadRequest)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db := database.Connect()
	db.Debug()
	err = db.AutoMigrate(model.User{}, model.CreditCard{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%#v\n", user)
	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	lib.OutputJSON(w, user)
}
