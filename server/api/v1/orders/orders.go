package orders

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/orders"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ordersReq "github.com/flipped-aurora/gin-vue-admin/server/model/orders/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OrdersApi struct {
}

var osService = service.ServiceGroupApp.OrdersServiceGroup.OrdersService


// CreateOrders 创建订单
// @Tags Orders
// @Summary 创建订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orders.Orders true "创建订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /os/createOrders [post]
func (osApi *OrdersApi) CreateOrders(c *gin.Context) {
	var os orders.Orders
	err := c.ShouldBindJSON(&os)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "User_id":{utils.NotEmpty()},
        "Account_id":{utils.NotEmpty()},
        "Order_no":{utils.NotEmpty()},
        "Direction":{utils.NotEmpty()},
        "Volume":{utils.NotEmpty()},
        "Price":{utils.NotEmpty()},
    }
	if err := utils.Verify(os, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := osService.CreateOrders(&os); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrders 删除订单
// @Tags Orders
// @Summary 删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orders.Orders true "删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /os/deleteOrders [delete]
func (osApi *OrdersApi) DeleteOrders(c *gin.Context) {
	var os orders.Orders
	err := c.ShouldBindJSON(&os)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osService.DeleteOrders(os); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrdersByIds 批量删除订单
// @Tags Orders
// @Summary 批量删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /os/deleteOrdersByIds [delete]
func (osApi *OrdersApi) DeleteOrdersByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := osService.DeleteOrdersByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrders 更新订单
// @Tags Orders
// @Summary 更新订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body orders.Orders true "更新订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /os/updateOrders [put]
func (osApi *OrdersApi) UpdateOrders(c *gin.Context) {
	var os orders.Orders
	err := c.ShouldBindJSON(&os)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "User_id":{utils.NotEmpty()},
          "Account_id":{utils.NotEmpty()},
          "Order_no":{utils.NotEmpty()},
          "Direction":{utils.NotEmpty()},
          "Volume":{utils.NotEmpty()},
          "Price":{utils.NotEmpty()},
      }
    if err := utils.Verify(os, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := osService.UpdateOrders(os); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrders 用id查询订单
// @Tags Orders
// @Summary 用id查询订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orders.Orders true "用id查询订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /os/findOrders [get]
func (osApi *OrdersApi) FindOrders(c *gin.Context) {
	var os orders.Orders
	err := c.ShouldBindQuery(&os)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reos, err := osService.GetOrders(os.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reos": reos}, c)
	}
}

// GetOrdersList 分页获取订单列表
// @Tags Orders
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ordersReq.OrdersSearch true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /os/getOrdersList [get]
func (osApi *OrdersApi) GetOrdersList(c *gin.Context) {
	var pageInfo ordersReq.OrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := osService.GetOrdersInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
