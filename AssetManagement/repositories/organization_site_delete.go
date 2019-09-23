package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//DeleteOrganizationSite is used to delete OrganizationSite data from database
func (r *Repository) DeleteOrganizationSite(ID uint) error {
	mOrganization := nmodel.OrganizationSite{}
	err := r.DB.Where("ID = ?", ID).First(&mOrganization).Delete(&mOrganization).Error

	if err != nil {
		return err
	}

	return nil
}
