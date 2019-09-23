package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganizationRegion is used to get all organization region data from database
func (r *Repository) GetAllOrganizationRegion() (mOrganization []nmodel.OrganizationRegion, er error) {
	r.DB.Find(&mOrganization)

	return mOrganization, nil
}
