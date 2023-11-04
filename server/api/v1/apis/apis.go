package apis

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apis"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orders"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/flipped-aurora/gin-vue-admin/server/pb"
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
var orderService = service.ServiceGroupApp.OrdersServiceGroup.OrdersService

func (uApi *ApisApi) Test(c *gin.Context) {
	// grpc 调用下单接口
	res, err := global.GVA_GrpcCLient.Order(context.Background(), &pb.OrderRequest{
		C:       "rb2403",
		V:       1,
		Buy:     true,
		Open:    true,
		OrderId: 134,
	})
	if err != nil {
		global.GVA_LOG.Error("grpc Order", zap.Error(err))
	}
	response.OkWithData(res, c)
}

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
	ss, err := utils.GetSymbol(req.Symbol)
	if err != nil {
		response.FailWithMessageWithCode(10002, "下单失败", c)
		return
	}
	id, _ := c.Get("uid")
	userID := cast.ToInt(id)
	price := int(req.Price)
	volume := int(req.Volume)
	direction := int(req.Direction)
	accountID := 1
	u, err := utils.GetUser(int64(userID))
	if err != nil {
		response.FailWithMessageWithCode(10002, "下单失败", c)
		return
	}
	allPrice := *ss.Multiple * price
	needPrice := float64(allPrice) * (float64(*ss.Bond) / 100)
	decrAmount := int(needPrice)
	if decrAmount > u.AvailableAmount {
		response.FailWithMessageWithCode(10002, "下单失败，金额不足", c)
		return
	}
	u.AvailableAmount = u.AvailableAmount - decrAmount
	u.FreezeAmount = u.FreezeAmount + decrAmount
	if err = global.GVA_DB.Where("id = ?", u.ID).Save(u).Error; err != nil {
		global.GVA_LOG.Error("update user err", zap.Error(err))
		response.FailWithMessageWithCode(10002, "下单失败", c)
		return
	}
	utils.AddAmountLog(int(u.ID), decrAmount, u.AvailableAmount, 2)
	order := &orders.Orders{
		User_id:    &userID,
		Account_id: &accountID,
		Price:      &price,
		Volume:     &volume,
		Direction:  &direction,
		Order_no:   utils.MD5(fmt.Sprintf("%d", time.Now().UnixNano())),
		SymbolID:   req.Symbol,
		SymbolName: ss.Name,
	}
	err = orderService.CreateOrders(order)
	if err != nil {
		response.FailWithMessageWithCode(10002, "下单失败", c)
		return
	}
	// grpc 调用下单接口
	res, err := global.GVA_GrpcCLient.SayHello(context.Background(), &pb.HelloRequest{Name: "name"})
	if err != nil {
		global.GVA_LOG.Error("SayHello", zap.Error(err))
	}
	global.GVA_LOG.Error("SayHello", zap.String("message", res.GetMessage()))

	response.OkWithMessage("success", c)
	return
}

// PriceData k线数据
// @Tags 前端接口API
// @Summary k线数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query apis.KData true "下单参数"
// @Success 200 {object} object "{"code":0,"data":{},"msg":"success"}"
// @Router /api/k/data [get]
func (uApi *ApisApi) PriceData(c *gin.Context) {
	var req apis.KData
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.Rows == 0 {
		req.Rows = 10
	}
	var yData []apis.YdData
	var dates interface{}
	dates = []data.Data{}
	if req.Period == 5 {
		dates = []data.Data5{}
	}
	if req.Period == 15 {
		dates = []data.Data15{}
	}
	if req.Period == 30 {
		dates = []data.Data30{}
	}
	if req.Period == 60 {
		dates = []data.Data60{}
	}
	if req.Period == 120 {
		dates = []data.Data120{}
	}
	if req.Period == 240 {
		dates = []data.Data240{}
	}
	if req.Period == 360 {
		dates = []data.Data360{}
	}
	if req.Period == 480 {
		dates = []data.Data480{}
	}
	if req.Period == 1440 {
		dates = []data.Data1440{}
	}
	err = global.GVA_DB.Where("symbol_id = ?", req.Symbol).Order("id DESC").Limit(int(req.Rows)).Find(&dates).Error
	if err != nil {
		global.GVA_LOG.Error("PriceData err", zap.Error(err))
		response.FailWithMessageWithCode(10002, "获取失败", c)
		return
	}
	var price float64
	switch v := dates.(type) {
	case []data.Data:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data5:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data15:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data30:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data60:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data120:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data240:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data360:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data480:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	case []data.Data1440:
		for k, d := range v {
			if k == 0 {
				price = d.PreClosePrice
			}
			yData = append(yData, apis.YdData{
				Uptime: d.InsertAt,
				Cjl:    d.OpenInterest, // 持仓量
				Close:  d.PreClosePrice,
				Vol:    float64(d.Volume),
				Open:   d.OpenInterest,
				High:   d.AveragePrice,
				Low:    d.LowerLimitPrice,
			})
		}
	default:
		response.FailWithMessageWithCode(10002, "获取失败", c)
		return
	}

	response.OkWithData(apis.KDataResp{
		YdClosePrice: price,
		Results:      yData,
	}, c)
	return
}

