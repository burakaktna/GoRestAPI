package controllers

import (
	"net/http"

	"github.com/burakaktna/GoRestAPI/config"
	"github.com/burakaktna/GoRestAPI/models"
	"github.com/gin-gonic/gin"
)

func ListArticles(ctx *gin.Context) {
	articles := []models.Article{}
	config.DB.Find(&articles)
	ctx.JSON(http.StatusOK, gin.H{"data": articles})
}

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	ctx.BindJSON(&article)
	if err := config.DB.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": article})
}

func ShowArticle(ctx *gin.Context) {
	var article models.Article
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&article).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": article})
}

func UpdateArticle(ctx *gin.Context) {
	var article models.Article
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&article).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.BindJSON(&article)
	if err := config.DB.Save(&article).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": article})
}

func DeleteArticle(ctx *gin.Context) {
	var article models.Article
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&article).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Delete(&article).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
