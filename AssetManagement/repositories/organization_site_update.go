package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateOrganizationSite is used to save updated data to database
func (r *Repository) UpdateOrganizationSite(ID uint, organizationSite *nmodel.OrganizationSite) (*nmodel.OrganizationSite, error) {
	var organizationSiteParam nmodel.OrganizationSite

	err := r.DB.Model(&organizationSiteParam).Where("ID = ?", ID).First(&organizationSiteParam).Updates(organizationSite).Error
	if err != nil {
		return nil, err
	}

	return organizationSite, nil
}
