package handel

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orders"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/flipped-aurora/gin-vue-admin/server/pb"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"time"
)

// HandelOrders 处理平仓
func HandelOrders(d data.Data) {
	now := time.Now()
	if now.Hour() < 9 {
		return
	}
	if now.Hour() == 10 && now.Minute() > 15 && now.Minute() < 30 {
		return
	}

	if now.Hour() == 11 && now.Minute() > 30 {
		return
	}
	if now.Hour() == 12 {
		return
	}
	if now.Hour() == 13 && now.Minute() < 30 {
		return
	}
	if now.Hour() == 15 && now.Minute() > 0 {
		return
	}
	var list []orders.Orders
	if err := global.GVA_DB.Model(orders.Orders{}).Where("status = 1").Order("id desc").Find(&list).Error; err != nil {
		global.GVA_LOG.Error("handelOrders err", zap.Error(err))
		return
	}
	var ss []symbols.Symbol
	if err := global.GVA_DB.Model(symbols.Symbol{}).Find(&ss).Error; err != nil {
		global.GVA_LOG.Error("handelOrders err", zap.Error(err))
		return
	}
	ms := make(map[string]symbols.Symbol)
	for _, s := range ss {
		ms[s.Code] = s
	}
	price := d.LastPrice //  最新价
	status := 5          // 平仓
	for _, order := range list {
		s, ok := ms[order.SymbolID]
		if !ok {
			continue
		}
		if d.SymbolId != order.SymbolID {
			continue
		}
		key := fmt.Sprintf("lock_order_%d_%s", order.ID, order.SymbolID)
		firstLock := utils.NewRedisLock(global.GVA_REDIS, key)
		firstLock.SetExpire(5)
		againAcquire, err := firstLock.Acquire(context.Background())
		if err != nil {
			global.GVA_LOG.Error("NewRedisLock Acquire", zap.Error(err))
			continue
		}
		if !againAcquire {
			continue
		}
		var isComplete bool
		var u users.Users
		var sPrice float64
		if err = global.GVA_DB.Where("id = ?", order.User_id).First(&u).Error; err != nil {
			global.GVA_LOG.Error("find user err", zap.Error(err), zap.Any("order", order))
			continue
		}
		logType := 4
		if *order.Direction == 1 { // 买多
			if float64(*order.Price+*s.PointSuccess) <= price { // 止赢
				order.Status = &status
				order.CompleteAt = model.LocalTime(time.Now())
				order.IsWin = 1
				order.WinAmount = int64(*s.PointSuccessPrice) * 100 // 赢的金额
				isComplete = true
				sPrice = float64(*order.Price + *s.PointSuccess)
			}
			if cast.ToFloat64(*order.Price)-cast.ToFloat64(*s.PointFail) > price { // 止损
				order.Status = &status
				order.CompleteAt = model.LocalTime(time.Now())
				order.IsWin = 2
				order.WinAmount = int64(*s.PointFailPrice) * 100 // 赢的金额
				isComplete = true
				logType = 5
				sPrice = cast.ToFloat64(*order.Price) - cast.ToFloat64(*s.PointFail)
			}
		}
		if *order.Direction == 2 { // 卖空
			if float64(*order.Price+*s.PointFail) < price { // 止损
				order.Status = &status
				order.CompleteAt = model.LocalTime(time.Now())
				order.IsWin = 2
				order.WinAmount = int64(*s.PointFailPrice) * 100 // 赢的金额
				isComplete = true
				logType = 5
				sPrice = cast.ToFloat64(*order.Price) - cast.ToFloat64(*s.PointFail)
			}
			if cast.ToFloat64(*order.Price)-cast.ToFloat64(*s.PointSuccess) >= price { // 止赢
				order.Status = &status
				order.CompleteAt = model.LocalTime(time.Now())
				order.IsWin = 1
				order.WinAmount = int64(*s.PointSuccessPrice) * 100 // 赢的金额
				isComplete = true
				sPrice = float64(*order.Price + *s.PointSuccess)
			}
		}
		if isComplete { // 平仓
			order.ClosePrice = int(price)
			global.GVA_LOG.Error("平仓", zap.Any("price", price), zap.Any("d_price", d.LastPrice),
				zap.Any("order", order), zap.Any("d", d))
			global.GVA_LOG.Error("SymbolID", zap.Any("s", s), zap.Any("order.SymbolID", order.SymbolID), zap.Any("ms", ms))
			if err := global.GVA_DB.Save(&order).Error; err != nil {
				global.GVA_LOG.Error("save order err", zap.Error(err), zap.Any("order", order))
				continue
			}
			// 添加金币
			amount := int(order.DecrAmount) + int(order.WinAmount)
			u.AvailableAmount += amount
			u.Amount += int(order.WinAmount)
			u.FreezeAmount -= amount
			if u.FreezeAmount < 0 {
				u.FreezeAmount = 0
			}
			if err = global.GVA_DB.Save(&u).Error; err != nil {
				global.GVA_LOG.Error("save user err", zap.Error(err), zap.Any("order", order))
				continue
			}
			if int(order.WinAmount) > 0 {
				utils.AddAmountLog(int(u.ID), int(order.WinAmount), u.AvailableAmount, logType)
			}
			utils.AddAmountLog(int(u.ID), int(order.DecrAmount), u.AvailableAmount+int(order.WinAmount), 6)
			reqClient := &pb.OrderRequest{
				C:       order.SymbolID,
				V:       1,
				Close:   true,
				OrderId: int32(order.ID),
				P:       float32(sPrice),
			}
			if *order.Direction == 1 {
				reqClient.Buy = true
				if u.OrderType == 2 {
					reqClient.Sell = true
				}
			}
			if *order.Direction == 2 {
				reqClient.Sell = true
				if u.OrderType == 2 {
					reqClient.Buy = true
				}
			}
			res, err := global.GVA_GrpcCLient.Order(context.Background(), reqClient)
			global.GVA_LOG.Info("grp order", zap.Any("res", res), zap.Error(err), zap.Any("reqClient", reqClient))
			ssss := "买入"
			sss := "盈利"
			if *order.Direction == 2 {
				ssss = "卖出"
			}
			if order.IsWin == 2 {
				sss = "亏损"
			}
			utils.AddMessage(int64(*order.User_id),
				fmt.Sprintf("【平仓通知】您%s的一手产品名已平仓，成交价：%d，平仓价：%d，%s:%d元",
					ssss, order.Price, order.ClosePrice, sss, order.WinAmount/100))
		}
		firstLock.Release(context.Background())
	}
}
