package util

import (
	permify "github.com/Permify/permify-gorm"
	"go-daily-work/log"
)

var permifyInstance *permify.Permify

func Permify() *permify.Permify {
	return permifyInstance
}

func PermifyInit() {
	var err error
	permifyInstance, err = permify.New(permify.Options{
		Migrate: true,
		DB:      Master(),
	})
	if err != nil {
		log.Error(err)
		return
	}
	permifyInstance.CreateRole("manager", "Manager desc")
	permifyInstance.CreateRole("admin", "Admin desc")
	permifyInstance.CreateRole("staff", "Staff desc")

	permifyInstance.CreatePermission("manage project", "")
	permifyInstance.CreatePermission("manage category", "")
	permifyInstance.CreatePermission("manage DWL", "")
	permifyInstance.CreatePermission("manage user", "")

	permifyInstance.AddPermissionsToRole("manager", []string{"manage project", "manage category", "manage DWL", "manage user"})
	permifyInstance.AddPermissionsToRole("admin", []string{"manage project", "manage category", "manage DWL"})
	permifyInstance.AddPermissionsToRole("staff", []string{"manage DWL"})
}
