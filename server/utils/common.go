package utils

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alog"
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
