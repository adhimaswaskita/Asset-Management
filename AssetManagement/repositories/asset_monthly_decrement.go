package repositories

import (
	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetDecrementStatistics get one Asset detail from Asset table
func (r *Repository) AssetDecrementStatistics(bulan int) (int, error) {
	var mDecrement nmodel.Decrement
	err := r.DB.Raw(`
		SELECT EXTRACT (MONTH FROM deleted_at) as month, 
			COUNT(*) 
			FROM assets AS decrease 
		WHERE deleted_at IS NOT NULL AND 
		EXTRACT (MONTH FROM purchase_date) = ?
		GROUP BY month;`, bulan).Scan(&mDecrement).Error
	if err != nil {
		return 0, err
	}

	return mDecrement.Decrease, nil
}
