package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetRecentInsertion get one Asset detail from Asset table
func (r *Repository) AssetRecentInsertion() (*[]nmodel.RecentInsertion, error) {
	var mRecentInsertion []nmodel.RecentInsertion
	err := r.DB.Raw(`
	SELECT product_types.name as category, assets.name as name, products.part_no as code, 
		to_char(assets.created_at, 'HH12:MIAM,DD/MM/YYYY') as time
		FROM assets INNER JOIN products ON assets.product_id = products.id
		INNER JOIN product_types ON products.product_type_id = product_types.id
	ORDER BY time DESC
	LIMIT 5;`).Scan(&mRecentInsertion).Error
	if err != nil {
		return nil, err
	}

	return &mRecentInsertion, nil
}
