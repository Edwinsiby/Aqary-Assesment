package repository

import "gorm.io/gorm"

func InitDatabase() (gorm.DB, error) {

	return gorm.DB{}, nil
}
