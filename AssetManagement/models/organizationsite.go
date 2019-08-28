package models

import (
	"github.com/jinzhu/gorm"
)

//OrganizationSite represent organization_site table
type OrganizationSite struct {
	*gorm.Model
	RegionID    int    `json:"regionID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
}
