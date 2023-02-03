package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/allEmployees", GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{eid}", GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employee/{eid}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{eid}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
