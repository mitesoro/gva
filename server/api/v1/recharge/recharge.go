package recharge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/recharge"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    rechargeReq "github.com/flipped-aurora/gin-vue-admin/server/model/recharge/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type RechargeApi struct {
}

var rgService = service.ServiceGroupApp.RechargeServiceGroup.RechargeService


// CreateRecharge 创建充值
// @Tags Recharge
// @Summary 创建充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body recharge.Recharge true "创建充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rg/createRecharge [post]
func (rgApi *RechargeApi) CreateRecharge(c *gin.Context) {
	var rg recharge.Recharge
	err := c.ShouldBindJSON(&rg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "User_id":{utils.NotEmpty()},
        "Amount":{utils.NotEmpty()},
    }
	if err := utils.Verify(rg, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := rgService.CreateRecharge(&rg); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRecharge 删除充值
// @Tags Recharge
// @Summary 删除充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body recharge.Recharge true "删除充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rg/deleteRecharge [delete]
func (rgApi *RechargeApi) DeleteRecharge(c *gin.Context) {
	var rg recharge.Recharge
	err := c.ShouldBindJSON(&rg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rgService.DeleteRecharge(rg); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRechargeByIds 批量删除充值
// @Tags Recharge
// @Summary 批量删除充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rg/deleteRechargeByIds [delete]
func (rgApi *RechargeApi) DeleteRechargeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rgService.DeleteRechargeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRecharge 更新充值
// @Tags Recharge
// @Summary 更新充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body recharge.Recharge true "更新充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rg/updateRecharge [put]
func (rgApi *RechargeApi) UpdateRecharge(c *gin.Context) {
	var rg recharge.Recharge
	err := c.ShouldBindJSON(&rg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "User_id":{utils.NotEmpty()},
          "Amount":{utils.NotEmpty()},
      }
    if err := utils.Verify(rg, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := rgService.UpdateRecharge(rg); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRecharge 用id查询充值
// @Tags Recharge
// @Summary 用id查询充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query recharge.Recharge true "用id查询充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rg/findRecharge [get]
func (rgApi *RechargeApi) FindRecharge(c *gin.Context) {
	var rg recharge.Recharge
	err := c.ShouldBindQuery(&rg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerg, err := rgService.GetRecharge(rg.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerg": rerg}, c)
	}
}

// GetRechargeList 分页获取充值列表
// @Tags Recharge
// @Summary 分页获取充值列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rechargeReq.RechargeSearch true "分页获取充值列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rg/getRechargeList [get]
func (rgApi *RechargeApi) GetRechargeList(c *gin.Context) {
	var pageInfo rechargeReq.RechargeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := rgService.GetRechargeInfoList(pageInfo); err != nil {
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
