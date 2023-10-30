package apis

import (
	"encoding/base64"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apis"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orders"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"math/rand"
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
var orderService = service.ServiceGroupApp.OrdersServiceGroup.OrdersService

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
		"time":  cast.ToString(time.Now().Unix() + 86400*30),
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
func (uApi *ApisApi) UploadFile(c *gin.Context) {
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

// UpdateUser 更新用户信息
// @Tags 前端接口API
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apis.ReqUpdateUser true "更新用户信息"
// @Success 200 {object} object "{"code":0,"data":{},"msg":"success"}"
// @Router /api/user/update [post]
func (uApi *ApisApi) UpdateUser(c *gin.Context) {
	var req apis.ReqUpdateUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"avatar":   {utils.NotEmpty()},
		"nickname": {utils.NotEmpty()},
	}
	if err = utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	id, _ := c.Get("uid")
	user, err := userService.GetUsers(cast.ToUint(id))
	if err != nil {
		response.FailWithMessageWithCode(10002, "请重新登录", c)
		return
	}
	user.Avatar = req.Avatar
	user.Nickname = req.Nickname
	if err = userService.UpdateUsers(user); err != nil {
		response.FailWithMessageWithCode(10003, "更新失败", c)
		return
	}
	response.OkWithMessage("success", c)
	return
}

// UpdatePhone 更换手机号
// @Tags 前端接口API
// @Summary 更换手机号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apis.ReqUpdatePhone true "更新手机号"
// @Success 200 {object} object "{"code":0,"data":{},"msg":"success"}"
// @Router /api/user/update-phone [post]
func (uApi *ApisApi) UpdatePhone(c *gin.Context) {
	var req apis.ReqUpdatePhone
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"phone": {utils.NotEmpty()},
		"code":  {utils.NotEmpty()},
	}
	if err = utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if !apisService.CheckSms(req.Phone, req.Code) {
		response.FailWithMessageWithCode(10002, "验证码错误", c)
		return
	}

	id, _ := c.Get("uid")
	user, err := userService.GetUsers(cast.ToUint(id))
	if err != nil {
		response.FailWithMessageWithCode(10002, "请重新登录", c)
		return
	}
	user.Phone = req.Phone
	if err = userService.UpdateUsers(user); err != nil {
		response.FailWithMessageWithCode(10003, "更新失败", c)
		return
	}
	response.OkWithMessage("success", c)
	return
}

// UpdatePassword 更换密码
// @Tags 前端接口API
// @Summary 更换密码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apis.ReqUpdatePassword true "更新密码"
// @Success 200 {object} object "{"code":0,"data":{},"msg":"success"}"
// @Router /api/user/update-password [post]
func (uApi *ApisApi) UpdatePassword(c *gin.Context) {
	var req apis.ReqUpdatePassword
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"password": {utils.NotEmpty()},
	}
	if err = utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	id, _ := c.Get("uid")
	user, err := userService.GetUsers(cast.ToUint(id))
	if err != nil {
		response.FailWithMessageWithCode(10002, "请重新登录", c)
		return
	}
	if user.Password != req.OldPassword {
		response.FailWithMessageWithCode(10002, "旧密码错误", c)
		return
	}
	user.Password = req.Password
	if err = userService.UpdateUsers(user); err != nil {
		response.FailWithMessageWithCode(10003, "更新失败", c)
		return
	}
	response.OkWithMessage("success", c)
	return
}

// OrdersCreate 下单
// @Tags 前端接口API
// @Summary 下单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apis.ReqOrders true "下单参数"
// @Success 200 {object} object "{"code":0,"data":{},"msg":"success"}"
// @Router /api/orders/create [post]
func (uApi *ApisApi) OrdersCreate(c *gin.Context) {
	var req apis.ReqOrders
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"volume":    {utils.NotEmpty()},
		"price":     {utils.NotEmpty()},
		"direction": {utils.NotEmpty()},
	}
	if err = utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	id, _ := c.Get("uid")
	userID := cast.ToInt(id)
	price := int(req.Price)
	volume := int(req.Volume)
	direction := int(req.Direction)
	accountID := 1
	order := &orders.Orders{
		User_id:    &userID,
		Account_id: &accountID,
		Price:      &price,
		Volume:     &volume,
		Direction:  &direction,
		Order_no:   utils.MD5(fmt.Sprintf("%d", time.Now().UnixNano())),
	}
	err = orderService.CreateOrders(order)
	if err != nil {
		response.FailWithMessageWithCode(10002, "下单失败", c)
		return
	}

	response.OkWithMessage("success", c)
	return
}

// PriceData k线数据
// @Tags 前端接口API
// @Summary k线数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apis.KData true "下单参数"
// @Success 200 {object} object "{"code":0,"data":{},"msg":"success"}"
// @Router /api/k/data [get]
func (uApi *ApisApi) PriceData(c *gin.Context) {
	var req apis.KData
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var yData []apis.YdData
	tm := time.Now().Unix()
	for i := 0; i < int(req.Rows); i++ {
		yData = append(yData, apis.YdData{
			Uptime: tm - int64(i),
			Open:   rand.Float64() * 10000,
			High:   rand.Float64() * 10000,
			Low:    rand.Float64() * 10000,
			Close:  rand.Float64() * 10000,
			Vol:    rand.Float64() * 10000,
			Cjl:    rand.Float64() * 10000,
		})
	}
	response.OkWithData(apis.KDataResp{
		YdClosePrice: rand.Float64() * 10000,
		Results:      yData,
	}, c)
	return
}
