package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/handel"
	"github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orders"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/pb"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"google.golang.org/grpc"
	"log"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}
	// 初始化grpc

	if global.GVA_GrpcCLient == nil {
		// 创建与gRPC服务器的连接
		conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("无法连接到gRPC服务器: %v", err)
		}
		defer conn.Close()

		global.GVA_GrpcCLient = pb.NewGreeterClient(conn)
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	utils.SafeGO(func() {
		if global.GVA_REDIS == nil {
			initialize.Redis()
		}
		sub := global.GVA_REDIS.Subscribe(context.Background(), "channel_name", "channel_notify_order")
		// 获取订阅的通道
		channel := sub.Channel()

		// 在单独的goroutine中处理接收到的消息
		utils.SafeGO(func() {
			for msg := range channel {
				if msg.Channel == "channel_name" {
					handelData(msg.Payload)
				}
				if msg.Channel == "channel_notify_order" {
					handelNotifyOrder(msg.Payload)
				}

			}
		})
	})

	// 缓存品种
	utils.SafeGO(func() {
		if global.GVA_REDIS == nil {
			initialize.Redis()
		}
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				var sbs []symbols.Symbol
				if err := global.GVA_DB.Find(&sbs).Error; err != nil {
					fmt.Println("GVA_DB find Symbol err", err)
					continue
				}
				mm := make(map[string]string)
				for _, sb := range sbs {
					mm[sb.Code] = sb.Code
				}
				if err := global.GVA_REDIS.HMSet(context.Background(), "symbol", mm).Err(); err != nil {
					fmt.Println("GVA_REDIS HMSet Symbol err", err)
				}
			}
		}
	})
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	//	fmt.Printf(`
	//	欢迎使用 gin-vue-admin
	//	当前版本:v2.5.7
	//    加群方式:微信号：shouzi_1994 QQ群：622360840
	//	插件市场:https://plugin.gin-vue-admin.com
	//	GVA讨论社区:https://support.qq.com/products/371961
	//	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	//	默认前端文件运行地址:http://127.0.0.1:8080
	//	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/coffee/index.html
	// `, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

