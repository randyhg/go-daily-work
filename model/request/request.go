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
	UUID        uuid.UUID
	ID          uint
	Username    string
	Email       string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
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
