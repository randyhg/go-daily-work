package controller

import (
	"fmt"
	"go-daily-work/WorkLog/service"
	"go-daily-work/middleware"
	"go-daily-work/model/request"
	"go-daily-work/model/response"

	"github.com/gin-gonic/gin"
)

type worklogcontroller struct{}

var WorkLogController = new(worklogcontroller)

func (w *worklogcontroller) GetWorkLog(c *gin.Context) {
	userId := middleware.GetUser(c).Id
	workLogs, err := service.WorkLogService.GetWorkLogService(userId)
	if err != nil {
		response.FailWithDetailed(err, "Failed to get work log record", c)
		return
	}
	fmt.Println(workLogs[1].Project.Name)
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
	var req request.WorkLog
	//userId := middleware.GetUser(c).Id
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}
}

func (w *worklogcontroller) DeleteWorkLog(c *gin.Context) {
	// delete worklog
}
