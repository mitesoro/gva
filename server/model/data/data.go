package data

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// k线数据 结构体  Data
type Data struct {
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
}
