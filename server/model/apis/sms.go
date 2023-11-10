package apis

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 短信 结构体  Sms
type ReqSms struct {
	Phone string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"` // 手机号
}

type ReqRegister struct {
	Phone      string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"` // 手机号
	Code       string `json:"code" form:"code" `                                            // 验证码
	Password   string `json:"password" form:"password" `                                    // 密码
	InviteCode string `json:"invite_code" form:"password"`                                  // 邀请码
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
	Direction int64  `json:"direction" form:"direction" form:"symbol"` // 类型 1买 2卖
	Symbol    string `json:"symbol" form:"symbol" binding:"required" ` // 品种
}

type KData struct {
	Period int64  `json:"period" form:"period"`                     // 周期
	Rows   int64  `json:"rows" form:"rows" binding:"required" `     // 返回条数
	Symbol string `json:"symbol" form:"symbol" binding:"required" ` // 品种
}

type ArticleReq struct {
	ID int64 `json:"id" form:"id" binding:"required" ` // 文章id
}

type ArticleListReq struct {
	ArticleCategoryID int64  `json:"article_category_id"  form:"article_category_id" ` // 分类id
	Page              int64  `json:"page"  form:"page" binding:"required" `            // 分页
	Symbol            string `json:"symbol" form:"symbol"`                             // 合约
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
	Status int64 `json:"status" form:"status" `                // 状态 0下单中 1成交 2取消 3失败 4平
}

type SymbolDataResp struct {
	List []SymbolData
}

type SymbolData struct {
	global.GVA_MODEL
	SymbolId           string  `json:"symbol_id"  form:"symbol_id" gorm:"column:symbol_id;comment:合约代码;"`                                  // 合约代码
	InstrumentRef      int64   `json:"instrument_ref"  form:"instrument_ref" gorm:"column:instrument_ref;comment:合约序号;"`                   // 合约序号
	TradingDay         int64   `json:"trading_day"  form:"trading_day" gorm:"column:trading_day;comment:交易日;"`                             // 交易日
	PreSettlementPrice float64 `json:"pre_settlement_price"  form:"pre_settlement_price" gorm:"column:pre_settlement_price;comment:前结算价;"` // 前结算价
	PreClosePrice      float64 `json:"pre_close_price"  form:"pre_close_price" gorm:"column:pre_close_price;comment:前收盘价;"`                // 前收盘价
	PreOpenInterest    float64 `json:"pre_open_interest"  form:"pre_open_interest" gorm:"column:pre_open_interest;comment:前持仓量;"`          // 前持仓量
	UpperLimitPrice    float64 `json:"upper_limit_price"  form:"upper_limit_price" gorm:"column:upper_limit_price;comment:涨停板价;"`          // 涨停板价
	LowerLimitPrice    float64 `json:"lower_limit_price"  form:"lower_limit_price" gorm:"column:lower_limit_price;comment:跌停板价;"`          // 跌停板价
	LastPrice          float64 `json:"last_price"  form:"last_price" gorm:"column:last_price;comment:最新价;"`                                // 最新价
	BidPrice           float64 `json:"bid_price"  form:"bid_price" gorm:"column:bid_price;comment:买入价。为零代表无买入价。;"`                         // 买入价。为零代表无买入价。
	AskPrice           float64 `json:"ask_price"  form:"ask_price" gorm:"column:ask_price;comment:卖出价。为零代表无卖出价。;"`                         // 卖出价。为零代表无卖出价。
	BidVolume          int64   `json:"bid_volume"  form:"bid_volume" gorm:"column:bid_volume;comment:买入量。为零代表无买入价。;"`                      // 买入量。为零代表无买入价。
	AskVolume          int64   `json:"ask_volume"  form:"ask_volume" gorm:"column:ask_volume;comment:卖出量。为零代表无卖出价。;"`                      // 卖出量。为零代表无卖出价。
	Turnover           float64 `json:"turnover"  form:"turnover" gorm:"column:turnover;comment:成交金额;"`                                     // 成交金额
	OpenInterest       float64 `json:"open_interest"  form:"open_interest" gorm:"column:open_interest;comment:持仓量;"`                       // 持仓量
	Volume             int64   `json:"volume"  form:"volume" gorm:"column:volume;comment:成交量;"`                                            // 成交量
	TimeStamp          int64   `json:"time_stamp"  form:"time_stamp" gorm:"column:time_stamp;comment:时间戳;"`                                // 时间戳
	AveragePrice       float64 `json:"average_price"  form:"average_price" gorm:"column:average_price;comment:市场均价;"`                      // 市场均价
	InsertAt           int64   `json:"insert_at"  form:"insert_at" gorm:"column:insert_at;comment:插入时间;"`                                  // 插入时间
	Name               string
}

type ReqSymbolInfo struct {
	SymbolID string `json:"symbol_id" form:"symbol_id" binding:"required" ` // 行情合约编号

}
