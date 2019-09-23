package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganization is used to get all manfacture data from database
func (r *Repository) GetAllOrganization() (mOrganization []nmodel.Organization, er error) {
	r.DB.Find(&mOrganization)

	return mOrganization, nil
}
