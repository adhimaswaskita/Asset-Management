package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//AssetIT get one Asset detail from Asset table
func (r *Repository) AssetIT() (*[]nmodel.Dashboard, error) {
	var AssetSummary []nmodel.Dashboard
	err := r.DB.Raw(`
	SELECT assets.name, COUNT(*) as quantity 
		FROM assets 
		INNER JOIN products ON products.id = assets.product_id 
		INNER JOIN product_types ON product_types.id = products.product_type_id
	WHERE product_types.category = 'IT'
	GROUP BY assets.name 
	ORDER BY quantity DESC;`).Scan(&AssetSummary).Error
	if err != nil {
		return nil, err
	}

	return &AssetSummary, nil
}
