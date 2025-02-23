package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"zumar-school/internal/dto"
	"zumar-school/internal/models"
	"zumar-school/internal/services"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	studentService services.StudentService
}

func NewStudentController(studentService services.StudentService) *StudentController {
	return &StudentController{studentService: studentService}
}

func (ctrl *StudentController) GetStudents(c *gin.Context) {
	students, err := ctrl.studentService.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (ctrl *StudentController) GetStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	student, err := ctrl.studentService.GetStudent(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (ctrl *StudentController) CreateStudent(c *gin.Context) {

	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		MidName:   req.MidName,
		ClassID:   req.ClassID,
	}

	if err := ctrl.studentService.CreateStudent(&student); err != nil {
		if strings.Contains(err.Error(), "idx_student_fullname") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "This combination of firstname, lastname and midname already exists.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.NewStudentResponse(student))
}

func (ctrl *StudentController) UpdateStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Notog'ri student ID"})
		return
	}
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student.ID = uint(id)
	if err := ctrl.studentService.UpdateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (ctrl *StudentController) DeleteStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	if err := ctrl.studentService.DeleteStudent(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student successfully deleted"})
}
