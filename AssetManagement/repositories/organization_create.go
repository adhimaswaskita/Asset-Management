package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganization is used to insert Organization data to database
func (r *Repository) CreateOrganization(Organization *nmodel.Organization) (*nmodel.Organization, error) {
	err := r.DB.Create(&Organization).Error
	if err != nil {
		return nil, err
	}

	return Organization, nil
}
