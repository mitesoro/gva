package symbols

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    symbolsReq "github.com/flipped-aurora/gin-vue-admin/server/model/symbols/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SymbolApi struct {
}

var sbService = service.ServiceGroupApp.SymbolsServiceGroup.SymbolService


// CreateSymbol 创建合约品种
// @Tags Symbol
// @Summary 创建合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbols.Symbol true "创建合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sb/createSymbol [post]
func (sbApi *SymbolApi) CreateSymbol(c *gin.Context) {
	var sb symbols.Symbol
	err := c.ShouldBindJSON(&sb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Code":{utils.NotEmpty()},
        "Multiple":{utils.NotEmpty()},
        "Bond":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
    }
	if err := utils.Verify(sb, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := sbService.CreateSymbol(&sb); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSymbol 删除合约品种
// @Tags Symbol
// @Summary 删除合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbols.Symbol true "删除合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sb/deleteSymbol [delete]
func (sbApi *SymbolApi) DeleteSymbol(c *gin.Context) {
	var sb symbols.Symbol
	err := c.ShouldBindJSON(&sb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sbService.DeleteSymbol(sb); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSymbolByIds 批量删除合约品种
// @Tags Symbol
// @Summary 批量删除合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sb/deleteSymbolByIds [delete]
func (sbApi *SymbolApi) DeleteSymbolByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sbService.DeleteSymbolByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSymbol 更新合约品种
// @Tags Symbol
// @Summary 更新合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body symbols.Symbol true "更新合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sb/updateSymbol [put]
func (sbApi *SymbolApi) UpdateSymbol(c *gin.Context) {
	var sb symbols.Symbol
	err := c.ShouldBindJSON(&sb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Code":{utils.NotEmpty()},
          "Multiple":{utils.NotEmpty()},
          "Bond":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
      }
    if err := utils.Verify(sb, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := sbService.UpdateSymbol(sb); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSymbol 用id查询合约品种
// @Tags Symbol
// @Summary 用id查询合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbols.Symbol true "用id查询合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sb/findSymbol [get]
func (sbApi *SymbolApi) FindSymbol(c *gin.Context) {
	var sb symbols.Symbol
	err := c.ShouldBindQuery(&sb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resb, err := sbService.GetSymbol(sb.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resb": resb}, c)
	}
}

// GetSymbolList 分页获取合约品种列表
// @Tags Symbol
// @Summary 分页获取合约品种列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query symbolsReq.SymbolSearch true "分页获取合约品种列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sb/getSymbolList [get]
func (sbApi *SymbolApi) GetSymbolList(c *gin.Context) {
	var pageInfo symbolsReq.SymbolSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sbService.GetSymbolInfoList(pageInfo); err != nil {
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
