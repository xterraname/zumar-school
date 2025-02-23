package main

import (
	"fmt"
	"log"
	"net/http"
	"zumar-school/internal/config"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

var cfg *config.Config

func init() {
	config.LoadENV()
	config.ConnectDB()
	cfg = config.Cfg
}

func main() {

	router := gin.Default()

	router.SetTrustedProxies(nil)

	api := router.Group("/api")
	{
		api.GET("/ping", PingHandler)

	}

	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
