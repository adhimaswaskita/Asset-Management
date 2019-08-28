package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganizationRegion is used to insert organizationRegion data to database
func (r *Repository) CreateOrganizationRegion(organizationRegion *nmodel.OrganizationRegion) (*nmodel.OrganizationRegion, error) {
	err := r.DB.Create(&organizationRegion).Error
	if err != nil {
		return nil, err
	}

	return organizationRegion, nil
}
