package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//ITAssetSummary get one Asset detail from Asset table
func (r *Repository) ITAssetSummary() (*[]nmodel.Dashboard, error) {
	var mDashboard []nmodel.Dashboard
	err := r.DB.Raw(`
	SELECT assets.name, COUNT(*) as quantity 
		FROM assets 
		INNER JOIN products ON products.id = assets.product_id 
		INNER JOIN product_types ON product_types.id = products.product_type_id
	WHERE product_types.category = 'IT'
	GROUP BY assets.name 
	ORDER BY quantity DESC;`).Scan(&mDashboard).Error
	if err != nil {
		return nil, err
	}

	return &mDashboard, nil
}
