package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ---------here is my all handlers function
func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp Employee                     // variabel for storing the input data
	json.NewDecoder(r.Body).Decode(&emp) // now decode the input json data with the help of json_decoder .
	Database.Create(&emp)                //data created now we have to send  data back to the database in json foramat.
	json.NewEncoder(w).Encode(emp)       //sending encoading data to the database thats means data is created .
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var allEmployees []Employee             //---variable for storing the fetching data .
	Database.Find(&allEmployees)            //---fetchig all record from database .
	json.NewEncoder(w).Encode(allEmployees) //----ecoading the fetch data for showing the requested client.
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(&employee) //----printing the data to database.

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employee) // this update the previous data to new data and save to database.
	Database.Save(&employee)
	json.NewEncoder(w).Encode(&employee)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.Delete(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("Employee is deleted from Database !!")

}
