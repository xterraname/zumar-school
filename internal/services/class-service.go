package services

import (
	"zumar-school/internal/models"
	"zumar-school/internal/repository"
)

type ClassService interface {
	GetClasses() ([]models.Class, error)
	GetClass(id uint) (models.Class, error)
	CreateClass(class *models.Class) error
	UpdateClass(class *models.Class) error
	DeleteClass(id uint) error
}

type classService struct {
	classRepo repository.ClassRepository
}

func NewClassService(classRepo repository.ClassRepository) ClassService {
	return &classService{classRepo: classRepo}
}

func (s *classService) GetClasses() ([]models.Class, error) {
	return s.classRepo.GetAll()
}

func (s *classService) GetClass(id uint) (models.Class, error) {
	return s.classRepo.GetByID(id)
}

func (s *classService) CreateClass(class *models.Class) error {
	return s.classRepo.Create(class)
}

func (s *classService) UpdateClass(class *models.Class) error {
	return s.classRepo.Update(class)
}

func (s *classService) DeleteClass(id uint) error {
	return s.classRepo.Delete(id)
}
