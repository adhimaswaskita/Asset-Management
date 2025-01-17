package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalAssetByRegion get one Asset detail from Asset table
func (r *Repository) TotalAssetByRegion() (*[]nmodel.Dashboard, error) {
	var mDashboard []nmodel.Dashboard
	err := r.DB.Raw(`
	SELECT organization_regions.name, COUNT(*) as quantity 
		FROM assets 
		RIGHT JOIN organization_sites ON organization_sites.id = assets.site_id 
			RIGHT JOIN organization_regions ON organization_regions.id = organization_sites.region_id
	GROUP BY organization_regions.name 
	ORDER BY quantity DESC;`).Scan(&mDashboard).Error
	if err != nil {
		return nil, err
	}

	return &mDashboard, nil
}
