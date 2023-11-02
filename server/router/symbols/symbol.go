package symbols

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SymbolRouter struct {
}

// InitSymbolRouter 初始化 合约品种 路由信息
func (s *SymbolRouter) InitSymbolRouter(Router *gin.RouterGroup) {
	sbRouter := Router.Group("sb").Use(middleware.OperationRecord())
	sbRouterWithoutRecord := Router.Group("sb")
	var sbApi = v1.ApiGroupApp.SymbolsApiGroup.SymbolApi
	{
		sbRouter.POST("createSymbol", sbApi.CreateSymbol)   // 新建合约品种
		sbRouter.DELETE("deleteSymbol", sbApi.DeleteSymbol) // 删除合约品种
		sbRouter.DELETE("deleteSymbolByIds", sbApi.DeleteSymbolByIds) // 批量删除合约品种
		sbRouter.PUT("updateSymbol", sbApi.UpdateSymbol)    // 更新合约品种
	}
	{
		sbRouterWithoutRecord.GET("findSymbol", sbApi.FindSymbol)        // 根据ID获取合约品种
		sbRouterWithoutRecord.GET("getSymbolList", sbApi.GetSymbolList)  // 获取合约品种列表
	}
}
