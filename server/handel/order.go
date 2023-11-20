package handel

import (
	"context"
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
	ms := make(map[string]*symbols.Symbol)
	for _, s := range ss {
		ms[s.Code] = &s
	}
	price := d.LastPrice //  最新价
	status := 5          // 平仓
	for _, order := range list {
		s := ms[order.SymbolID]
		if s == nil {
			continue
		}
		var isComplete bool
		var u users.Users
		var sPrice float64
		if err := global.GVA_DB.Where("id = ?", order.User_id).First(&u).Error; err != nil {
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
			if err := global.GVA_DB.Save(&order).Error; err != nil {
				global.GVA_LOG.Error("save order err", zap.Error(err), zap.Any("order", order))
				continue
			}
			// 添加金币
			amount := int(order.DecrAmount) + int(order.WinAmount)
			u.Amount += amount
			u.AvailableAmount += amount
			u.FreezeAmount -= amount
			if err := global.GVA_DB.Save(&u).Error; err != nil {
				global.GVA_LOG.Error("save user err", zap.Error(err), zap.Any("order", order))
				continue
			}
			if int(order.WinAmount) > 0 {
				utils.AddAmountLog(int(u.ID), int(order.WinAmount), u.AvailableAmount, logType)
			}
			utils.AddAmountLog(int(u.ID), *order.Price*100, u.AvailableAmount+int(order.WinAmount), 6)
			reqClient := &pb.OrderRequest{
				C:       order.SymbolID,
				V:       1,
				Close:   true,
				Sell:    true,
				OrderId: int32(order.ID),
				P:       float32(sPrice),
			}
			res, err := global.GVA_GrpcCLient.Order(context.Background(), reqClient)
			global.GVA_LOG.Info("grp order", zap.Any("res", res), zap.Error(err), zap.Any("reqClient", reqClient))
		}
	}
}
