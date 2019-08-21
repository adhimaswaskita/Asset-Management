package models

import (
	"github.com/jinzhu/gorm"
)

//ProductType represent the product_type table
type ProductType struct {
	*gorm.Model
	Name        string `json:"Name"`
	Type        string `json:"Type"`
	Category    string `json:"Category"`
	Description string `json:"Description"`
}
