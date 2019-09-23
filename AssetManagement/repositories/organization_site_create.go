package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganizationSite is used to insert organizationSite data to database
func (r *Repository) CreateOrganizationSite(mOrganization *nmodel.OrganizationSite) (*nmodel.OrganizationSite, error) {
	err := r.DB.Create(&mOrganization).Error
	if err != nil {
		return nil, err
	}

	return mOrganization, nil
}
