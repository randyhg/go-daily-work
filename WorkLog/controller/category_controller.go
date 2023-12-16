package controller

import (
	"github.com/gin-gonic/gin"
	"go-daily-work/WorkLog/service"
	"go-daily-work/model/request"
	"go-daily-work/model/response"
	"strconv"
)

type categorycontroller struct{}

var CategoryController = new(categorycontroller)

func (w *categorycontroller) GetCategory(c *gin.Context) {
	length, _ := strconv.Atoi(c.Param("length"))
	start, _ := strconv.Atoi(c.Param("start"))
	length = 10
	start = 0
	categoryList, err := service.CategoryService.GetCategoryService(length, start)
	if err != nil {
		response.FailWithDetailed(err, "Failed to get category records", c)
		return
	}
	response.OkWithData(categoryList, c)
}

func (w *categorycontroller) AddCategory(c *gin.Context) {
	var req request.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}

	if err := service.CategoryService.AddCategoryService(req); err != nil {
		response.FailWithDetailed(err, "Failed to create category record", c)
		return
	}
	response.OkWithMessage("Category record successfully created", c)
}

func (w *categorycontroller) EditCategory(c *gin.Context) {
	var req request.UpdateCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}

	if err := service.CategoryService.EditCategoryService(req); err != nil {
		response.FailWithDetailed(err, "Failed to update category record", c)
		return
	}
	response.OkWithMessage("Category record successfully updated", c)
}

func (w *categorycontroller) DeleteCategory(c *gin.Context) {
	var req request.UpdateCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}

	if err := service.CategoryService.DeleteCategoryService(req.CategoryId); err != nil {
		response.FailWithDetailed(err, "Failed to delete category record", c)
		return
	}
	response.OkWithMessage("Category record successfully deleted", c)
}
