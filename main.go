package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stumbra/go-rest-bank-api/controllers"
	"github.com/stumbra/go-rest-bank-api/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")

	}

	db := database.NewPostgresStore()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Attach the database instance to the Gin context using middleware
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	v1 := router.Group("/api/v1")

	controllers.InjectAccountsRoutes(v1)

	router.Run(":" + os.Getenv("SERVER_PORT"))
}
