package controller

import (
	"go-daily-work/WorkLog/service"
	"go-daily-work/middleware"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/model/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type worklogcontroller struct{}

var WorkLogController = new(worklogcontroller)

func (w *worklogcontroller) GetWorkLog(c *gin.Context) {
	length, _ := strconv.Atoi(c.Param("length"))
	start, _ := strconv.Atoi(c.Param("start"))
	length = 10
	start = 0
	userId := middleware.GetUser(c).Id
	workLogs, err := service.WorkLogService.GetWorkLogService(userId, length, start)
	if err != nil {
		response.FailWithDetailed(err, "Failed to get work log record", c)
		return
	}
	response.OkWithData(workLogs, c)
}

func (w *worklogcontroller) AddWorkLog(c *gin.Context) {
	var req request.WorkLog

	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}
	user := middleware.GetUser(c)

	if err := service.WorkLogService.AddWorkLogService(req, user); err != nil {
		response.FailWithDetailed(err, "Failed to create work log record", c)
		return
	}
	response.OkWithMessage("Work log record successfully created", c)
}

func (w *worklogcontroller) EditWorkLog(c *gin.Context) {
	// edit worklog
	var req model.WorkLog
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}
	if err := service.WorkLogService.EditWorkLogServiceV2(req); err != nil {
		response.FailWithDetailed(err, "Failed to update work log record", c)
		return
	}
	response.OkWithMessage("Work log record successfully updated", c)
}

func (w *worklogcontroller) DeleteWorkLog(c *gin.Context) {
	var req model.WorkLog
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}

	if err := service.WorkLogService.DeleteWorkLogService(req); err != nil {
		response.FailWithMessage("Failed to delete work log record", c)
		return
	}
	response.OkWithMessage("Work log record successfully deleted", c)
}
