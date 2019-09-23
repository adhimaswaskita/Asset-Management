package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//TotalAsset get one Asset detail from Asset table
func (r *Repository) TotalAsset() (*nmodel.Dashboard, error) {
	var total int

	err := r.DB.Table("assets").Count(&total).Error
	if err != nil {
		return nil, err
	}

	var result = &nmodel.Dashboard{
		Name:     "Total Asset",
		Quantity: total,
	}

	return result, nil
}
