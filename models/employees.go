package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Position  string  `json:"position"`
	Salary    float64 `json:"salary"`
	FullTime  bool    `json:"full_time"`
}
