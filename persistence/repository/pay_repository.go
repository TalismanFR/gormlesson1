package repository

import (
	"github.com/jinzhu/gorm"
	"gormlesson/persistence/model"
)

type PayRepository struct {
	db *gorm.DB
}

func NewPayRepository(db *gorm.DB) *PayRepository {
	return &PayRepository{db: db}
}

func (r *PayRepository) Save(pay *model.Pay) error {
	return r.db.Create(pay).Error
}
