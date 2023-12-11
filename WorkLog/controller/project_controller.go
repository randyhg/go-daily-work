package controller

import (
	"github.com/gin-gonic/gin"
	"go-daily-work/WorkLog/service"
	"go-daily-work/model/request"
	"go-daily-work/model/response"
)

type projectcontroller struct{}

var ProjectController = new(projectcontroller)

func (p *projectcontroller) GetProject(c *gin.Context) {
	lists, err := service.ProjectService.GetProjectService()
	if err != nil {
		response.FailWithDetailed(err, "Failed to get project records", c)
		return
	}
	response.OkWithData(lists, c)
}

func (p *projectcontroller) AddProject(c *gin.Context) {
	var req request.Project
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}

	if err := service.ProjectService.AddProjectService(req); err != nil {
		response.FailWithDetailed(err, "Failed to create project record", c)
		return
	}
	response.OkWithMessage("Project record successfully created", c)
}

func (p *projectcontroller) EditProject(c *gin.Context) {
	var req request.UpdateProject
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}

	if err := service.ProjectService.EditProjectService(req); err != nil {
		response.FailWithDetailed(err, "Failed to update project record", c)
		return
	}
	response.OkWithMessage("Project record successfully updated", c)
}

func (p *projectcontroller) DeleteProject(c *gin.Context) {
	//
	var req request.UpdateProject
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}

	if err := service.ProjectService.DeleteProjectService(req.ProjectId); err != nil {
		response.FailWithDetailed(err, "Failed to delete project record", c)
		return
	}
	response.OkWithMessage("Project record successfully deleted", c)
}
