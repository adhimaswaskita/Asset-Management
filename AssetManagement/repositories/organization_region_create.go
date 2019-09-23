package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganizationRegion is used to insert organizationRegion data to database
func (r *Repository) CreateOrganizationRegion(mOrganization *nmodel.OrganizationRegion) (*nmodel.OrganizationRegion, error) {
	err := r.DB.Create(&mOrganization).Error
	if err != nil {
		return nil, err
	}

	return mOrganization, nil
}
