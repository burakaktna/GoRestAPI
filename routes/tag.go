package routes

import (
	"github.com/burakaktna/GoRestAPI/controllers"
	"github.com/gin-gonic/gin"
)

func TagRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/tags", controllers.ListTags)
	v1.POST("/tags", controllers.CreateTag)
	v1.GET("/tags/:id", controllers.ShowTag)
	v1.PUT("/tags/:id", controllers.UpdateTag)
	v1.DELETE("/tags/:id", controllers.DeleteTag)
}
