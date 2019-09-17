package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//CountAsset get one Asset detail from Asset table
func (r *Repository) CountAsset() (*nmodel.StatusCount, error) {
	var total int

	err := r.DB.Table("assets").Count(&total).Error
	if err != nil {
		return nil, err
	}

	var count = &nmodel.StatusCount{
		Name:     "Total Asset",
		Quantity: total,
	}

	return count, nil
}
