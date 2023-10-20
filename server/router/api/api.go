package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

// InitApiRouter 初始化 api 路由信息
func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	uRouter := Router.Group("api")
	var aApi = v1.ApiGroupApp.ApisApiGroup.ApisApi
	{
		uRouter.POST("sms/send", aApi.GetSmsCode) // 发送短信验证码
		uRouter.POST("register", aApi.Register)   // 注册
	}

}
