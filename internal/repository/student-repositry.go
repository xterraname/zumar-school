package repository

import (
	"zumar-school/internal/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	GetAll() ([]models.Student, error)
	GetByID(id uint) (models.Student, error)
	Create(student *models.Student) error
	Update(student *models.Student) error
	Delete(id uint) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) GetAll() ([]models.Student, error) {
	var students []models.Student
	if err := r.db.Preload("Class").Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepository) GetByID(id uint) (models.Student, error) {
	var student models.Student
	if err := r.db.Preload("Class").First(&student, id).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (r *studentRepository) Create(student *models.Student) error {
	if err := r.db.Create(student).Error; err != nil {
		return err
	}
	return r.db.Preload("Class").First(student, student.ID).Error
}

func (r *studentRepository) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}
