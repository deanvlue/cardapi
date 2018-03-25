package main

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"net/http"
	"time"
	"log"

	"github.com/gorilla/mux"
)

// FindCard :  Regresa una tarjeta con sus correspondientes imágenes
func FindCard(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "por implementar")
}

// UpdateCard : Actualiza tarjeta
func UpdateCard(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "por implementar")
}

// DeleteCard : Borra tarjeta de la base de datos
func DeleteCard(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "por implementar")
}

// NewCard : Crea una nueva tarjeta y la guarda a la base de datos
func NewCard(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "por implementar")
}

// AllCards : Regresa todas las tarjetas
func AllCards(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "por implementar")
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cards", AllCards).Methods("GET")
	r.HandleFunc("/cards", NewCard).Methods("POST")
	r.HandleFunc("/cards", UpdateCard).Methods("PUT")
	r.HandleFunc("/cards", DeleteCard).Methods("DELETE")
	r.HandleFunc("/cards/{bin}", FindCard).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(time.Now(), err)
	}
}

// Card : Struct to send back the cards
type Card struct {
	ID			bson.ObjectId `bson:"_id" json:"id"`
	Promo      string `json:"promo"`
	Inicialbin int32 `json:"inicialbin"`
	Finalbin   int32 `json:"finalbin"`
	Active     bool   `json:"active"`
	Imageurl   []struct {
		ImageType string `json:"imageType"`
		URI       string `json:"uri"`
	} `json:"imageurl"`
}