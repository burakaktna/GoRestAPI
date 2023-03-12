package main

import (
	"github.com/burakaktna/GoRestAPI/config"
	"github.com/burakaktna/GoRestAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.ArticleRoutes(router)
	routes.TagRoutes(router)
	routes.ArticleTagRoutes(router)
	router.Run(":8096")
}
