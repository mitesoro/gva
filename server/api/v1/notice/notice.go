package notice

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/notice"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    noticeReq "github.com/flipped-aurora/gin-vue-admin/server/model/notice/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type NoticeApi struct {
}

var noService = service.ServiceGroupApp.NoticeServiceGroup.NoticeService


// CreateNotice 创建公告
// @Tags Notice
// @Summary 创建公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body notice.Notice true "创建公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /no/createNotice [post]
func (noApi *NoticeApi) CreateNotice(c *gin.Context) {
	var no notice.Notice
	err := c.ShouldBindJSON(&no)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := noService.CreateNotice(&no); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNotice 删除公告
// @Tags Notice
// @Summary 删除公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body notice.Notice true "删除公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /no/deleteNotice [delete]
func (noApi *NoticeApi) DeleteNotice(c *gin.Context) {
	var no notice.Notice
	err := c.ShouldBindJSON(&no)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := noService.DeleteNotice(no); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNoticeByIds 批量删除公告
// @Tags Notice
// @Summary 批量删除公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /no/deleteNoticeByIds [delete]
func (noApi *NoticeApi) DeleteNoticeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := noService.DeleteNoticeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNotice 更新公告
// @Tags Notice
// @Summary 更新公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body notice.Notice true "更新公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /no/updateNotice [put]
func (noApi *NoticeApi) UpdateNotice(c *gin.Context) {
	var no notice.Notice
	err := c.ShouldBindJSON(&no)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := noService.UpdateNotice(no); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNotice 用id查询公告
// @Tags Notice
// @Summary 用id查询公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query notice.Notice true "用id查询公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /no/findNotice [get]
func (noApi *NoticeApi) FindNotice(c *gin.Context) {
	var no notice.Notice
	err := c.ShouldBindQuery(&no)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reno, err := noService.GetNotice(no.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reno": reno}, c)
	}
}

// GetNoticeList 分页获取公告列表
// @Tags Notice
// @Summary 分页获取公告列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query noticeReq.NoticeSearch true "分页获取公告列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /no/getNoticeList [get]
func (noApi *NoticeApi) GetNoticeList(c *gin.Context) {
	var pageInfo noticeReq.NoticeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := noService.GetNoticeInfoList(pageInfo); err != nil {
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
