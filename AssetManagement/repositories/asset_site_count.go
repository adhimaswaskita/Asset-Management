package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//AssetBySite get one Asset detail from Asset table
func (r *Repository) AssetBySite() (*[]nmodel.Dashboard, error) {
	var AssetSummary []nmodel.Dashboard
	err := r.DB.Raw(`
	SELECT organization_sites.name, COUNT(*) as quantity 
		FROM assets 
		RIGHT JOIN organization_sites ON organization_sites.id = assets.site_id 
	GROUP BY organization_sites.name 
	ORDER BY quantity DESC;`).Scan(&AssetSummary).Error
	if err != nil {
		return nil, err
	}

	return &AssetSummary, nil
}
