package v1

import (
	user2 "go-temp/api/v1/user"
)

type ApiGroup struct {
	UserApiGroup user2.UserApiGroup
}

var ApiGroupApp = new(ApiGroup)
