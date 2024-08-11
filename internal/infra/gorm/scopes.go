package gorm

import "gorm.io/gorm"

func FilterSoftDelete(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}
