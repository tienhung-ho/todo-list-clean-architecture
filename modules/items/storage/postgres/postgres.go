package postgres

import "gorm.io/gorm"

type postgresStore struct {
	db *gorm.DB
}

func NewPostgresStore(db *gorm.DB) *postgresStore {
	return &postgresStore{db: db}
}
