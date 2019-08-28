package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganizationSite is used to get all organization Site data from database
func (r *Repository) GetAllOrganizationSite() (organizationSite []nmodel.OrganizationSite, er error) {
	r.DB.Find(&organizationSite)

	return organizationSite, nil
}
