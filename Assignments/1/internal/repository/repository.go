package repository

import (
	"github.com/FakharzadehH/CloudComputing-Fall-1402/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetByNationalID(nationalID int) (domain.User, error) {
	return domain.User{}, nil
}
