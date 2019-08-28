package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//UpdateOrganizationRegion is used to save updated data to database
func (r *Repository) UpdateOrganizationRegion(ID uint, organizationRegion *nmodel.OrganizationRegion) (*nmodel.OrganizationRegion, error) {
	var organizationRegionParam nmodel.OrganizationRegion

	err := r.DB.Model(&organizationRegionParam).Where("ID = ?", ID).First(&organizationRegionParam).Updates(organizationRegion).Error
	if err != nil {
		return nil, err
	}

	return organizationRegion, nil
}
