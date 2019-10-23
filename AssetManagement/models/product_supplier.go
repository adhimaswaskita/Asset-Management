package models

import (
	"github.com/jinzhu/gorm"
)

//ProductSupplier represent product supplier table
type ProductSupplier struct {
	*gorm.Model
	Name        string `json:"Name"`
	Address     string `json:"Address"`
	City        string `json:"City"`
	State       string `json:"State"`
	Country     string `json:"Country"`
	Zip         int    `json:"Zip"`
	ContactName string `json:"ContactName"`
	Phone       string `json:"Phone"`
	Email       string `json:"Email"`
	URL         string `json:"URL"`
}
