package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateOrganization is used to save updated data to database
func (r *Repository) UpdateOrganization(ID uint, Organization *nmodel.Organization) (*nmodel.Organization, error) {
	var organizationParam nmodel.Organization

	err := r.DB.Model(&organizationParam).Where("ID = ?", ID).First(&organizationParam).Updates(Organization).Error
	if err != nil {
		return nil, err
	}

	return Organization, nil
}
