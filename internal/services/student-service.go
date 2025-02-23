package services

import (
	"zumar-school/internal/models"
	"zumar-school/internal/repository"
)

type StudentService interface {
	GetStudents() ([]models.Student, error)
	GetStudent(id uint) (models.Student, error)
	CreateStudent(student *models.Student) error
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

type studentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &studentService{studentRepo: studentRepo}
}

func (s *studentService) GetStudents() ([]models.Student, error) {
	return s.studentRepo.GetAll()
}

func (s *studentService) GetStudent(id uint) (models.Student, error) {
	return s.studentRepo.GetByID(id)
}

func (s *studentService) CreateStudent(student *models.Student) error {
	return s.studentRepo.Create(student)
}

func (s *studentService) UpdateStudent(student *models.Student) error {
	return s.studentRepo.Update(student)
}

func (s *studentService) DeleteStudent(id uint) error {
	return s.studentRepo.Delete(id)
}
