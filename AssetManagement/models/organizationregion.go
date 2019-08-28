package models

import (
	"github.com/jinzhu/gorm"
)

//OrganizationRegion represent organization_region 
type OrganizationRegion struct {
	*gorm.Model
	OrganizationID int    `json:"organizationID"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}
