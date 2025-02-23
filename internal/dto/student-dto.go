package dto

import "zumar-school/internal/models"

type CreateStudentRequest struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=32"`
	LastName  string `json:"last_name" binding:"required,min=3,max=32"`
	MidName   string `json:"mid_name" binding:"required,min=3,max=32"`
	ClassID   uint   `json:"class_id" binding:"required"`
}

type StudentResponse struct {
	ID        uint          `json:"id"`
	FirstName string        `json:"first_name" binding:"required,min=3,max=32"`
	LastName  string        `json:"last_name" binding:"required,min=3,max=32"`
	MidName   string        `json:"mid_name" binding:"required,min=3,max=32"`
	Class     ClassResponse `json:"class"`
}

func NewStudentResponse(student models.Student) StudentResponse {
	return StudentResponse{
		ID:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		MidName:   student.MidName,
		Class:     ClassResponse(student.Class),
	}
}
