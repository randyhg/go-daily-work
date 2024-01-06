package service

import (
	"go-daily-work/log"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/model/response"
	"go-daily-work/util"
)

type worklogservice struct{}

var WorkLogService = new(worklogservice)

func (w *worklogservice) GetWorkLogService(userId int64, limit, offset int) (workLogs []response.WorkLogResp, err error) {
	if err = util.Master().
		Table("work_logs wl").
		Select("wl.id Id, wl.description Description, p.name ProjectName, tc.name TaskCategory").
		Joins("JOIN users u ON wl.user_id = u.id").
		Joins("JOIN projects p ON wl.task_project_id = p.id").
		Joins("JOIN task_categories tc ON wl.task_category_id = tc.id").
		Where("user_id = ?", userId).Order("Id").Limit(limit).Offset(offset).
		Scan(&workLogs).Error; err != nil {
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

func (w *worklogservice) EditWorkLogServiceV2(req model.WorkLog) error {
	var existingWorklog model.WorkLog
	result := util.Master().First(&existingWorklog, req.Id)
	if result.Error != nil {
		log.Error(result.Error)
		return result.Error
	}

	if err := util.Master().Model(&existingWorklog).Updates(&req).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (w *worklogservice) DeleteWorkLogService(req model.WorkLog) error {
	if err := util.Master().Delete(&req, req.Id).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

//func (w *worklogservice) EditWorkLogService(req request.UpdateWorkLog) error {
//	updateData := map[string]interface{}{
//		"task_project_id":  req.TaskProject,
//		"task_category_id": req.TaskCategory,
//		"description":      req.Description,
//	}
//	if err := util.Master().Model(model.WorkLog{}).Where("id = ?", req.WorkLogId).Updates(updateData).Error; err != nil {
//		log.Error(err)
//		return err
//	}
//	return nil
//}
