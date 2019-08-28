package models

import (
	"github.com/jinzhu/gorm"
)

//Organization represent organization table
type Organization struct {
	*gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	City        string `json:"city"`
	Address     string `json:"address"`
	PostalCode  int    `json:"postalcode"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	URL         string `json:"url"`
}
