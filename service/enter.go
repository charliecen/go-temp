package service

import (
	"go-temp/service/user"
)

type ServiceGroup struct {
	UserServiceGroup user.UserServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
