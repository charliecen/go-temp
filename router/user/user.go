package user

import (
	"github.com/gin-gonic/gin"
	v12 "go-temp/api/v1"
	"go-temp/middleware"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	ur := rg.Group("user").Use(middleware.OperationRecord())
	var baseApi = v12.ApiGroupApp.UserApiGroup.BaseApi
	{
		ur.POST("register", baseApi.Register) // 用户注册
		//ur.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		//ur.POST("getUserList", baseApi.GetUserList)               // 分页获取用户列表
		//ur.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		//ur.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		//ur.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		//ur.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		//ur.GET("getUserInfo", baseApi.GetUserInfo)
	}
}
