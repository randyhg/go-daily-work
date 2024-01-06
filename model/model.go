package model

import (
	"go-daily-work/log"
	"go-daily-work/util"
	"gorm.io/gorm"
	"time"
)

const (
	ManagerRoleId = uint(1)
	AdminRoleId   = uint(2)
	StaffRoleId   = uint(3)
)

type Model struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

//var SecretKey = []byte("sasjdakdlkasjk")

type WorkLog struct {
	Model
	UserId         int64          `json:"user_id"`
	User           User           `gorm:"foreignKey:UserId; not null;" json:"user"`
	Project        Project        `gorm:"foreignKey:TaskProjectId; not null;" json:"project"`
	TaskProjectId  int64          `json:"task_project_id"`
	TaskCategory   TaskCategory   `gorm:"foreignKey:TaskCategoryId; not null;" json:"task_category"`
	TaskCategoryId int64          `json:"task_category_id"`
	Description    string         `gorm:"type:longtext" json:"description"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type User struct {
	Model
	Name      string         `gorm:"not null; size:32;" json:"name"`
	Email     string         `gorm:"not null; size:128; uniqueIndex;" json:"email"`
	Password  string         `gorm:"not null; size:512" json:"password"`
	Position  string         `gorm:"not null;" json:"position"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	//RoleId    uint           `gorm:"not null;" json:"role_id"`
}

type TaskCategory struct {
	Model
	Name      string         `gorm:"not null; size:64;" json:"task_name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Project struct {
	Model
	Name      string         `gorm:"not null; size:64;" json:"project_name"`
	Status    int            `gorm:"not null; size:64;" json:"project_status"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func Migration() {
	err := util.Master().AutoMigrate(&User{}, &TaskCategory{}, &Project{}, &WorkLog{})
	if err != nil {
		log.Fatal(err)
		return
	}
}
