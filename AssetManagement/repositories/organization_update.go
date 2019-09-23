package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateOrganization is used to save updated data to database
func (r *Repository) UpdateOrganization(ID uint, mOrganization *nmodel.Organization) (*nmodel.Organization, error) {
	var mpOrganization nmodel.Organization

	err := r.DB.Model(&mpOrganization).Where("ID = ?", ID).First(&mpOrganization).Updates(mOrganization).Error
	if err != nil {
		return nil, err
	}

	return mOrganization, nil
}
