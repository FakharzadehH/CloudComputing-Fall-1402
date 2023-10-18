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

func (r *Repository) Upsert(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *Repository) GetByID(id int, user *domain.User) error {
	return r.db.First(user, id).Error
}

func (r *Repository) GetByNationalID(nationalID string, user *domain.User) error {

	return r.db.First(user, "national_id = ?", nationalID).Error
}
