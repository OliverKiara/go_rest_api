package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//contact struct

type ContactModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

var contacts []ContactModel

//get all contacts
func allContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

//get single contact
func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range contacts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&ContactModel{})

}

//create contact
func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact ContactModel
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = strconv.Itoa(rand.Intn(100))
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

//update contact
func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range contacts {
		if item.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			var contact ContactModel
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.ID = params["id"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

//delete contact
func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range contacts {
		if item.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

func main() {

	mux := mux.NewRouter()

	contacts = append(contacts, ContactModel{ID: "1", Name: "Nathan", Contact: "0797378891"})
	contacts = append(contacts, ContactModel{ID: "2", Name: "Jane", Contact: "0788234576"})
	contacts = append(contacts, ContactModel{ID: "3", Name: "John", Contact: "0723467589"})

	mux.HandleFunc("/api/contacts", allContacts).Methods("GET")
	mux.HandleFunc("/api/contact/{id}", getContact).Methods("GET")
	mux.HandleFunc("/api/contact", createContact).Methods("POST")
	mux.HandleFunc("/api/contact/{id}", updateContact).Methods("PUT")
	mux.HandleFunc("/api/contact/{id}", deleteContact).Methods("DELETE")

	log.Println("Server starting on :4040")
	err := http.ListenAndServe(":4040", mux)
	log.Fatal(err)

}
