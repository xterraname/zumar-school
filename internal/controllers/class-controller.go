package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"zumar-school/internal/models"
	"zumar-school/internal/services"

	"github.com/gin-gonic/gin"
)

type ClassController struct {
	classService services.ClassService
}

func NewClassController(classService services.ClassService) *ClassController {
	return &ClassController{classService: classService}
}

func (ctrl *ClassController) GetClasses(c *gin.Context) {
	classes, err := ctrl.classService.GetClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}

func (ctrl *ClassController) GetClass(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	class, err := ctrl.classService.GetClass(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class topilmadi"})
		return
	}
	c.JSON(http.StatusOK, class)
}

func (ctrl *ClassController) CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.classService.CreateClass(&class); err != nil {
		if strings.Contains(err.Error(), "idx_degree_group") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "This combination of degree and group already exists.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, class)
}

func (ctrl *ClassController) UpdateClass(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	class.ID = uint(id)
	if err := ctrl.classService.UpdateClass(&class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
}

// Classni o'chirish
func (ctrl *ClassController) DeleteClass(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	if err := ctrl.classService.DeleteClass(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
