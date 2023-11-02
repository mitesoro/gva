package alog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlogRouter struct {
}

// InitAlogRouter 初始化 金额纪录 路由信息
func (s *AlogRouter) InitAlogRouter(Router *gin.RouterGroup) {
	alRouter := Router.Group("al").Use(middleware.OperationRecord())
	alRouterWithoutRecord := Router.Group("al")
	var alApi = v1.ApiGroupApp.AlogApiGroup.AlogApi
	{
		alRouter.POST("createAlog", alApi.CreateAlog)   // 新建金额纪录
		alRouter.DELETE("deleteAlog", alApi.DeleteAlog) // 删除金额纪录
		alRouter.DELETE("deleteAlogByIds", alApi.DeleteAlogByIds) // 批量删除金额纪录
		alRouter.PUT("updateAlog", alApi.UpdateAlog)    // 更新金额纪录
	}
	{
		alRouterWithoutRecord.GET("findAlog", alApi.FindAlog)        // 根据ID获取金额纪录
		alRouterWithoutRecord.GET("getAlogList", alApi.GetAlogList)  // 获取金额纪录列表
	}
}
