package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//ID      int    `gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Persons int    `json:"persons"`
}
