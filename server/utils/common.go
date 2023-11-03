package utils

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"go.uber.org/zap"
	"strconv"
)

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// AddAmountLog 添加记录
func AddAmountLog(userID, amount, curAmount, logType int) {
	al := alog.Alog{
		User_id:     &userID,
		Amount:      &amount,
		Amount_type: &logType,
		Cur_amount:  &curAmount,
	}
	if err := global.GVA_DB.Create(&al).Error; err != nil {
		global.GVA_LOG.Error("AddAmountLog err", zap.Error(err))
	}
}

// GetUser 获取用户
func GetUser(userID int64) (*users.Users, error) {
	var u users.Users
	if err := global.GVA_DB.Where("id = ?", userID).First(&u).Error; err != nil {
		global.GVA_LOG.Error("GetUserAvailableAmount err", zap.Error(err))
		return nil, err
	}

	return &u, nil
}

// GetSymbol 获取品种下单价格
func GetSymbol(code string) (*symbols.Symbol, error) {
	var s *symbols.Symbol
	if err := global.GVA_DB.Where("code = ?", code).First(&s).Error; err != nil {
		global.GVA_LOG.Error("GetSymbolPrice err", zap.Error(err))
		return nil, err
	}

	return s, nil
}