// SymbolData 行情列表
// @Tags 前端接口API
// @Summary 行情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} apis.SymbolDataResp "{"code":0,"data":{},"msg":"success"}"
// @Router /api/symbol/data/list [get]
func (uApi *ApisApi) SymbolData(c *gin.Context) {
	var sb []*symbols.Symbol
	if err := global.GVA_DB.Find(&sb).Error; err != nil {
		response.FailWithMessageWithCode(10002, "获取失败", c)
		return
	}
	var ds []apis.SymbolData
	for _, s := range sb {
		key := fmt.Sprintf("s:info:%s", s.Code)
		res, err := global.GVA_REDIS.Get(c.Request.Context(), key).Result()
		if err != nil {
			global.GVA_LOG.Error("SymbolData hgetall err ", zap.Error(err), zap.String("key", key))
			continue
		}
		var d apis.SymbolData
		if err = json.Unmarshal([]byte(res), &d); err != nil {
			global.GVA_LOG.Error("SymbolData Unmarshal err ", zap.Error(err), zap.String("res", res))
			continue
		}
		d.Name = s.Name
		ds = append(ds, d)

	}

	response.OkWithData(apis.SymbolDataResp{
		List: ds,
	}, c)
	return
}

// SymbolDataInfo 行情详情
// @Tags 前端接口API
// @Summary 行情详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query apis.ReqSymbolInfo true "参数"
// @Success 200 {object} apis.SymbolData "{"code":0,"data":{},"msg":"success"}"
// @Router /api/symbol/data/info [get]
func (uApi *ApisApi) SymbolDataInfo(c *gin.Context) {
	var req apis.ReqSymbolInfo
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var s *symbols.Symbol
	if err = global.GVA_DB.Where("code= ?", req.SymbolID).Find(&s).Error; err != nil {
		response.FailWithMessageWithCode(10002, "获取失败", c)
		return
	}
	key := fmt.Sprintf("s:info:%s", s.Code)
	res, err := global.GVA_REDIS.Get(c.Request.Context(), key).Result()
	if err != nil {
		global.GVA_LOG.Error("SymbolData hgetall err ", zap.Error(err), zap.String("key", key))
		response.FailWithMessageWithCode(10002, "获取失败", c)
		return
	}
	var d apis.SymbolData
	if err = json.Unmarshal([]byte(res), &d); err != nil {
		global.GVA_LOG.Error("SymbolData Unmarshal err ", zap.Error(err), zap.String("res", res))
		response.FailWithMessageWithCode(10002, "获取失败", c)
		return
	}
	d.Name = s.Name

	response.OkWithData(d, c)
	return
}

// GetUserInfo 获取用户信息
// @Tags 前端接口API
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} users.Users "{"code":0,"data":{},"msg":"success"}"
// @Router /api/user/info [get]
func (uApi *ApisApi) GetUserInfo(c *gin.Context) {

	id, _ := c.Get("uid")
	userID := cast.ToInt(id)
	var u users.Users
	err := global.GVA_DB.Where("id = ?", userID).First(&u).Error
	if err != nil {
		global.GVA_LOG.Error("GVA_DB get user err", zap.Error(err))
		response.FailWithMessageWithCode(10002, "获取失败", c)
		return
	}
	response.OkWithData(u, c)
	return
}

// OrdersList 交易记录
// @Tags 前端接口API
// @Summary 交易记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query apis.ReqTrade true "查询参数"
// @Success 200 {array} orders.Orders "{"code":0,"data":{},"msg":"success"}"
// @Router /api/orders/list [get]
func (uApi *ApisApi) OrdersList(c *gin.Context) {
	var req apis.ReqTrade
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	page := 1
	if req.Page > 1 {
		page = int(req.Page)
	}
	limit := 20
	offset := (page - 1) * limit
	id, _ := c.Get("uid")
	userID := cast.ToInt(id)
	var os []*orders.Orders
	err = global.GVA_DB.Where("user_id = ?", userID).Order("id DESC").Offset(offset).Limit(limit).Find(&os).Error
	if err != nil {
		global.GVA_LOG.Error("PriceData err", zap.Error(err))
		response.FailWithMessageWithCode(10002, "获取失败", c)
		return
	}
	response.OkWithData(os, c)
	return
}
