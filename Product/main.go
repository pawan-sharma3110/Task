package main

import (
	"log"
	"net/http"
	"product/database"
	"product/handler"
)

func main() {
	db := database.DbIn()
	if db == nil {
		log.Fatal()
		return
	}
	defer db.Close()
	http.HandleFunc("/add", handler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
