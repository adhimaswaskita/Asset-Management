package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Asset represent asset table
type Asset struct {
	*gorm.Model
	ProductID       int       `json:"productID"`
	SiteID          int       `json:"siteID"`
	ProductStatusID int       `json:"productStatusID"`
	Name            string    `json:"name"`
	PurchaseDate    time.Time `json:"purchaseDate"`
	PurchaseCost    float32   `json:"purchaseCost"`
	Warranty        int       `json:"warranty"`
}
