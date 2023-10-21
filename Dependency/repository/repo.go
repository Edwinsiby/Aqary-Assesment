package repository

import "gorm.io/gorm"

type Repository struct {
	Db gorm.DB
}

type RepositoryMethods interface {
}

func InitRepository(db gorm.DB) Repository {
	return Repository{
		Db: db,
	}
}
