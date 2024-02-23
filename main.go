package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stumbra/go-rest-bank-api/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")

	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	gin.SetMode(gin.ReleaseMode)

	v1 := router.Group("/api/v1")

	controllers.InjectAccountsRoutes(v1)

	port := os.Getenv("SERVER_PORT")

	err = router.Run(":" + port)

	if err != nil {
		fmt.Print("Server read at http://localhost:", port, "/api/v1")
		fmt.Println()
	}

}

// pgadmin
// 5432
