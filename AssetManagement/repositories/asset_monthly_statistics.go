package repositories

import nmodel "github.com/adhimaswaskita/AssetManagement/models"

//AssetStatistics get one Asset detail from Asset table
func (r *Repository) AssetStatistics(year int) (*[]nmodel.Statistics, error) {
	var mStatistics []nmodel.Statistics
	err := r.DB.Raw(`
	SELECT EXTRACT (MONTH FROM assets.purchase_date) AS month, COUNT(*) as quantity 
		FROM assets 
	WHERE EXTRACT(YEAR FROM assets.purchase_date) = ?
	GROUP BY month
	ORDER BY month ASC;`, year).Scan(&mStatistics).Error
	if err != nil {
		return nil, err
	}

	return &mStatistics, nil
}
