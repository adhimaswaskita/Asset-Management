package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//GetAllOrganization is used to get all manfacture data from database
func (r *Repository) GetAllOrganization() (Organization []nmodel.Organization, er error) {
	r.DB.Find(&Organization)

	return Organization, nil
}
