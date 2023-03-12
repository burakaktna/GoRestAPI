package controllers

import (
	"net/http"

	"github.com/burakaktna/GoRestAPI/config"
	"github.com/burakaktna/GoRestAPI/models"
	"github.com/gin-gonic/gin"
)

func ListTags(ctx *gin.Context) {
	tags := []models.Tag{}
	config.DB.Find(&tags)
	ctx.JSON(http.StatusOK, gin.H{"data": tags})
}

func CreateTag(ctx *gin.Context) {
	var tag models.Tag
	ctx.BindJSON(&tag)
	if err := config.DB.Create(&tag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": tag})
}

func ShowTag(ctx *gin.Context) {
	var tag models.Tag
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&tag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": tag})
}

func UpdateTag(ctx *gin.Context) {
	var tag models.Tag
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&tag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.BindJSON(&tag)
	if err := config.DB.Save(&tag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": tag})
}

func DeleteTag(ctx *gin.Context) {
	var tag models.Tag
	if err := config.DB.Where("id = ?", ctx.Param("id")).First(&tag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Delete(&tag).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
