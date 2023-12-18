package handel

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data"
	"github.com/flipped-aurora/gin-vue-admin/server/model/kdata"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
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
	ctx := context.Background()
	dateFormat := "2006-01-02-15:04"
	now := time.Now()
	kd := kdata.KData{
		Uptime: now.Unix(),
	}
	key := now.Format(dateFormat)
	// 日k 开盘
	if err := global.GVA_REDIS.SetNX(ctx, fmt.Sprintf("k_data_day_start_%s_%s", now.Format("2006-01-02"), d.SymbolId), d.LastPrice, 24*time.Hour).Err(); err != nil {
		global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
	}
	// 日k 收盘
	if err := global.GVA_REDIS.SetEx(ctx, fmt.Sprintf("k_data_day_end_%s_%s", now.Format("2006-01-02"), d.SymbolId), d.LastPrice, 24*time.Hour).Err(); err != nil {
		global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
	}
	// 日k 最高价
	highDayKey := fmt.Sprintf("k_data_day_high_%s_%s", now.Format("2006-01-02"), d.SymbolId)
	if err := global.GVA_REDIS.Eval(ctx, luaHighScript, []string{highDayKey}, d.LastPrice).Err(); err != nil {
		global.GVA_LOG.Error("DoKData luaHighScript1 Eval err:", zap.Error(err), zap.Any("d", d))
	}
	// 日k 最低价
	lowLowKey := fmt.Sprintf("k_data_day_low_%s_%s", now.Format("2006-01-02"), d.SymbolId)
	if err := global.GVA_REDIS.Eval(ctx, luaLowScript, []string{lowLowKey}, d.LastPrice).Err(); err != nil {
		global.GVA_LOG.Error("DoKData luaLowScript1 Eval err:", zap.Error(err), zap.Any("d", d))
	}
	// 日k 交易日
	if err := global.GVA_REDIS.SetEx(ctx, fmt.Sprintf("k_data_day_trading_day_%s_%s", now.Format("2006-01-02"), d.SymbolId), d.TradingDay, 24*time.Hour).Err(); err != nil {
		global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("d", d))
	}

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
	// 有数据丢失，废弃
	if now.Minute()%1 == 0 && now.Second() == 0 && false { // 分钟
		lockKey := fmt.Sprintf("lock_k_data_%d_%s", now.Unix(), d.SymbolId)
		firstLock := utils.NewRedisLock(global.GVA_REDIS, lockKey)
		firstLock.SetExpire(5)
		againAcquire, err := firstLock.Acquire(context.Background())
		if err != nil {
			global.GVA_LOG.Error("DoKData Acquire err:", zap.Error(err), zap.Any("d", d))
		}
		if !againAcquire {
			global.GVA_LOG.Error("DoKData luck", zap.Any("now", now), zap.Any("d", d))
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
		if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
			return
		}
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

	return
}

