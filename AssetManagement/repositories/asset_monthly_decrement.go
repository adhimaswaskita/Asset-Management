package repositories

import (
	"fmt"

	nmodel "github.com/adhimaswaskita/AssetManagement/models"
)

//AssetDecrementStatistics get one Asset detail from Asset table
func (r *Repository) AssetDecrementStatistics(bulan int) (int, error) {
	var mDecrement nmodel.Decrement
	var count []int
	query := fmt.Sprintf(`
	select count(*) as decrease from assets where EXTRACT (MONTH FROM purchase_date) = %d;`, bulan)

	err := r.DB.Raw(query).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	mDecrement.Month = fmt.Sprintf("%v", bulan)
	mDecrement.Decrease = count[0]

	return mDecrement.Decrease, nil
}
