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
	"time"
)

var (
	luaHighScript = `
		local current_price = tonumber(ARGV[1])
		local stored_price = redis.call('get', KEYS[1])
		if not stored_price then
			redis.call('set', KEYS[1], current_price, 'EX', 86400)
			return "NO_TIMES"
		else
			stored_price = tonumber(stored_price)
			if current_price > stored_price then
				redis.call('set', KEYS[1], current_price, 'EX', 86400)
			end
			return "OK"
		end

	`
	luaLowScript = `
		local current_price = tonumber(ARGV[1])
		local stored_price = redis.call('get', KEYS[1])
		if not stored_price  then
			redis.call('set', KEYS[1], current_price, 'EX', 86400)
			return "NO_TIMES"
		else
			stored_price = tonumber(stored_price)
			if current_price < stored_price then
				redis.call('set', KEYS[1], current_price, 'EX', 86400)
			end
			return "OK"
		end
	`
)

// DoKData 处理k线
func DoKData(d data.Data) {
	dateFormat := "2006-01-02-15:04"
	ctx := context.Background()
	now := time.Now()
	kd := kdata.KData{
		Uptime: now.Unix(),
	}
	key := now.Format(dateFormat)
	// 开始 开盘
	if err := global.GVA_REDIS.SetNX(ctx, fmt.Sprintf("k_data_1_start_%s_%s", key, d.SymbolId), d.LastPrice, 24*time.Hour).Err(); err != nil {
		global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
	}
	// 结束 收盘
	if err := global.GVA_REDIS.SetEx(ctx, fmt.Sprintf("k_data_1_end_%s_%s", key, d.SymbolId), d.LastPrice, 24*time.Hour).Err(); err != nil {
		global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
	}
	// 最高价
	highKey := fmt.Sprintf("k_data_1_high_%s_%s", key, d.SymbolId)
	if err := global.GVA_REDIS.Eval(ctx, luaHighScript, []string{highKey}, d.LastPrice).Err(); err != nil {
		global.GVA_LOG.Error("DoKData luaHighScript1 Eval err:", zap.Error(err), zap.Any("d", d))
	}
	// 最低价
	lowKey := fmt.Sprintf("k_data_1_low_%s_%s", key, d.SymbolId)
	if err := global.GVA_REDIS.Eval(ctx, luaLowScript, []string{lowKey}, d.LastPrice).Err(); err != nil {
		global.GVA_LOG.Error("DoKData luaLowScript1 Eval err:", zap.Error(err), zap.Any("d", d))
	}

	if now.Minute()%1 == 0 && now.Second() == 0 { // 分钟
		lockKey := fmt.Sprintf("lock_k_data_%d_%s", now.Unix(), d.SymbolId)
		firstLock := utils.NewRedisLock(global.GVA_REDIS, lockKey)
		firstLock.SetExpire(5)
		againAcquire, err := firstLock.Acquire(context.Background())
		if err != nil {
			global.GVA_LOG.Error("DoKData Acquire err:", zap.Error(err), zap.Any("d", d))
		}
		if !againAcquire {
			return
		}
		dataMinute := now.Add(-1 * time.Minute).Format(dateFormat)
		// 获取开盘
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_start_%s_%s", dataMinute, d.SymbolId)).Result(); err1 == nil {
			kd.Open = cast.ToInt64(res)
		}
		// 获取收盘
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_end_%s_%s", dataMinute, d.SymbolId)).Result(); err1 == nil {
			kd.Close = cast.ToInt64(res)
		}
		// 获取最高
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_high_%s_%s", dataMinute, d.SymbolId)).Result(); err1 == nil {
			kd.High = cast.ToInt64(res)
		}
		// 获取最低
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_low_%s_%s", dataMinute, d.SymbolId)).Result(); err1 == nil {
			kd.Low = cast.ToInt64(res)
		} else {
			global.GVA_LOG.Error("DoKData:", zap.Error(err1), zap.Any("res", res))
		}
		kd.SymbolId = d.SymbolId
		if err = global.GVA_DB.Create(&kd).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储每分钟数据
		value := map[string]interface{}{
			"open":  kd.Open,
			"high":  kd.High,
			"low":   kd.Low,
			"close": kd.Close,
		}
		if err = global.GVA_REDIS.HMSet(ctx, fmt.Sprintf("k_data_%s_%s", key, d.SymbolId), value).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
	}
	if now.Minute()%5 == 0 && now.Second() == 0 { // 5分钟
		kd5 := kdata.KData5(kd)
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data_%s_%s", now.Add(-1*time.Minute).Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data_%s_%s", now.Add(-2*time.Minute).Format(dateFormat), d.SymbolId))
		v3 := utils.GetKd(ctx, fmt.Sprintf("k_data_%s_%s", now.Add(-3*time.Minute).Format(dateFormat), d.SymbolId))
		v4 := utils.GetKd(ctx, fmt.Sprintf("k_data_%s_%s", now.Add(-4*time.Minute).Format(dateFormat), d.SymbolId))
		v5 := utils.GetKd(ctx, fmt.Sprintf("k_data_%s_%s", now.Add(-5*time.Minute).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low, v3.Low, v4.Low, v5.Low})
		max := utils.FindMax([]int64{v1.High, v2.High, v3.High, v4.High, v5.High})
		kd5.Open = v5.Open
		kd5.Close = v1.Close
		kd5.High = max
		kd5.Low = mix
		kd5.SymbolId = d.SymbolId
		if err := global.GVA_DB.Create(&kd5).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
	}
	if now.Minute()%15 == 0 && now.Second() == 0 { // 15分钟
		kd15 := kdata.KData15(kd)
		if err := global.GVA_DB.Create(&kd15).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
	}
}
