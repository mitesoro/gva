package recharge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RechargeRouter struct {
}

// InitRechargeRouter 初始化 充值 路由信息
func (s *RechargeRouter) InitRechargeRouter(Router *gin.RouterGroup) {
	rgRouter := Router.Group("rg").Use(middleware.OperationRecord())
	rgRouterWithoutRecord := Router.Group("rg")
	var rgApi = v1.ApiGroupApp.RechargeApiGroup.RechargeApi
	{
		rgRouter.POST("createRecharge", rgApi.CreateRecharge)   // 新建充值
		rgRouter.DELETE("deleteRecharge", rgApi.DeleteRecharge) // 删除充值
		rgRouter.DELETE("deleteRechargeByIds", rgApi.DeleteRechargeByIds) // 批量删除充值
		rgRouter.PUT("updateRecharge", rgApi.UpdateRecharge)    // 更新充值
	}
	{
		rgRouterWithoutRecord.GET("findRecharge", rgApi.FindRecharge)        // 根据ID获取充值
		rgRouterWithoutRecord.GET("getRechargeList", rgApi.GetRechargeList)  // 获取充值列表
	}
}
