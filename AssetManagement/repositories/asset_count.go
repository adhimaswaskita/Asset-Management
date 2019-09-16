package repositories

//CountAsset get one Asset detail from Asset table
func (r *Repository) CountAsset() (int, error) {
	// var Asset nmodel.Asset
	var total int

	err := r.DB.Table("assets").Count(&total).Error
	if err != nil {
		return -1, err
	}

	return total, nil
}
