package service

import (
	"go-daily-work/log"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/util"
)

type worklogservice struct{}

var WorkLogService = new(worklogservice)

func (w *worklogservice) GetWorkLogService(userId int64) ([]model.WorkLog, error) {
	var workLogs []model.WorkLog
	if err := util.Master().
		Preload("Project.Name").
		//Preload("TaskCategory.Name").
		Where("user_id = ?", userId).Find(&workLogs).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return workLogs, nil
}

func (w *worklogservice) AddWorkLogService(req request.WorkLog, user *model.User) error {
	workLog := model.WorkLog{
		UserId:         user.Id,
		TaskCategoryId: req.TaskCategory,
		TaskProjectId:  req.TaskProject,
		Description:    req.Description,
	}

	if err := util.Master().Create(&workLog).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (w *worklogservice) EditWorkLogService(req request.UpdateWorkLog, userId int64) error {
	updateData := map[string]interface{}{
		"user_id":          userId,
		"task_project_id":  req.TaskProject,
		"task_category_id": req.TaskCategory,
		"description":      req.Description,
	}
	if err := util.Master().Model(model.WorkLog{}).Where("id = ?", req.WorkLogId).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}
