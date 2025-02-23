package main

import (
	"fmt"
	"log"
	"net/http"
	"zumar-school/internal/config"
	"zumar-school/internal/controllers"
	"zumar-school/internal/repository"
	"zumar-school/internal/routers"
	"zumar-school/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

var cfg *config.Config
var DB *gorm.DB

func init() {
	config.LoadENV()
	config.ConnectDB()
	cfg = config.Cfg
	DB = config.DB
}

func main() {

	studentRepo := repository.NewStudentRepository(DB)
	classRepo := repository.NewClassRepository(DB)

	studentService := services.NewStudentService(studentRepo)
	classService := services.NewClassService(classRepo)

	studentController := controllers.NewStudentController(studentService)
	classController := controllers.NewClassController(classService)

	router := routers.SetupRouter(studentController, classController)
	router.SetTrustedProxies(nil)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
