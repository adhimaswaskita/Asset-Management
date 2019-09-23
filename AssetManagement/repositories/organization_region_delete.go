package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//DeleteOrganizationRegion is used to delete OrganizationRegion data from database
func (r *Repository) DeleteOrganizationRegion(ID uint) error {
	mOrganization := nmodel.OrganizationRegion{}
	err := r.DB.Where("ID = ?", ID).First(&mOrganization).Delete(&mOrganization).Error

	if err != nil {
		return err
	}

	return nil
}
