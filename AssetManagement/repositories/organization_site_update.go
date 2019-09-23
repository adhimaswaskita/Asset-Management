package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateOrganizationSite is used to save updated data to database
func (r *Repository) UpdateOrganizationSite(ID uint, mOrganization *nmodel.OrganizationSite) (*nmodel.OrganizationSite, error) {
	var mpOrganization nmodel.OrganizationSite

	err := r.DB.Model(&mpOrganization).Where("ID = ?", ID).First(&mpOrganization).Updates(mOrganization).Error
	if err != nil {
		return nil, err
	}

	return mOrganization, nil
}
