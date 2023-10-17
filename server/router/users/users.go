package users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UsersRouter struct {
}

// InitUsersRouter 初始化 用户 路由信息
func (s *UsersRouter) InitUsersRouter(Router *gin.RouterGroup) {
	uRouter := Router.Group("u").Use(middleware.OperationRecord())
	uRouterWithoutRecord := Router.Group("u")
	var uApi = v1.ApiGroupApp.UsersApiGroup.UsersApi
	{
		uRouter.POST("createUsers", uApi.CreateUsers)   // 新建用户
		uRouter.DELETE("deleteUsers", uApi.DeleteUsers) // 删除用户
		uRouter.DELETE("deleteUsersByIds", uApi.DeleteUsersByIds) // 批量删除用户
		uRouter.PUT("updateUsers", uApi.UpdateUsers)    // 更新用户
	}
	{
		uRouterWithoutRecord.GET("findUsers", uApi.FindUsers)        // 根据ID获取用户
		uRouterWithoutRecord.GET("getUsersList", uApi.GetUsersList)  // 获取用户列表
	}
}
