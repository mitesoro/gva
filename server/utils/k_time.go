package utils

import (
	"fmt"
	"time"
)

func IsK30(now time.Time) bool {
	if now.Hour() == 9 && now.Minute() == 31 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 10 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 10 && now.Minute() == 46 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 11 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 13 && now.Minute() == 46 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 14 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 14 && now.Minute() == 46 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 15 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 21 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 21 && now.Minute() == 31 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 22 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 22 && now.Minute() == 31 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 23 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	return false
}

func IsK60(now time.Time) bool {
	if now.Hour() == 10 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 11 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 13 && now.Minute() == 46 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 14 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 15 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	if now.Hour() == 22 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 23 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	return false
}

func IsK120(now time.Time) bool {
	if now.Hour() == 11 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 15 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 23 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	return false
}

func IsK240(now time.Time) bool {
	if now.Hour() == 11 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 15 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	return false
}

// IsK30Ye 夜盘
func IsK30Ye(now time.Time) bool {
	if now.Hour() == 9 && now.Minute() == 31 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 10 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 10 && now.Minute() == 46 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 11 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 13 && now.Minute() == 46 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 14 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 14 && now.Minute() == 46 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 15 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 21 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 21 && now.Minute() == 31 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 22 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 22 && now.Minute() == 31 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 23 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 23 && now.Minute() == 31 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 0 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 0 && now.Minute() == 31 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 1 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	return false
}

func IsK60Ye(now time.Time) bool {
	if now.Hour() == 10 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 11 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 13 && now.Minute() == 46 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 14 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 15 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	if now.Hour() == 22 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 23 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 0 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 1 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	return false
}

func IsK120Ye(now time.Time) bool {
	if now.Hour() == 11 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 15 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 23 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 1 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	return false
}

func IsK240Ye(now time.Time) bool {
	if now.Hour() == 11 && now.Minute() == 16 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 15 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}
	if now.Hour() == 1 && now.Minute() == 1 && now.Second() == 0 {
		return true
	}

	return false
}

func IsYe(code string) bool {
	if code == "ni2401" {
		return true
	}
	if code == "al2401" {
		return true
	}

	return false
}

func IsCloseTime(now time.Time) bool {
	if now.Hour() == 15 && now.Minute() == 1 {
		return true
	}
	if now.Hour() == 10 && now.Minute() == 16 {
		return true
	}
	if now.Hour() == 11 && now.Minute() == 31 {
		return true
	}
	if now.Hour() == 23 && now.Minute() == 1 {
		return true
	}
	return false
}

func IsCloseTimeYe(now time.Time) bool {
	if now.Hour() == 15 && now.Minute() == 1 {
		return true
	}
	if now.Hour() == 10 && now.Minute() == 16 {
		return true
	}
	if now.Hour() == 11 && now.Minute() == 31 {
		return true
	}
	if now.Hour() == 1 && now.Minute() == 1 {
		return true
	}
	return false
}

func IsOpenTime(now time.Time) bool {
	if now.Hour() == 9 && now.Minute() == 0 {
		return true
	}
	return false
}

func GetTime(inputDate int64) int64 {

	// 指定时区
	timezone := "Asia/Shanghai"
	location, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println("时区加载错误:", err)
		return 0
	}

	// 将整数转换为时间对象，并指定时区
	dateTime, err := time.ParseInLocation("20060102", fmt.Sprint(inputDate), location)
	if err != nil {
		fmt.Println("日期解析错误:", err)
		return 0
	}

	// 将时间的时、分、秒、纳秒部分清零
	zeroTime := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, location)

	// 转换为时间戳
	return zeroTime.Unix()
}
