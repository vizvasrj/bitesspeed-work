package main

import (
	"log"
	"src/pkg/db"
	"src/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := db.GetConnection()
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	go routes.SetupRouter(router)
	router.Run(":8080")

}
