package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//DeleteOrganization is used to delete Organization data from database
func (r *Repository) DeleteOrganization(ID uint) error {
	Organization := nmodel.Organization{}
	err := r.DB.Where("ID = ?", ID).First(&Organization).Delete(&Organization).Error

	if err != nil {
		return err
	}

	return nil
}
