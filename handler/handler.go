package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"GET": "aaaa",
	})
}

func GetOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"GET": "aaaa",
	})
}

func CreateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"GET": "aaaa",
	})
}

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"GET": "aaaa",
	})
}

func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"GET": "aaaa",
	})
}
