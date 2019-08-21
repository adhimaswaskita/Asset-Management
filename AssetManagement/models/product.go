package models

import (
	"github.com/jinzhu/gorm"
)

//Product represent the product table
type Product struct {
	*gorm.Model
	Name              string `json:"Name"`
	ManufactureID     uint   `json:"ManufactureID"`
	ProductTypeID     uint   `json:"ProductTypeID"`
	ProductSupplierID uint   `json:"ProductSupplierID"`
	Manufacturer      string `json:"Manufacturer"`
	PartNo            int    `json:"PartNo"`
}
