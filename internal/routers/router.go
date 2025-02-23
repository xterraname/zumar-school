package routers

import (
	ct "zumar-school/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(studentCt *ct.StudentController, classCt *ct.ClassController) *gin.Engine {
	router := gin.Default()

	studentRoutes := router.Group("/students")
	{
		studentRoutes.GET("/", studentCt.GetStudents)
		studentRoutes.GET("/:id", studentCt.GetStudent)
		studentRoutes.POST("/", studentCt.CreateStudent)
		studentRoutes.PUT("/:id", studentCt.UpdateStudent)
		studentRoutes.DELETE("/:id", studentCt.DeleteStudent)
	}

	classRoutes := router.Group("/classes")
	{
		classRoutes.GET("/", classCt.GetClasses)
		classRoutes.GET("/:id", classCt.GetClass)
		classRoutes.POST("/", classCt.CreateClass)
		classRoutes.PUT("/:id", classCt.UpdateClass)
		classRoutes.DELETE("/:id", classCt.DeleteClass)
	}

	return router
}
