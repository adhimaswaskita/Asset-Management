package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CreateOrganizationSite is used to insert organizationSite data to database
func (r *Repository) CreateOrganizationSite(organizationSite *nmodel.OrganizationSite) (*nmodel.OrganizationSite, error) {
	err := r.DB.Create(&organizationSite).Error
	if err != nil {
		return nil, err
	}

	return organizationSite, nil
}