// handelData 处理行情数据
func handelData(msg string) {
	now := time.Now()
	var d data.Data
	err := json.Unmarshal([]byte(msg), &d)
	if err != nil {
		global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
		return
	}
	utils.SafeGO(func() {
		handel.HandelOrders(d)
	})
	d.InsertAt = now.Unix()
	d.PreSettlementPrice = utils.Decimal(d.PreSettlementPrice)
	d.PreClosePrice = utils.Decimal(d.PreClosePrice)
	d.PreOpenInterest = utils.Decimal(d.PreOpenInterest)
	d.UpperLimitPrice = utils.Decimal(d.UpperLimitPrice)
	d.LowerLimitPrice = utils.Decimal(d.LowerLimitPrice)
	d.LastPrice = utils.Decimal(d.LastPrice)
	d.BidPrice = utils.Decimal(d.BidPrice)
	d.AskPrice = utils.Decimal(d.AskPrice)
	d.Turnover = utils.Decimal(d.Turnover)
	d.OpenInterest = utils.Decimal(d.OpenInterest)
	d.AveragePrice = utils.Decimal(d.AveragePrice)

	err = global.GVA_DB.Create(&d).Error
	if err != nil {
		global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
		return
	}
	if now.Minute()%5 == 0 && now.Second() == 0 {
		dd := data.Data5(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	if now.Minute()%15 == 0 && now.Second() == 0 {
		dd := data.Data15(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	if now.Minute()%30 == 0 && now.Second() == 0 {
		dd := data.Data30(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	if now.Minute() == 0 && now.Second() == 0 {
		dd := data.Data60(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	if now.Minute() == 0 && now.Second() == 0 && now.Hour()%2 == 0 {
		dd := data.Data120(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	if now.Minute() == 0 && now.Second() == 0 && now.Hour()%4 == 0 {
		dd := data.Data240(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	if now.Minute() == 0 && now.Second() == 0 && now.Hour()%6 == 0 {
		dd := data.Data360(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	if now.Minute() == 0 && now.Second() == 0 && now.Hour()%8 == 0 {
		dd := data.Data480(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	if now.Minute() == 0 && now.Second() == 0 && now.Hour() == 0 {
		dd := data.Data1440(d)
		err = global.GVA_DB.Create(&dd).Error
		if err != nil {
			global.GVA_LOG.Error("Received message:", zap.Error(err), zap.String("Payload", msg))
			return
		}
	}
	// 缓存最新的数据
	key := fmt.Sprintf("s:info:%s", d.SymbolId)
	if err = global.GVA_REDIS.Set(context.Background(), key, msg, 0).Err(); err != nil {
		global.GVA_LOG.Error("Received message: redis set err ", zap.Error(err), zap.String("Payload", msg))
	}

}

func handelNotifyOrder(msg string) {
	global.GVA_LOG.Info("handelNotifyOrder", zap.String("order", msg))
	var mo MsgOrder
	if err := json.Unmarshal([]byte(msg), &mo); err != nil {
		global.GVA_LOG.Error("handelNotifyOrder Unmarshal", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("handelNotifyOrder struct", zap.Any("mo", mo))
	var o orders.Orders
	err := global.GVA_DB.Where("order_ref = ?", mo.OrderRef).First(&o).Error
	if err != nil {
		global.GVA_LOG.Error("get Order", zap.Error(err), zap.Any("order_ref", mo.OrderRef))
		return
	}
	if mo.OrderSysID <= 0 {
		return
	}
	o.OrderSysID = mo.OrderSysID
	err = global.GVA_DB.Save(&o).Error
	if err != nil {
		global.GVA_LOG.Error("save Order", zap.Error(err))
		return
	}
	// if *o.Status != 0 {
	// 	return
	// }
	// 查询并取消

	// 绑定OrderRef 和 OrderSysID
	res, err1 := global.GVA_GrpcCLient.QueryOrder(context.Background(), &pb.QueryOrderRequest{
		Or: int32(o.OrderSysID),
	})
	if err1 != nil {
		global.GVA_LOG.Error("QueryOrder", zap.Error(err1), zap.Any("OrderSysID", o.OrderSysID))
		return
	}
	global.GVA_LOG.Info("QueryOrder", zap.Any("res", res))
	// 成交
	status := 0
	o.SuccessAt = model.LocalTime(time.Now())
	a := *o.Price
	o.SuccessPrice = int64(a)
	if res.Os == 3 {
		status = 1
	} else {
		status = 2
	}
	o.Status = &status
	err = global.GVA_DB.Save(&o).Error
	if err != nil {
		global.GVA_LOG.Error("save Order", zap.Error(err))
		return
	}

	if status == 2 {
		// 去掉订单
		_, err1 = global.GVA_GrpcCLient.CancelOrder(context.Background(), &pb.QueryOrderRequest{
			Or: int32(mo.OrderSysID),
		})
		if err1 != nil {
			global.GVA_LOG.Error("CancelOrder", zap.Error(err))
		}
		// 返还金钱
		u, _ := utils.GetUser(int64(*o.User_id))
		utils.AddAmountLog(int(u.ID), int(o.DecrAmount), u.AvailableAmount, 3)
		u.AvailableAmount = u.AvailableAmount + int(o.DecrAmount)
		u.FreezeAmount = u.FreezeAmount - int(o.DecrAmount)
		err = global.GVA_DB.Save(&u).Error
		if err != nil {
			global.GVA_LOG.Error("save user", zap.Error(err))
			return
		}
	}
}

type MsgOrder struct {
	ErrorNo        int `json:"error_no"`
	LongOrderSysID int `json:"long_order_sys_id"`
	OrderLocalID   int `json:"order_local_id"`
	OrderRef       int `json:"order_ref"`
	OrderStatus    int `json:"order_status"`
	OrderSysID     int `json:"order_sys_id"`
	TradeVolume    int `json:"trade_volume"`
}
