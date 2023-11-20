package main

import (
	"golang-gorm/handler"
	"golang-gorm/middleware"
	"log"
	"net/http"
)

var port string = ":8080"

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/insertOne", handler.HandlePostOne)
	mux.HandleFunc("/insertMultiple", handler.HandlePostMultiple)
	mux.HandleFunc("/insertSelect", handler.HandlePostInsertSelect)
	mux.HandleFunc("/association", handler.HandlePostAssociation)

	var handler http.Handler = mux
	handler = middleware.MiddlewareBasicAuth(handler)

	var server http.Server
	server.Addr = port
	server.Handler = handler

	log.Printf("server running on localhost%s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

}
