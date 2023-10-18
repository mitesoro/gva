package apis

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apis"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type ApisApi struct {
}

var apisService = service.ServiceGroupApp.ApisServiceGroup.ApisService

// GetSmsCode 获取短信验证码
// @Tags Sms
// @Summary 获取短信验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body users.Users true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sms/send [post]
func (uApi *ApisApi) GetSmsCode(c *gin.Context) {
	var req apis.ReqSms
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Phone": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	s := apis.Sms{
		Phone:    req.Phone,
		ExpireAt: time.Now().Unix() + 300,
		Code:     utils.GenerateRandomCode(6),
	}
	if err := apisService.CreateSms(&s); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessageWithCode(10001, "获取验证码失败", c)
	} else {
		response.OkWithMessage("success", c)
	}
}
