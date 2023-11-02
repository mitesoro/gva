package alog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/alog"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    alogReq "github.com/flipped-aurora/gin-vue-admin/server/model/alog/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AlogApi struct {
}

var alService = service.ServiceGroupApp.AlogServiceGroup.AlogService


// CreateAlog 创建金额纪录
// @Tags Alog
// @Summary 创建金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alog.Alog true "创建金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /al/createAlog [post]
func (alApi *AlogApi) CreateAlog(c *gin.Context) {
	var al alog.Alog
	err := c.ShouldBindJSON(&al)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "User_id":{utils.NotEmpty()},
        "Amount_type":{utils.NotEmpty()},
        "Amount":{utils.NotEmpty()},
        "Cur_amount":{utils.NotEmpty()},
    }
	if err := utils.Verify(al, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := alService.CreateAlog(&al); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAlog 删除金额纪录
// @Tags Alog
// @Summary 删除金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alog.Alog true "删除金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /al/deleteAlog [delete]
func (alApi *AlogApi) DeleteAlog(c *gin.Context) {
	var al alog.Alog
	err := c.ShouldBindJSON(&al)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := alService.DeleteAlog(al); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAlogByIds 批量删除金额纪录
// @Tags Alog
// @Summary 批量删除金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /al/deleteAlogByIds [delete]
func (alApi *AlogApi) DeleteAlogByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := alService.DeleteAlogByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAlog 更新金额纪录
// @Tags Alog
// @Summary 更新金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alog.Alog true "更新金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /al/updateAlog [put]
func (alApi *AlogApi) UpdateAlog(c *gin.Context) {
	var al alog.Alog
	err := c.ShouldBindJSON(&al)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "User_id":{utils.NotEmpty()},
          "Amount_type":{utils.NotEmpty()},
          "Amount":{utils.NotEmpty()},
          "Cur_amount":{utils.NotEmpty()},
      }
    if err := utils.Verify(al, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := alService.UpdateAlog(al); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAlog 用id查询金额纪录
// @Tags Alog
// @Summary 用id查询金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alog.Alog true "用id查询金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /al/findAlog [get]
func (alApi *AlogApi) FindAlog(c *gin.Context) {
	var al alog.Alog
	err := c.ShouldBindQuery(&al)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if real, err := alService.GetAlog(al.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"real": real}, c)
	}
}

// GetAlogList 分页获取金额纪录列表
// @Tags Alog
// @Summary 分页获取金额纪录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alogReq.AlogSearch true "分页获取金额纪录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /al/getAlogList [get]
func (alApi *AlogApi) GetAlogList(c *gin.Context) {
	var pageInfo alogReq.AlogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := alService.GetAlogInfoList(pageInfo); err != nil {
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
