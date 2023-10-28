package orders

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrdersRouter struct {
}

// InitOrdersRouter 初始化 订单 路由信息
func (s *OrdersRouter) InitOrdersRouter(Router *gin.RouterGroup) {
	osRouter := Router.Group("os").Use(middleware.OperationRecord())
	osRouterWithoutRecord := Router.Group("os")
	var osApi = v1.ApiGroupApp.OrdersApiGroup.OrdersApi
	{
		osRouter.POST("createOrders", osApi.CreateOrders)   // 新建订单
		osRouter.DELETE("deleteOrders", osApi.DeleteOrders) // 删除订单
		osRouter.DELETE("deleteOrdersByIds", osApi.DeleteOrdersByIds) // 批量删除订单
		osRouter.PUT("updateOrders", osApi.UpdateOrders)    // 更新订单
	}
	{
		osRouterWithoutRecord.GET("findOrders", osApi.FindOrders)        // 根据ID获取订单
		osRouterWithoutRecord.GET("getOrdersList", osApi.GetOrdersList)  // 获取订单列表
	}
}