// KDataDB 落库
func KDataDB() {
	// 计算到下一分钟的等待时间
	next := time.Now().Truncate(time.Minute).Add(time.Minute)
	wait := time.Until(next)

	// 等待到下一分钟
	time.Sleep(wait)
	// 创建一个每分钟触发一次的Ticker
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				add()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func add() {
	ctx := context.Background()
	dateFormat := "2006-01-02-15:04"
	now := time.Now()
	if now.Hour()%2 == 0 && now.Minute() == 0 && now.Second() == 0 {
		global.GVA_LOG.Error("add", zap.Any("time", time.Now().Format(time.DateTime)))
	}
	if now.Hour() == 8 && now.Minute() < 59 {
		return
	}
	if now.Hour() == 20 && now.Minute() < 59 {
		return
	}
	if now.Hour() == 15 && now.Minute() > 1 {
		return
	}

	var ss []symbols.Symbol
	if err := global.GVA_DB.Find(&ss).Error; err != nil {
		global.GVA_LOG.Error("Symbol err", zap.Error(err))
		return
	}
	if now.Hour()%2 == 0 && now.Minute() == 0 && now.Second() == 0 {
		global.GVA_LOG.Error("symbols", zap.Any("ss", ss))
	}
	for _, sss := range ss {
		kd := kdata.KData{
			Uptime: now.Unix(),
		}
		dataMinute := now.Add(-1 * time.Minute).Format(dateFormat)
		// 获取开盘
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_start_%s_%s", dataMinute, sss.Code)).Result(); err1 == nil {
			kd.Open = cast.ToInt64(utils.Decimal(cast.ToFloat64(res)))
		}
		// 获取收盘
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_end_%s_%s", dataMinute, sss.Code)).Result(); err1 == nil {
			kd.Close = cast.ToInt64(utils.Decimal(cast.ToFloat64(res)))
		}
		// 获取最高
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_high_%s_%s", dataMinute, sss.Code)).Result(); err1 == nil {
			kd.High = cast.ToInt64(utils.Decimal(cast.ToFloat64(res)))
		}
		// 获取最低
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_1_low_%s_%s", dataMinute, sss.Code)).Result(); err1 == nil {
			kd.Low = cast.ToInt64(utils.Decimal(cast.ToFloat64(res)))
		} else {
			global.GVA_LOG.Error("DoKData:", zap.Error(err1), zap.Any("res", res))
		}
		if utils.IsOpenTime(now) {
			// 设置开盘价
			if err := global.GVA_REDIS.Set(ctx, fmt.Sprintf("k_data_1_start_%s_%s", now.Format(dateFormat), sss.Code), kd.Open, 24*time.Hour).Err(); err != nil {
				global.GVA_LOG.Error("DoKData:", zap.Error(err), zap.Any("kd", kd))
			}
			continue
		}
		kd.SymbolId = sss.Code
		if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
			global.GVA_LOG.Error("add err:", zap.Any("kd", kd))
			continue
		}
		if utils.IsCloseTime(now) && !utils.IsYe(sss.Code) {
			global.GVA_LOG.Error("update k_data close:", zap.Any("close", kd.Close), zap.Any("time", now.Add(-1*time.Minute).Unix()))
			if err := global.GVA_DB.Model(kdata.KData{}).Where("uptime = ? and symbol_id = ?", now.Add(-1*time.Minute).Unix(), sss.Code).Update("close", kd.Close).Error; err != nil {
				global.GVA_LOG.Error("update k_data err:", zap.Error(err), zap.Any("kd", kd))
			}
			continue
		}
		if utils.IsCloseTimeYe(now) && utils.IsYe(sss.Code) {
			global.GVA_LOG.Error("update k_data close:", zap.Any("close", kd.Close), zap.Any("time", now.Add(-1*time.Minute).Unix()))
			if err := global.GVA_DB.Model(kdata.KData{}).Where("uptime = ? and symbol_id = ?", now.Add(-1*time.Minute).Unix(), sss.Code).Update("close", kd.Close).Error; err != nil {
				global.GVA_LOG.Error("update k_data err:", zap.Error(err), zap.Any("kd", kd))
			}
			continue
		}
		if err := global.GVA_DB.Create(&kd).Error; err != nil {
			global.GVA_LOG.Error("add err:", zap.Error(err), zap.Any("kd", kd))
			// return
		}
		// 存储每分钟数据
		value := map[string]interface{}{
			"open":  kd.Open,
			"high":  kd.High,
			"low":   kd.Low,
			"close": kd.Close,
		}
		key := now.Format(dateFormat)
		rKey := fmt.Sprintf("k_data_%s_%s", key, sss.Code)
		if err := global.GVA_REDIS.HMSet(ctx, rKey, value).Err(); err != nil {
			global.GVA_LOG.Error("add err:", zap.Error(err), zap.Any("kd", kd))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rKey, time.Hour*24)

		// 处理日k
		// 获取开盘
		// 日k 交易日
		tradeDay := int64(0)
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_day_trading_day_%s_%s", now.Format("2006-01-02"), sss.Code)).Result(); err1 != nil {
			tradeDay = cast.ToInt64(res)
		}
		zeroTime := utils.GetTime(tradeDay)
		if tradeDay == 0 || zeroTime == 0 {
			global.GVA_LOG.Error("k_data_day_trading_day err:", zap.Any("sss.Code", sss.Code), zap.Any("tradeDay", tradeDay))
		}
		var dkd kdata.KData1440
		var isAdd bool
		if err := global.GVA_DB.Where("uptime = ? and symbol_id = ?", zeroTime, sss.Code).First(&dkd).Error; err != nil {
			global.GVA_LOG.Error("add err:", zap.Error(err), zap.Any("zeroTime", zeroTime), zap.Any("sss.Code", sss.Code))
		}
		if dkd.ID == 0 && dkd.Uptime == 0 {
			dkd = kdata.KData1440{
				Uptime: zeroTime,
			}
			isAdd = true
		}
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_day_start_%s_%s", now.Format("2006-01-02"), sss.Code)).Result(); err1 == nil {
			dkd.Open = cast.ToInt64(utils.Decimal(cast.ToFloat64(res)))
		}
		// 获取收盘
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_day_end_%s_%s", now.Format("2006-01-02"), sss.Code)).Result(); err1 == nil {
			dkd.Close = cast.ToInt64(utils.Decimal(cast.ToFloat64(res)))
		}
		// 获取最高
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_day_high_%s_%s", now.Format("2006-01-02"), sss.Code)).Result(); err1 == nil {
			dkd.High = cast.ToInt64(utils.Decimal(cast.ToFloat64(res)))
		}
		// 获取最低
		if res, err1 := global.GVA_REDIS.Get(ctx, fmt.Sprintf("k_data_day_low_%s_%s", now.Format("2006-01-02"), sss.Code)).Result(); err1 == nil {
			dkd.Low = cast.ToInt64(utils.Decimal(cast.ToFloat64(res)))
		}
		dkd.SymbolId = sss.Code
		if dkd.Open == 0 || dkd.Close == 0 || dkd.High == 0 || dkd.Low == 0 {
			global.GVA_LOG.Error("add err:", zap.Any("dkd", dkd))
			continue
		}
		if isAdd { // 添加
			if err := global.GVA_DB.Create(&dkd).Error; err != nil {
				global.GVA_LOG.Error("add err:", zap.Error(err), zap.Any("dkd", dkd))
				// return
			}
		} else { // 更新
			if err := global.GVA_DB.Save(&dkd).Error; err != nil {
				global.GVA_LOG.Error("add err:", zap.Error(err), zap.Any("dkd", dkd))
				// return
			}
		}

		// 存储每分钟数据
		value1 := map[string]interface{}{
			"open":        dkd.Open,
			"high":        dkd.High,
			"low":         dkd.Low,
			"close":       dkd.Close,
			"trading_day": tradeDay,
		}

		rDayKey := fmt.Sprintf("k_data_day_%s_%s", now.Format("2006-01-02"), sss.Code)
		if err := global.GVA_REDIS.HMSet(ctx, rDayKey, value1).Err(); err != nil {
			global.GVA_LOG.Error("add err:", zap.Error(err), zap.Any("kd", kd))
			// return
		}
		global.GVA_REDIS.Expire(ctx, rDayKey, time.Hour*24)
	}

}

