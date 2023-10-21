package apis

import (
	"encoding/base64"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apis"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

var CosClient *cos.Client

func init() {
	// 创建 COS 客户端
	u, _ := url.Parse("https://zd-1321537534.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	CosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDrhQEziNcmnbEgvxmNl180hOyFKH4d1GR",
			SecretKey: "DF8oERSXWIuJXzIZmocEduzCF5l4C8Vl",
		},
	})
}

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
// @Param data body apis.ReqSms true "获取验证码"
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
// @Param data body apis.ReqRegister true "注册"
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

// Login 登录
// @Tags 前端接口API
// @Summary 登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apis.ReqLogin true "登录"
// @Success 200 {object} string "{"code":0,"data":{"access_token":"ZWdUJTJCUGliYWNZQ0NGVUwzQUtUbFZwd3ZlR3FDWURGMFFlZUdMWFJaQzdMY3BWbGxiUnR6VjJOc0JYVmRBMEpsdlJCayUyQjVLWHUxNHRLVXUwekZ5dE5nJTNEJTNE"},"msg":"查询成功"}"
// @Router /api/login [post]
func (uApi *ApisApi) Login(c *gin.Context) {
	var req apis.ReqLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"phone":    {utils.NotEmpty()},
		"password": {utils.NotEmpty()},
	}
	if err = utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := userService.GetUsersByPhoneAndPassword(req.Phone, req.Password)
	if err != nil || user.ID == 0 {
		response.FailWithMessageWithCode(10002, "手机号或密码错误", c)
		return
	}
	p := map[string]string{
		"uid":   cast.ToString(user.ID),
		"phone": user.Phone,
		"time":  cast.ToString(time.Now().Unix()),
	}

	data := map[string]interface{}{
		"access_token": base64.StdEncoding.EncodeToString([]byte(utils.AESEncodeNormal(p, utils.Sign))),
	}
	response.OkWithData(data, c)
	return
}

// UploadFile 上传文件
// @Tags 前端接口API
// @Summary 上传文件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file   true  "上传文件示例"
// @Success   200   {object}  response.Response{data=object,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /api/file/upload [post]
func (a *ApisApi) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessageWithCode(10002, "文件错误", c)
		return
	}
	remoteFilePath := "uploads/" + file.Filename // 远程 COS 文件路径
	f, _ := file.Open()
	_, err = CosClient.Object.Put(c.Request.Context(), remoteFilePath, f, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{
		"path": fmt.Sprintf("https://zd-1321537534.cos.ap-nanjing.myqcloud.com/%s", remoteFilePath),
	}
	response.OkWithData(data, c)
	return
}
