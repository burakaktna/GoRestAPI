package routes

import (
	"github.com/burakaktna/GoRestAPI/controllers"
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/articles", controllers.ListArticles)
	v1.POST("/articles", controllers.CreateArticle)
	v1.GET("/articles/:id", controllers.ShowArticle)
	v1.PUT("/articles/:id", controllers.UpdateArticle)
	v1.DELETE("/articles/:id", controllers.DeleteArticle)
}
