package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganizationRegion is used to get all organization region data from database
func (r *Repository) GetAllOrganizationRegion() (organizationRegion []nmodel.OrganizationRegion, er error) {
	r.DB.Find(&organizationRegion)

	return organizationRegion, nil
}
