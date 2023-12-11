package service

import (
	"go-daily-work/log"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/util"
)

type projectservice struct{}

var ProjectService = new(projectservice)

func (p *projectservice) GetProjectService() (project []model.Project, err error) {
	if err := util.Master().Model(model.Project{}).Find(&project).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return project, nil
}

func (p *projectservice) AddProjectService(req request.Project) error {
	project := model.Project{
		Name:   req.Name,
		Status: req.Status,
	}
	if err := util.Master().Create(&project).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (p *projectservice) EditProjectService(req request.UpdateProject) error {
	updateData := map[string]interface{}{
		"name":   req.Name,
		"status": req.Status,
	}
	if err := util.Master().Model(model.Project{}).Where("id = ?", req.ProjectId).Updates(updateData).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (p *projectservice) DeleteProjectService(projectId int64) error {
	if err := util.Master().Delete(&model.Project{}, projectId).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}
