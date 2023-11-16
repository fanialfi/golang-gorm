package main

import (
	"golang-gorm/database"
	"log"
	"time"
)

func main() {
	start := time.Now()
	log.Println(start, start.Format(time.RFC3339))

	database.Connect()

	end := time.Since(start)
	log.Printf("running on %s seconds\n", end.String())
}
