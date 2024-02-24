package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stumbra/go-rest-bank-api/database"
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

	db := ctx.MustGet("db").(*database.PostgresStore)

	accounts, err := db.GetAccounts()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	ctx.IndentedJSON(http.StatusOK, accounts)
}

func handleGetAccount(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": http.StatusBadRequest})
	}

	db := ctx.MustGet("db").(*database.PostgresStore)

	account, err := db.GetAccount(id)

	fmt.Println(account, err)

	if err != nil {
		message := fmt.Sprintf("Account with ID - %d not found", id)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": message, "status": http.StatusNotFound})
		return
	}

	ctx.IndentedJSON(http.StatusFound, account)
}

func handleCreateAccount(ctx *gin.Context) {
	createAccountRequest := new(models.CreateAccountRequest)

	db := ctx.MustGet("db").(*database.PostgresStore)

	if ctx.Request.ContentLength == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Request body must be a valid JSON", "status": http.StatusBadRequest})
		return
	}

	if err := ctx.ShouldBindJSON(&createAccountRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": http.StatusBadRequest})
		return
	}

	account := models.NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName, createAccountRequest.Email)

	if err := db.CreateAccount(account); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	ctx.IndentedJSON(http.StatusOK, account)
}

func handleDeleteAccount(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": http.StatusBadRequest})
		return
	}

	db := ctx.MustGet("db").(*database.PostgresStore)

	if err := db.DeleteAccount(id); err != nil {
		message := fmt.Sprintf("Account with ID - %d not found", id)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": message, "status": http.StatusNotFound})
		return
	}

	message := fmt.Sprintf("Account with ID - %d successfully deleted", id)

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": message, "status": http.StatusOK})
}
