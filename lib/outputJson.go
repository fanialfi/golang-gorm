package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

func OutputJSON(w http.ResponseWriter, data any) {
	res, err := json.Marshal(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("error on encode data to json")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
