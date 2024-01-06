package request

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomClaims struct {
	UUID     uuid.UUID
	ID       uint
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Position string `json:"position,omitempty"`
	jwt.StandardClaims
	//BufferTime int64
}

type WorkLog struct {
	TaskProject  int64  `json:"task_project"`
	TaskCategory int64  `json:"task_category"`
	Description  string `json:"description"`
}

type UpdateWorkLog struct {
	WorkLogId int64 `json:"work_log_id"`
	WorkLog
}

type Category struct {
	Name string `json:"category_name"`
}

type UpdateCategory struct {
	CategoryId int64 `json:"category_id"`
	Category
}

type Project struct {
	Name   string `json:"project_name"`
	Status int    `json:"project_status"`
}

type UpdateProject struct {
	ProjectId int64 `json:"project_id"`
	Project
}
