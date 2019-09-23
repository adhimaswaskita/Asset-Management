package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganizationSite is used to get all organization Site data from database
func (r *Repository) GetAllOrganizationSite() (mOrganization []nmodel.OrganizationSite, er error) {
	r.DB.Find(&mOrganization)

	return mOrganization, nil
}
