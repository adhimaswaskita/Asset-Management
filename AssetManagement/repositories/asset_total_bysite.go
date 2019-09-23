package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalAssetBySite get one Asset detail from Asset table
func (r *Repository) TotalAssetBySite() (*[]nmodel.Dashboard, error) {
	var mDashboard []nmodel.Dashboard
	err := r.DB.Raw(`
	SELECT organization_sites.name, COUNT(*) as quantity 
		FROM assets 
		RIGHT JOIN organization_sites ON organization_sites.id = assets.site_id 
	GROUP BY organization_sites.name 
	ORDER BY quantity DESC;`).Scan(&mDashboard).Error
	if err != nil {
		return nil, err
	}

	return &mDashboard, nil
}
