package routes

import (
	"github.com/burakaktna/GoRestAPI/controllers"
	"github.com/gin-gonic/gin"
)

func ArticleTagRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/article-tags", controllers.ListArticleTags)
	v1.POST("/article-tags", controllers.CreateArticleTag)
	v1.GET("/article-tags/:id", controllers.ShowArticleTag)
	v1.PUT("/article-tags/:id", controllers.UpdateArticleTag)
	v1.DELETE("/article-tags/:id", controllers.DeleteArticleTag)
}
