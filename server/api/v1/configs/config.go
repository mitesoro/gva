package configs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/configs"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    configsReq "github.com/flipped-aurora/gin-vue-admin/server/model/configs/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ConfigApi struct {
}

var configService = service.ServiceGroupApp.ConfigsServiceGroup.ConfigService


// CreateConfig 创建配置管理
// @Tags Config
// @Summary 创建配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body configs.Config true "创建配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /config/createConfig [post]
func (configApi *ConfigApi) CreateConfig(c *gin.Context) {
	var config configs.Config
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Field":{utils.NotEmpty()},
        "Value":{utils.NotEmpty()},
        "Desc":{utils.NotEmpty()},
    }
	if err := utils.Verify(config, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := configService.CreateConfig(&config); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteConfig 删除配置管理
// @Tags Config
// @Summary 删除配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body configs.Config true "删除配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /config/deleteConfig [delete]
func (configApi *ConfigApi) DeleteConfig(c *gin.Context) {
	var config configs.Config
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configService.DeleteConfig(config); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteConfigByIds 批量删除配置管理
// @Tags Config
// @Summary 批量删除配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /config/deleteConfigByIds [delete]
func (configApi *ConfigApi) DeleteConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configService.DeleteConfigByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateConfig 更新配置管理
// @Tags Config
// @Summary 更新配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body configs.Config true "更新配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /config/updateConfig [put]
func (configApi *ConfigApi) UpdateConfig(c *gin.Context) {
	var config configs.Config
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Field":{utils.NotEmpty()},
          "Value":{utils.NotEmpty()},
          "Desc":{utils.NotEmpty()},
      }
    if err := utils.Verify(config, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := configService.UpdateConfig(config); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindConfig 用id查询配置管理
// @Tags Config
// @Summary 用id查询配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query configs.Config true "用id查询配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /config/findConfig [get]
func (configApi *ConfigApi) FindConfig(c *gin.Context) {
	var config configs.Config
	err := c.ShouldBindQuery(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reconfig, err := configService.GetConfig(config.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reconfig": reconfig}, c)
	}
}

// GetConfigList 分页获取配置管理列表
// @Tags Config
// @Summary 分页获取配置管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query configsReq.ConfigSearch true "分页获取配置管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /config/getConfigList [get]
func (configApi *ConfigApi) GetConfigList(c *gin.Context) {
	var pageInfo configsReq.ConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := configService.GetConfigInfoList(pageInfo); err != nil {
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
