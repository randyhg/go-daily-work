package service

import (
	"go-daily-work/log"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/util"
)

type categoryservice struct{}

var CategoryService = new(categoryservice)

func (c *categoryservice) GetCategoryService(limit, offset int) (category *[]model.TaskCategory, err error) {
	if err := util.Master().Model(model.TaskCategory{}).Limit(limit).Offset(offset).Find(&category).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return category, nil
}

func (c *categoryservice) AddCategoryService(req request.Category) error {
	category := model.TaskCategory{
		Name: req.Name,
	}
	if err := util.Master().Create(&category).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (c *categoryservice) EditCategoryService(req request.UpdateCategory) error {
	updateData := map[string]interface{}{
		"name": req.Name,
	}
	if err := util.Master().Model(model.TaskCategory{}).Where("id = ?", req.CategoryId).Updates(updateData).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (c *categoryservice) DeleteCategoryService(categoryId int64) error {
	if err := util.Master().Delete(&model.TaskCategory{}, categoryId).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}
