package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//DeleteOrganizationRegion is used to delete OrganizationRegion data from database
func (r *Repository) DeleteOrganizationRegion(ID uint) error {
	OrganizationRegion := nmodel.OrganizationRegion{}
	err := r.DB.Where("ID = ?", ID).First(&OrganizationRegion).Delete(&OrganizationRegion).Error

	if err != nil {
		return err
	}

	return nil
}
