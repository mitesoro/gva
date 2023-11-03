package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

// InitApiRouter 初始化 api 路由信息
func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	uRouter := Router.Group("api")
	var aApi = v1.ApiGroupApp.ApisApiGroup.ApisApi
	{
		uRouter.GET("k/data", aApi.PriceData)                                             // k线
		uRouter.POST("sms/send", aApi.GetSmsCode)                                         // 发送短信验证码
		uRouter.POST("register", aApi.Register)                                           // 注册
		uRouter.POST("login", aApi.Login)                                                 // 登录
		uRouter.POST("file/upload", aApi.UploadFile)                                      // 上传文件
		uRouter.POST("user/update", aApi.UpdateUser).Use(middleware.Token())              // 修改用户信息
		uRouter.POST("user/update-phone", aApi.UpdatePhone).Use(middleware.Token())       // 修改手机号
		uRouter.POST("user/update-password", aApi.UpdatePassword).Use(middleware.Token()) // 修改密码
		uRouter.POST("orders/create", aApi.OrdersCreate).Use(middleware.Token())          // 下单
		uRouter.GET("user/info", aApi.GetUserInfo).Use(middleware.Token())                // 下单
		uRouter.GET("orders/list", aApi.OrdersList).Use(middleware.Token())               // 交易记录
	}

}
