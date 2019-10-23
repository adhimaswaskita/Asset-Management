package models

import (
	"github.com/jinzhu/gorm"
)

//ProductStatus represent product_status table
type ProductStatus struct {
	*gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
