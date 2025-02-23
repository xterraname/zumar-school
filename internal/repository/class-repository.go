package repository

import (
	"zumar-school/internal/models"

	"gorm.io/gorm"
)

type ClassRepository interface {
	GetAll() ([]models.Class, error)
	GetByID(id uint) (models.Class, error)
	Create(class *models.Class) error
	Update(class *models.Class) error
	Delete(id uint) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) GetAll() ([]models.Class, error) {
	var classes []models.Class
	if err := r.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepository) GetByID(id uint) (models.Class, error) {
	var class models.Class
	if err := r.db.First(&class, id).Error; err != nil {
		return class, err
	}
	return class, nil
}

func (r *classRepository) Create(class *models.Class) error {
	return r.db.Create(class).Error
}

func (r *classRepository) Update(class *models.Class) error {
	return r.db.Save(class).Error
}

func (r *classRepository) Delete(id uint) error {
	return r.db.Delete(&models.Class{}, id).Error
}
