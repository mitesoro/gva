package configs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConfigRouter struct {
}

// InitConfigRouter 初始化 配置管理 路由信息
func (s *ConfigRouter) InitConfigRouter(Router *gin.RouterGroup) {
	configRouter := Router.Group("config").Use(middleware.OperationRecord())
	configRouterWithoutRecord := Router.Group("config")
	var configApi = v1.ApiGroupApp.ConfigsApiGroup.ConfigApi
	{
		configRouter.POST("createConfig", configApi.CreateConfig)   // 新建配置管理
		configRouter.DELETE("deleteConfig", configApi.DeleteConfig) // 删除配置管理
		configRouter.DELETE("deleteConfigByIds", configApi.DeleteConfigByIds) // 批量删除配置管理
		configRouter.PUT("updateConfig", configApi.UpdateConfig)    // 更新配置管理
	}
	{
		configRouterWithoutRecord.GET("findConfig", configApi.FindConfig)        // 根据ID获取配置管理
		configRouterWithoutRecord.GET("getConfigList", configApi.GetConfigList)  // 获取配置管理列表
	}
}
