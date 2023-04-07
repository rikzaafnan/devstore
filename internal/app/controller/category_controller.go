package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rikzaafnan/devstore/internal/app/schema"
	"github.com/rikzaafnan/devstore/internal/app/service"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {

	resp, err := cc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {

	var req schema.CreateCategoryReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = cc.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success create category"})
}

func (cc *CategoryController) DetailCategory(ctx *gin.Context) {

	paramCategoryID := ctx.Param("categoryID")

	categoryID, _ := strconv.ParseInt(paramCategoryID, 0, 64)

	category, err := cc.service.FindOneByID(int(categoryID))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": category})
}

func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {

	var req schema.CreateCategoryReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	paramCategoryID := ctx.Param("categoryID")

	categoryID, _ := strconv.ParseInt(paramCategoryID, 0, 64)

	category, err := cc.service.Update(int(categoryID), req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": category})
}

func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {

	paramCategoryID := ctx.Param("categoryID")

	categoryID, _ := strconv.ParseInt(paramCategoryID, 0, 64)

	err := cc.service.DeleteByID(int(categoryID))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success delete category"})
}
