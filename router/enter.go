package router

import (
	"go-temp/router/user"
)

type RouterGroup struct {
	User user.UserRouteGroup
}

var RouterGroupApp = new(RouterGroup)
