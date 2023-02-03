package main

import "gorm.io/gorm"

// ----this is the entity
type Employee struct {
	gorm.Model
	EmpName   string  `json:"empname"`
	EmpSalary float64 `json:"empsalary"`
	Email     string  `json:"email"`
}
