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
	uRouter.Use(middleware.DefaultLogger())
	var aApi = v1.ApiGroupApp.ApisApiGroup.ApisApi
	{
		uRouter.Any("test", aApi.Test)                                                    // test
		uRouter.GET("article/category", aApi.GetArticleCategory)                          // 文章分类
		uRouter.GET("article/list", aApi.GetArticleList)                                  // 文章列表
		uRouter.GET("article/info", aApi.GetArticleInfo)                                  // 文章详情
		uRouter.GET("k/data", aApi.PriceData)                                             // k线
		uRouter.GET("symbol/data/list", aApi.SymbolData)                                  // 行情列表
		uRouter.GET("symbol/data/info", aApi.SymbolDataInfo)                              // 行情详情
		uRouter.POST("sms/send", aApi.GetSmsCode)                                         // 发送短信验证码
		uRouter.POST("register", aApi.Register)                                           // 注册
		uRouter.POST("login", aApi.Login)                                                 // 登录
		uRouter.POST("file/upload", aApi.UploadFile)                                      // 上传文件
		uRouter.Use(middleware.Token()).POST("user/update", aApi.UpdateUser)              // 修改用户信息
		uRouter.Use(middleware.Token()).POST("user/update-phone", aApi.UpdatePhone)       // 修改手机号
		uRouter.Use(middleware.Token()).POST("user/update-password", aApi.UpdatePassword) // 修改密码
		uRouter.Use(middleware.Token()).POST("orders/create", aApi.OrdersCreate)          // 下单
		uRouter.Use(middleware.Token()).GET("user/info", aApi.GetUserInfo)                // 下单
		uRouter.Use(middleware.Token()).GET("orders/list", aApi.OrdersList)               // 交易记录
	}

}
