package handel

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data"
	"github.com/flipped-aurora/gin-vue-admin/server/model/kdata"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"log"
	"time"
)

var (
	luaHighScript = `
		local current_price = tonumber(ARGV[1])
		local stored_price = tonumber(redis.call('get', KEYS[1]))
		if stored_price == nil or current_price > stored_price then
			redis.call('set', KEYS[1], current_price)
		end
	`
	luaLowScript = `
		local current_price = tonumber(ARGV[1])
		local stored_price = tonumber(redis.call('get', KEYS[1]))
		if stored_price == nil or current_price < stored_price then
			redis.call('set', KEYS[1], current_price)
		end
	`
)

// DoKData 处理k线
func DoKData(d data.Data) {
	dateFormat := "2006-01-02 15:04"
	ctx := context.Background()
	now := time.Now()
	kd := kdata.KData{
		Uptime: now.Unix(),
	}
	key := now.Format(dateFormat)
	// 开始 开盘
	if err := global.GVA_REDIS.SetNX(ctx, fmt.Sprintf("k_data_1_start_%s", key), d.LastPrice, 2*time.Minute).Err(); err != nil {
		global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
	}
	// 结束 收盘
	if err := global.GVA_REDIS.SetEx(ctx, fmt.Sprintf("k_data_1_end_%s", key), d.LastPrice, 2*time.Minute).Err(); err != nil {
		global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
	}
	// 最高价
	highKey := fmt.Sprintf("k_data_1_high_%s", key)
	if err := global.GVA_REDIS.Eval(ctx, luaHighScript, []string{highKey}, d.LastPrice).Err(); err != nil {
		global.GVA_LOG.Error("DoKData Eval err:", zap.Error(err), zap.Any("d", d))
	}
	// 最低价
	lowKey := fmt.Sprintf("k_data_1_low_%s", key)
	if err := global.GVA_REDIS.Eval(ctx, luaHighScript, []string{lowKey}, d.LastPrice).Err(); err != nil {
		global.GVA_LOG.Error("DoKData Eval err:", zap.Error(err), zap.Any("d", d))
	}

	if now.Minute()%1 == 0 && now.Second() == 0 { // 分钟
		lockKey := fmt.Sprintf("lock_k_data_%d", now.Unix())
		firstLock := utils.NewRedisLock(global.GVA_REDIS, lockKey)
		firstLock.SetExpire(5)
		againAcquire, err := firstLock.Acquire(context.Background())
		if err != nil {
			global.GVA_LOG.Error("DoKData Acquire err:", zap.Error(err), zap.Any("d", d))
		}
		if !againAcquire {
			return
		}
		dataMinute := now.Add(-1 * time.Minute)
		// 获取开盘
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_start_%s", dataMinute)).Result(); err1 == nil {
			kd.Open = cast.ToFloat64(res)
		}
		// 获取收盘
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_end_%s", dataMinute)).Result(); err1 == nil {
			kd.Close = cast.ToFloat64(res)
		}
		// 获取最高
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_high_%s", dataMinute)).Result(); err1 == nil {
			kd.High = cast.ToFloat64(res)
		}
		// 获取最低
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_low_%s", dataMinute)).Result(); err1 == nil {
			kd.Low = cast.ToFloat64(res)
		}
		if err = global.GVA_DB.Create(&kd).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			return
		}
		// 存储每分钟数据
		value := map[string]interface{}{
			"open":  kd.Open,
			"high":  kd.High,
			"low":   kd.Low,
			"close": kd.Close,
		}
		if err = global.GVA_REDIS.HMSet(ctx, fmt.Sprintf("k_data_%s", key), value).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			return
		}
	}
	if now.Minute()%5 == 0 && now.Second() == 0 { // 5分钟
		key5 := now.Add(-1 * time.Minute).Format(dateFormat)
		key4 := now.Add(-2 * time.Minute).Format(dateFormat)
		key3 := now.Add(-3 * time.Minute).Format(dateFormat)
		key2 := now.Add(-4 * time.Minute).Format(dateFormat)
		key1 := now.Add(-5 * time.Minute).Format(dateFormat)
		log.Println(key1, key2, key3, key4, key5)

		kd5 := kdata.KData5(kd)
		if err := global.GVA_DB.Create(&kd5).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			return
		}
	}
}
