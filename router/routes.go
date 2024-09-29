package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcospsw/first-project-go/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/openings", handler.GetAllOpeningsHandler)
	v1.GET("/opening", handler.GetOpeningHandler)
	v1.POST("/opening", handler.CreateOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)
}
