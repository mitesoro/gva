package core

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/pb"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: false,
		})
		conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(creds))
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
		sub := global.GVA_REDIS.Subscribe(context.Background(), "channel_name")
		// 获取订阅的通道
		channel := sub.Channel()

		// 在单独的goroutine中处理接收到的消息
		utils.SafeGO(func() {
			for msg := range channel {
				var d data.Data
				err := json.Unmarshal([]byte(msg.Payload), &d)
				if err != nil {
					fmt.Println("Received message:", msg.Payload, err)
					continue
				}
				d.InsertAt = time.Now().Unix()
				// 获取当前时间
				now := time.Now()

				// 判断是否是1分钟K线时间点
				if now.Minute()%1 == 0 && now.Second() == 0 {
					fmt.Println("当前时间是1分钟K线时间点")
				}

				// 判断是否是5分钟K线时间点
				if now.Minute()%5 == 0 && now.Second() == 0 {
					fmt.Println("当前时间是5分钟K线时间点")
				}

				err = global.GVA_DB.Create(&d).Error
				if err != nil {
					fmt.Println(err)
				}
				res, err := global.GVA_GrpcCLient.SayHello(context.Background(), &pb.HelloRequest{Name: "name"})
				if err != nil {
					fmt.Println("err", err)
				}

				fmt.Println(res.GetMessage())
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
