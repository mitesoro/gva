package utils

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/message"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// AddMessage 添加消息
func AddMessage(userID int64, msg string) {
	al := message.Message{
		UserId:  userID,
		Content: msg,
	}
	if err := global.GVA_DB.Create(&al).Error; err != nil {
		global.GVA_LOG.Error("AddMessage err", zap.Error(err))
	}
}

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
	layout := "15:04:05" // Added seconds to the layout
	parsedStartTime, _ := time.Parse(layout, start)
	parsedEndTime, _ := time.Parse(layout, end)

	// Combine the current date with the parsed start and end times
	startTime := time.Date(t.Year(), t.Month(), t.Day(), parsedStartTime.Hour(), parsedStartTime.Minute(), parsedStartTime.Second(), 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), parsedEndTime.Hour(), parsedEndTime.Minute(), parsedEndTime.Second(), 0, t.Location())

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

func RandStr(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b []rune
	for i := 0; i < length; i++ {
		b = append(b, chars[r.Intn(len(chars))])
	}
	return string(b)
}

func GetKDataValue(ctx context.Context, key string) float64 {
	v, _ := global.GVA_REDIS.Get(ctx, key).Result()
	return cast.ToFloat64(v)
}

func FindMin(arr []int64) int64 {
	min := int64(math.MaxInt64)
	for _, v := range arr {
		if v == 0 {
			continue
		}
		if v < min {
			min = v
		}
	}
	return min
}

func FindMax(arr []int64) int64 {
	max := int64(-math.MaxInt64)
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

type KD struct {
	Open  int64 `json:"open"  form:"open" gorm:"column:open;comment:开盘价;"`    // 开盘价
	High  int64 `json:"high"  form:"high" gorm:"column:high;comment:最高价;"`    // 最高价
	Low   int64 `json:"low"  form:"low" gorm:"column:low;comment:最低价;"`       // 最低价
	Close int64 `json:"close"  form:"close" gorm:"column:close;comment:收盘价;"` // 收盘价
}

func GetKd(ctx context.Context, key string) KD {
	res, _ := global.GVA_REDIS.HGetAll(ctx, key).Result()
	return KD{
		Open:  cast.ToInt64(res["open"]),
		High:  cast.ToInt64(res["high"]),
		Low:   cast.ToInt64(res["low"]),
		Close: cast.ToInt64(res["close"]),
	}
}
