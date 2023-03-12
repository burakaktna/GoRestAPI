package controllers

import (
	"net/http"

	"github.com/burakaktna/GoRestAPI/config"
	"github.com/burakaktna/GoRestAPI/models"
	"github.com/gin-gonic/gin"
)

func ListArticleTags(ctx *gin.Context) {
	ArticleTags := []models.ArticleTag{}
	config.DB.Find(&ArticleTags)
	ctx.JSON(http.StatusOK, gin.H{"data": ArticleTags})
}

func CreateArticleTag(ctx *gin.Context) {
	var ArticleTag models.ArticleTag
	ctx.BindJSON(&ArticleTag)
	if err := config.DB.Create(&ArticleTag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": ArticleTag})
}

func ShowArticleTag(ctx *gin.Context) {
	var ArticleTag models.ArticleTag
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&ArticleTag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": ArticleTag})
}

func UpdateArticleTag(ctx *gin.Context) {
	var ArticleTag models.ArticleTag
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&ArticleTag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.BindJSON(&ArticleTag)
	if err := config.DB.Save(&ArticleTag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": ArticleTag})
}

func DeleteArticleTag(ctx *gin.Context) {
	var ArticleTag models.ArticleTag
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&ArticleTag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Delete(&ArticleTag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
