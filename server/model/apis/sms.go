package apis

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 短信 结构体  Sms
type ReqSms struct {
	Phone string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"` // 手机号
}

type ReqRegister struct {
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"` // 手机号
	Code     string `json:"code" form:"code" `                                            // 验证码
	Password string `json:"password" form:"password" `                                    // 密码
}

type ReqLogin struct {
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"` // 手机号
	Password string `json:"password" form:"password" `                                    // 密码
}

// 用户 结构体  Users
type Sms struct {
	global.GVA_MODEL
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"`     // 手机号
	Code     string `json:"code" form:"code" gorm:"column:code;comment:验证码;"`                 // 验证码
	ExpireAt int64  `json:"expire_at" form:"expire_at" gorm:"column:expire_at;comment:过期时间;"` // 过期时间
	Status   int64  `json:"status" form:"status" gorm:"column:status;comment:状态;"`            // 状态 0 可用 1已使用
}

type ReqUpdateUser struct {
	Avatar   string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:手机号;size:129;"` // 头像
	Nickname string `json:"nickname" form:"nickname" `                                       // 昵称
}

type ReqUpdatePhone struct {
	Phone string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"` // 手机号
	Code  string `json:"code" form:"code" `                                            // 验证码
}

type ReqUpdatePassword struct {
	Password    string `json:"password" form:"password" `         // 密码
	OldPassword string `json:"old_password" form:"old_password" ` // 旧密码
}

type ReqOrders struct {
	Volume    int64  `json:"volume" form:"volume" binding:"required" ` // 几手
	Price     int64  `json:"price" form:"price" form:"symbol"`         // 价格(分)
	Direction int64  `json:"direction" form:"direction" form:"symbol"` // 类型 1止赢 2止损
	Symbol    string `json:"symbol" form:"symbol" binding:"required" ` // 品种
}

type KData struct {
	Period int64  `json:"period" form:"period"`                     // 周期
	Rows   int64  `json:"rows" form:"rows" binding:"required" `     // 返回条数
	Symbol string `json:"symbol" form:"symbol" binding:"required" ` // 品种
}

type KDataResp struct {
	YdClosePrice float64  `json:"yd_close_price"`
	Results      []YdData `json:"results"`
}

type YdData struct {
	Uptime int64   `json:"uptime"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Vol    float64 `json:"vol"`
	Cjl    float64 `json:"cjl"`
}

type ReqTrade struct {
	Page   int64 `json:"page" form:"page" binding:"required" ` // 分页
	Status int64 `json:"status" form:"status" `                // 状态 0下单中 1成功 2取消 3失败 4盈利 5平
}
