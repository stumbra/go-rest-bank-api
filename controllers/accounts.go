package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stumbra/go-rest-bank-api/models"
)

func InjectAccountsRoutes(router *gin.RouterGroup) {
	routes := router.Group("./accounts")

	routes.GET("", handleGetAccounts)
	routes.POST("", handleCreateAccount)
	routes.GET("/:id", handleGetAccount)
	routes.DELETE("/:id", handleDeleteAccount)
}

func handleGetAccounts(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, &[]models.Account{})
}

func handleGetAccount(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusFound, &models.Account{})
}

func handleCreateAccount(ctx *gin.Context) {

}

func handleDeleteAccount(ctx *gin.Context) {

}
