package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Brennon-Oliveira/academiQ/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	//env or default 7301
	port := os.Getenv("PORT")
	if port == "" {
		port = "7301"
	}

	log.Println("Starting academiQ-api")

	database.Connect()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run(port)
}
