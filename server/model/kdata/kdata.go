package kdata

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type KData struct {
	global.GVA_MODEL
	SymbolId string  `json:"symbol_id"  form:"symbol_id" gorm:"column:symbol_id;comment:合约代码;"` // 合约代码
	Uptime   int64   `json:"uptime"  form:"uptime" gorm:"column:uptime;comment:更新时间;"`          // 更新时间
	Open     float64 `json:"open"  form:"open" gorm:"column:open;comment:开盘价;"`                 // 开盘价
	High     float64 `json:"high"  form:"high" gorm:"column:high;comment:最高价;"`                 // 最高价
	Low      float64 `json:"low"  form:"low" gorm:"column:low;comment:最低价;"`                    // 最低价
	Close    float64 `json:"close"  form:"close" gorm:"column:close;comment:收盘价;"`              // 收盘价
	Vol      float64 `json:"vol"  form:"vol" gorm:"column:vol;comment:成交量;"`                    // 成交量
	Cjl      float64 `json:"cjl"  form:"cjl" gorm:"column:cjl;comment:成交量;"`                    // 成交量
}

type KData5 KData

type KData15 KData

type KData30 KData

type KData60 KData

type KData120 KData

type KData240 KData

type KData360 KData

type KData480 KData

type KData1440 KData
