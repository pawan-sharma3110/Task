package handler

import (
	"encoding/json"
	"net/http"
	"product/controller"
	"product/database"
	"product/model"
	"time"

	"github.com/google/uuid"
)

var db = database.DbIn()

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "Post" {
		http.Error(w, "Only post methord required", http.StatusBadRequest)
		return
	}
	var newProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newProduct.ID = uuid.New()
	newProduct.AddedOn = time.Now()
	res, err := controller.InsertProduct(db, newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
