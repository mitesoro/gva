package apis

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apis"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type ApisApi struct {
}

var apisService = service.ServiceGroupApp.ApisServiceGroup.ApisService
var userService = service.ServiceGroupApp.UsersServiceGroup.UsersService

// GetSmsCode 获取短信验证码
// @Tags 前端接口API
// @Summary 获取短信验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body users.Users true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/sms/send [post]
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
	code := utils.GenerateRandomCode(6)
	s := apis.Sms{
		Phone:    req.Phone,
		ExpireAt: time.Now().Unix() + 300,
		Code:     code,
	}
	if err := apisService.CreateSms(&s); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessageWithCode(10001, "获取验证码失败", c)
	} else {
		if !utils.SendSms(req.Phone, code) {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessageWithCode(10001, "获取验证码失败", c)
		}
		response.OkWithMessage("success", c)
	}
}

// Register 注册
// @Tags 前端接口API
// @Summary 注册
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body users.Users true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/register [post]
func (uApi *ApisApi) Register(c *gin.Context) {
	var req apis.ReqRegister
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"phone":    {utils.NotEmpty()},
		"code":     {utils.NotEmpty()},
		"password": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if user, err := userService.GetUsersByPhone(req.Phone); err == nil && user.ID > 0 {
		response.FailWithMessageWithCode(10002, "手机号已注册", c)
		return
	}

	if !apisService.CheckSms(req.Phone, req.Code) {
		response.FailWithMessageWithCode(10002, "验证码错误", c)
		return
	}
	s := users.Users{
		Phone:    req.Phone,
		Password: req.Password,
		Nickname: fmt.Sprintf("zd_%s", utils.GenerateRandomCode(6)),
		Avatar:   "",
	}
	if err := userService.CreateUsers(&s); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessageWithCode(10001, "注册失败", c)
	} else {
		response.OkWithMessage("success", c)
	}
}
