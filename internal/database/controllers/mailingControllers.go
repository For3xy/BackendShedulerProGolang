package controllers

import (
	"backendShedulerPro/internal/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MailingController struct {
	DB *gorm.DB
}

// Constructor
func NewMailingController(db *gorm.DB) MailingController {
	return MailingController{db}
}

func (mc *MailingController) CreateMailing(ctx *gin.Context) {
	var mailing *models.CreateMailing

	if err := ctx.ShouldBindJSON(&mailing); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newMailing := models.Mailing{
		TypeMailing:    mailing.TypeMailing,
		Filter:         mailing.Filter,
		GroupFlag:      mailing.GroupFlag,
		SingleFileFlag: mailing.SingleFileFlag,
		Address:        mailing.Address,
		ChatId:         mailing.ChatId,
	}

	result := mc.DB.Create(&newMailing)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newMailing})
}

func (mc *MailingController) UpdateMailing(ctx *gin.Context) {
	mailingId := ctx.Param("id")

	var mailing *models.UpdateMailing
	if err := ctx.ShouldBindJSON(&mailing); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedMailing models.Mailing
	result := mc.DB.First(&updatedMailing, "id = ?", mailingId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No mailing found with that id"})
		return
	}

	mailingToUpdate := models.Mailing{
		TypeMailing:    mailing.TypeMailing,
		Filter:         mailing.Filter,
		GroupFlag:      mailing.GroupFlag,
		SingleFileFlag: mailing.SingleFileFlag,
		Address:        mailing.Address,
		ChatId:         mailing.ChatId,
	}

	mc.DB.Model(&updatedMailing).Updates(mailingToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedMailing})
}

func (mc *MailingController) FindMailingById(ctx *gin.Context) {
	mailingId := ctx.Param("id")

	var mailing models.Mailing
	result := mc.DB.First(&mailing, "id = ?", mailingId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No mailing found with that id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": mailing})
}

func (mc *MailingController) FindAllMailing(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPages, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPages - 1) * intLimit

	var mailings []models.Mailing
	results := mc.DB.Limit(intLimit).Offset(offset).Find(&mailings)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(mailings), "data": mailings})
}

func (mc *MailingController) DeleteMailing(ctx *gin.Context) {
	mailingId := ctx.Param("id")
	result := mc.DB.Delete(&models.Mailing{}, "id = ?", mailingId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No mailing found with that id"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
