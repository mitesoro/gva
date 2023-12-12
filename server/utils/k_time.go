package utils

import "time"

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
