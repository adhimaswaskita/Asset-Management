package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//DeleteOrganization is used to delete Organization data from database
func (r *Repository) DeleteOrganization(ID uint) error {
	mOrganization := nmodel.Organization{}
	err := r.DB.Where("ID = ?", ID).First(&mOrganization).Delete(&mOrganization).Error

	if err != nil {
		return err
	}

	return nil
}
