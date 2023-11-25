package notice

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NoticeRouter struct {
}

// InitNoticeRouter 初始化 公告 路由信息
func (s *NoticeRouter) InitNoticeRouter(Router *gin.RouterGroup) {
	noRouter := Router.Group("no").Use(middleware.OperationRecord())
	noRouterWithoutRecord := Router.Group("no")
	var noApi = v1.ApiGroupApp.NoticeApiGroup.NoticeApi
	{
		noRouter.POST("createNotice", noApi.CreateNotice)   // 新建公告
		noRouter.DELETE("deleteNotice", noApi.DeleteNotice) // 删除公告
		noRouter.DELETE("deleteNoticeByIds", noApi.DeleteNoticeByIds) // 批量删除公告
		noRouter.PUT("updateNotice", noApi.UpdateNotice)    // 更新公告
	}
	{
		noRouterWithoutRecord.GET("findNotice", noApi.FindNotice)        // 根据ID获取公告
		noRouterWithoutRecord.GET("getNoticeList", noApi.GetNoticeList)  // 获取公告列表
	}
}
