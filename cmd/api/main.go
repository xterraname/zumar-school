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

func init() {
	config.LoadEnv()
}

func main() {

	router := gin.Default()

	router.SetTrustedProxies(nil)

	api := router.Group("/api")
	{
		api.GET("/ping", PingHandler)

	}

	addr := fmt.Sprintf(":%s", config.GetPort())

	if err := router.Run(addr); err != nil {
		log.Fatalf("Serverni ishga tushirishda xatolik: %v", err)
	}
}
