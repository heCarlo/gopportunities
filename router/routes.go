package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heCarlo/gopportunities/handler"
	docs "github.com/heCarlo/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func initializeRoutes(router *gin.Engine) {

	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

	// Initialize Swagger
	router.GET("swagger/*anny", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

