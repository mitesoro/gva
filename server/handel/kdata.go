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
		rKey := fmt.Sprintf("k_data_%s_%s", key, d.SymbolId)
		if err = global.GVA_REDIS.HMSet(ctx, rKey, value).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}
	if (now.Minute()%10 == 6 && now.Second() == 0) || (now.Minute()%10 == 1 && now.Second() == 0) { // 5分钟
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
		kd5.Uptime = now.Add(-1 * time.Minute).Unix()
		if err := global.GVA_DB.Create(&kd5).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储5分钟数据
		value5 := map[string]interface{}{
			"open":  kd5.Open,
			"high":  kd5.High,
			"low":   kd5.Low,
			"close": kd5.Close,
		}
		rKey := fmt.Sprintf("k_data5_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value5).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}
	if (now.Minute() == 16 && now.Second() == 0) || (now.Minute() == 31 && now.Second() == 0) ||
		(now.Minute() == 46 && now.Second() == 0) || (now.Minute() == 1 && now.Second() == 0) { // 15分钟(15,30,45,00)
		kd15 := kdata.KData15(kd)
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data5_%s_%s", now.Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data5_%s_%s", now.Add(-5*time.Minute).Format(dateFormat), d.SymbolId))
		v3 := utils.GetKd(ctx, fmt.Sprintf("k_data5_%s_%s", now.Add(-10*time.Minute).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low, v3.Low})
		max := utils.FindMax([]int64{v1.High, v2.High, v3.High})
		kd15.Open = v3.Open
		kd15.Close = v1.Close
		kd15.High = max
		kd15.Low = mix
		kd15.SymbolId = d.SymbolId
		kd15.Uptime = now.Add(-1 * time.Minute).Unix()
		if err := global.GVA_DB.Create(&kd15).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储15分钟数据
		value15 := map[string]interface{}{
			"open":  kd15.Open,
			"high":  kd15.High,
			"low":   kd15.Low,
			"close": kd15.Close,
		}
		rKey := fmt.Sprintf("k_data15_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value15).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}

	if (now.Minute()%31 == 0 && now.Second() == 0) || (now.Minute()%1 == 0 && now.Second() == 0) { // 30分钟(31,01)
		kd30 := kdata.KData30(kd)
		kd30.Uptime = now.Add(-1 * time.Minute).Unix()
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data15_%s_%s", now.Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data15_%s_%s", now.Add(-15*time.Minute).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low})
		max := utils.FindMax([]int64{v1.High, v2.High})
		kd30.Open = v2.Open
		kd30.Close = v1.Close
		kd30.High = max
		kd30.Low = mix
		kd30.SymbolId = d.SymbolId
		if err := global.GVA_DB.Create(&kd30).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储15分钟数据
		value30 := map[string]interface{}{
			"open":  kd30.Open,
			"high":  kd30.High,
			"low":   kd30.Low,
			"close": kd30.Close,
		}
		rKey := fmt.Sprintf("k_data30_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value30).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}
	if now.Minute() == 1 && now.Second() == 0 { // 小时线
		kd60 := kdata.KData60(kd)
		kd60.Uptime = now.Add(-1 * time.Minute).Unix()
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data30_%s_%s", now.Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data30_%s_%s", now.Add(-30*time.Minute).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low})
		max := utils.FindMax([]int64{v1.High, v2.High})
		kd60.Open = v2.Open
		kd60.Close = v1.Close
		kd60.High = max
		kd60.Low = mix
		kd60.SymbolId = d.SymbolId
		if err := global.GVA_DB.Create(&kd60).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储15分钟数据
		value60 := map[string]interface{}{
			"open":  kd60.Open,
			"high":  kd60.High,
			"low":   kd60.Low,
			"close": kd60.Close,
		}
		rKey := fmt.Sprintf("k_data60_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value60).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}

	if now.Minute() == 1 && now.Second() == 0 && now.Hour()%2 == 0 { // 2小时
		kd120 := kdata.KData120(kd)
		kd120.Uptime = now.Add(-1 * time.Minute).Unix()
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data60_%s_%s", now.Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data60_%s_%s", now.Add(-60*time.Minute).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low})
		max := utils.FindMax([]int64{v1.High, v2.High})
		kd120.Open = v2.Open
		kd120.Close = v1.Close
		kd120.High = max
		kd120.Low = mix
		kd120.SymbolId = d.SymbolId
		if err := global.GVA_DB.Create(&kd120).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储15分钟数据
		value120 := map[string]interface{}{
			"open":  kd120.Open,
			"high":  kd120.High,
			"low":   kd120.Low,
			"close": kd120.Close,
		}
		rKey := fmt.Sprintf("k_data120_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value120).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}
	if now.Minute() == 1 && now.Second() == 0 && now.Hour()%4 == 0 { // 4小时
		kd240 := kdata.KData240(kd)
		kd240.Uptime = now.Add(-1 * time.Minute).Unix()
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data120_%s_%s", now.Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data120_%s_%s", now.Add(-2*time.Hour).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low})
		max := utils.FindMax([]int64{v1.High, v2.High})
		kd240.Open = v2.Open
		kd240.Close = v1.Close
		kd240.High = max
		kd240.Low = mix
		kd240.SymbolId = d.SymbolId
		if err := global.GVA_DB.Create(&kd240).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储15分钟数据
		value240 := map[string]interface{}{
			"open":  kd240.Open,
			"high":  kd240.High,
			"low":   kd240.Low,
			"close": kd240.Close,
		}
		rKey := fmt.Sprintf("k_data240_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value240).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}
	if now.Minute() == 1 && now.Second() == 0 && now.Hour()%6 == 0 { // 6小时
		kd360 := kdata.KData360(kd)
		kd360.Uptime = now.Add(-1 * time.Minute).Unix()
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data120_%s_%s", now.Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data120_%s_%s", now.Add(-2*time.Hour).Format(dateFormat), d.SymbolId))
		v3 := utils.GetKd(ctx, fmt.Sprintf("k_data120_%s_%s", now.Add(-4*time.Hour).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low, v3.Low})
		max := utils.FindMax([]int64{v1.High, v2.High, v3.High})
		kd360.Open = v3.Open
		kd360.Close = v1.Close
		kd360.High = max
		kd360.Low = mix
		kd360.SymbolId = d.SymbolId
		if err := global.GVA_DB.Create(&kd360).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储15分钟数据
		value360 := map[string]interface{}{
			"open":  kd360.Open,
			"high":  kd360.High,
			"low":   kd360.Low,
			"close": kd360.Close,
		}
		rKey := fmt.Sprintf("k_data360_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value360).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}

	if now.Minute() == 1 && now.Second() == 0 && now.Hour()%8 == 0 { // 8小时
		kd480 := kdata.KData480(kd)
		kd480.Uptime = now.Add(-1 * time.Minute).Unix()
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data240_%s_%s", now.Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data240_%s_%s", now.Add(-4*time.Hour).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low})
		max := utils.FindMax([]int64{v1.High, v2.High})
		kd480.Open = v2.Open
		kd480.Close = v1.Close
		kd480.High = max
		kd480.Low = mix
		kd480.SymbolId = d.SymbolId
		if err := global.GVA_DB.Create(&kd480).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储15分钟数据
		value480 := map[string]interface{}{
			"open":  kd480.Open,
			"high":  kd480.High,
			"low":   kd480.Low,
			"close": kd480.Close,
		}
		rKey := fmt.Sprintf("k_data480_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value480).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}
	if now.Minute() == 1 && now.Second() == 0 && now.Hour() == 0 { // 24小时
		kd1440 := kdata.KData1440(kd)
		kd1440.Uptime = now.Add(-1 * time.Minute).Unix()
		v1 := utils.GetKd(ctx, fmt.Sprintf("k_data480_%s_%s", now.Format(dateFormat), d.SymbolId))
		v2 := utils.GetKd(ctx, fmt.Sprintf("k_data480_%s_%s", now.Add(-8*time.Hour).Format(dateFormat), d.SymbolId))
		v3 := utils.GetKd(ctx, fmt.Sprintf("k_data480_%s_%s", now.Add(-16*time.Hour).Format(dateFormat), d.SymbolId))
		mix := utils.FindMin([]int64{v1.Low, v2.Low, v3.Low})
		max := utils.FindMax([]int64{v1.High, v2.High, v3.High})
		kd1440.Open = v3.Open
		kd1440.Close = v1.Close
		kd1440.High = max
		kd1440.Low = mix
		kd1440.SymbolId = d.SymbolId
		if err := global.GVA_DB.Create(&kd1440).Error; err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		// 存储15分钟数据
		value1440 := map[string]interface{}{
			"open":  kd1440.Open,
			"high":  kd1440.High,
			"low":   kd1440.Low,
			"close": kd1440.Close,
		}
		rKey := fmt.Sprintf("k_data1440_%s_%s", key, d.SymbolId)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value1440).Err(); err != nil {
			global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)
	}
}
