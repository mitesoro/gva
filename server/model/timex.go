package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	if tTime.IsZero() {
		return []byte("\"--\""), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// UnmarshalJSON 从 JSON 字符串解析 LocalTime
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return err
	}

	*t = LocalTime(parsedTime)
	return nil
}

// Value 实现 driver.Valuer 接口，用于数据库插入
func (t LocalTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// Scan 实现 sql.Scanner 接口，用于从数据库扫描值
func (t *LocalTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*t = LocalTime(v)
		return nil
	default:
		return fmt.Errorf("Unsupported type for LocalTime: %T", value)
	}
}
