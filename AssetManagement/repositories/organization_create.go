package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganization is used to insert Organization data to database
func (r *Repository) CreateOrganization(mOrganization *nmodel.Organization) (*nmodel.Organization, error) {
	err := r.DB.Create(&mOrganization).Error
	if err != nil {
		return nil, err
	}

	return mOrganization, nil
}
