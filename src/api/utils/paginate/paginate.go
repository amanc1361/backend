package paginate

import (
	"back-account/src/api/utils/pagination"
	"math"

	"gorm.io/gorm"
)

func Paginate(value int64, pagination *pagination.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	pagination.TotalRows = value
	totalPage := int(math.Ceil((math.Ceil(float64(value)) / math.Ceil(float64(pagination.GetLimit())))))
	pagination.TotalPages = totalPage
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
