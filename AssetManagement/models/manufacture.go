package models

import (
	"github.com/jinzhu/gorm"
)

//Manufacture represent manufacture table
type Manufacture struct {
	*gorm.Model
	Name    string `json:"Name"`
	Comment string `json:"Comment"`
}