func LopKData() {
	// 计算到下一分钟的等待时间
	next := time.Now().Truncate(time.Minute).Add(time.Minute)
	wait := time.Until(next)

	// 等待到下一分钟
	time.Sleep(wait)
	// 创建一个每分钟触发一次的Ticker
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				lopKData()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func lopKData() {
	var ss []symbols.Symbol
	if err := global.GVA_DB.Find(&ss).Error; err != nil {
		global.GVA_LOG.Error("Symbol err", zap.Error(err))
		return
	}
	now := time.Now()
	global.GVA_LOG.Error("lopKData", zap.Any("time", now.Format(time.DateTime)))

	for _, sss := range ss {
		s := sss
		// 5分钟
		if (now.Minute()%10 == 6 && now.Second() == 0) || (now.Minute()%10 == 1 && now.Second() == 0) {
			var kds []kdata.KData
			if err := global.GVA_DB.Where("symbol_id = ? and uptime between ? and ? ", s.Code, now.Add(-6*time.Minute).Unix(), now.Unix()-1).Order("uptime desc").Limit(5).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 5 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd5 := kdata.KData5{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd5.Open == 0 || kd5.Close == 0 || kd5.High == 0 || kd5.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd5).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		// 15分钟(15,30,45,00)
		if (now.Minute() == 16 && now.Second() == 0) || (now.Minute() == 31 && now.Second() == 0) ||
			(now.Minute() == 46 && now.Second() == 0) || (now.Minute() == 1 && now.Second() == 0) {
			var kds []kdata.KData5
			if err := global.GVA_DB.Where("symbol_id = ? and uptime between ? and ?", s.Code, now.Add(-16*time.Minute).Unix(), now.Unix()-1).Order("uptime desc").Limit(3).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 3 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd15 := kdata.KData15{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd15.Open == 0 || kd15.Close == 0 || kd15.High == 0 || kd15.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd15).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		// 30分钟(31,01)
		if utils.IsK30(now) && !utils.IsYe(s.Code) {
			var kds []kdata.KData15
			if err := global.GVA_DB.Where("symbol_id = ? and uptime <= ?", s.Code, now.Unix()).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData30{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		if utils.IsK30Ye(now) && utils.IsYe(s.Code) {
			var kds []kdata.KData15
			if err := global.GVA_DB.Where("symbol_id = ? and uptime <= ?", s.Code, now.Unix()).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData30{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		// 小时线
		if utils.IsK60(now) && !utils.IsYe(s.Code) {
			var kds []kdata.KData30
			if err := global.GVA_DB.Where("symbol_id = ? and uptime <= ?", s.Code, now.Unix()).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData60{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		if utils.IsK60Ye(now) && utils.IsYe(s.Code) {
			var kds []kdata.KData30
			if err := global.GVA_DB.Where("symbol_id = ? and uptime <= ?", s.Code, now.Unix()).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData60{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		// 2小时
		if utils.IsK120(now) && !utils.IsYe(s.Code) {
			var kds []kdata.KData60
			if err := global.GVA_DB.Where("symbol_id = ? and uptime <= ?", s.Code, now.Unix()).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData120{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		if utils.IsK120Ye(now) && utils.IsYe(s.Code) {
			var kds []kdata.KData60
			if err := global.GVA_DB.Where("symbol_id = ? and uptime <= ?", s.Code, now.Unix()).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData120{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		// 4小时
		if utils.IsK240(now) && !utils.IsYe(s.Code) {
			var kds []kdata.KData120
			if err := global.GVA_DB.Where("symbol_id = ? and uptime <= ?", s.Code, now.Unix()).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData240{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		if utils.IsK240Ye(now) && utils.IsYe(s.Code) {
			var kds []kdata.KData120
			if err := global.GVA_DB.Where("symbol_id = ? and uptime <= ?", s.Code, now.Unix()).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData240{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		// 6小时
		if now.Minute() == 1 && now.Second() == 0 && now.Hour()%6 == 0 {
			var kds []kdata.KData120
			if err := global.GVA_DB.Where("symbol_id = ? and uptime between ? and ?", s.Code, now.Add(-361*time.Minute).Unix(), now.Unix()-1).Order("uptime desc").Limit(3).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 3 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData360{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		// 8小时
		if now.Minute() == 1 && now.Second() == 0 && now.Hour()%8 == 0 {
			var kds []kdata.KData240
			if err := global.GVA_DB.Where("symbol_id = ? and uptime between ? and ?", s.Code, now.Add(-481*time.Minute).Unix(), now.Unix()-1).Order("uptime desc").Limit(2).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 2 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData480{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
		// 24小时
		if now.Minute() == 1 && now.Second() == 0 && now.Hour() == 0 {
			var kds []kdata.KData480
			if err := global.GVA_DB.Where("symbol_id = ? and uptime between ? and ?", s.Code, now.Add(-(1*time.Minute + time.Hour*24)).Unix(), now.Unix()-1).Order("uptime desc").Limit(3).Find(&kds).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err))
				continue
			}
			if len(kds) != 3 {
				continue
			}
			var low []int64
			var high []int64
			for _, kd := range kds {
				low = append(low, kd.Low)
				high = append(high, kd.High)
			}
			min := utils.FindMin(low)
			max := utils.FindMax(high)
			kd := kdata.KData1440{
				Uptime:   now.Add(-1 * time.Minute).Unix(),
				Open:     kds[len(kds)-1].Open,
				High:     max,
				Low:      min,
				Close:    kds[0].Close,
				SymbolId: s.Code,
			}
			if kd.Open == 0 || kd.Close == 0 || kd.High == 0 || kd.Low == 0 {
				return
			}
			if err := global.GVA_DB.Create(&kd).Error; err != nil {
				global.GVA_LOG.Error("LopKData:", zap.Error(err), zap.Any("kds", kds))
				// return
			}
		}
	}
}
