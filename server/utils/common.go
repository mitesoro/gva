package utils

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"go.uber.org/zap"
	"strconv"
	"time"
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

func IsWithinBusinessHours(t time.Time, start string, end string) bool {
	// Parse the start and end times
	layout := "15:04"
	parsedStartTime, _ := time.Parse(layout, start)
	parsedEndTime, _ := time.Parse(layout, end)

	// Combine the current date with the parsed start and end times
	startTime := time.Date(t.Year(), t.Month(), t.Day(), parsedStartTime.Hour(), parsedStartTime.Minute(), 0, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), parsedEndTime.Hour(), parsedEndTime.Minute(), 0, 0, t.Location())

	// Check if the time is between start and end
	isBetween := t.After(startTime) && t.Before(endTime)

	// Check if the day of the week is between Monday and Friday
	isWeekday := t.Weekday() >= time.Monday && t.Weekday() <= time.Friday

	return isBetween && isWeekday
}

func IsWithinRange(t time.Time, start string, end string) bool {
	// Define the layout string for parsing the start and end times
	layout := "2006-01-02 15:04:05"

	// Parse the start and end times
	startTime, _ := time.Parse(layout, start)
	endTime, _ := time.Parse(layout, end)

	// Check if the current time is between start and end
	isBetween := t.After(startTime) && t.Before(endTime)

	return isBetween
}
