package user

import (
	"go-temp/service"
)

type UserApiGroup struct {
	BaseApi
}

var userService = service.ServiceGroupApp.UserServiceGroup.UserService
var jwtService = service.ServiceGroupApp.UserServiceGroup.JwtService
