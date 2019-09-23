package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateOrganizationRegion is used to save updated data to database
func (r *Repository) UpdateOrganizationRegion(ID uint, mOrganization *nmodel.OrganizationRegion) (*nmodel.OrganizationRegion, error) {
	var mpOrganization nmodel.OrganizationRegion

	err := r.DB.Model(&mpOrganization).Where("ID = ?", ID).First(&mpOrganization).Updates(mOrganization).Error
	if err != nil {
		return nil, err
	}

	return mOrganization, nil
}
