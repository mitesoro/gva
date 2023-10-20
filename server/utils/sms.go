package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"math/rand"
)

func GenerateRandomCode(length int) string {
	charSet := "0123456789"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charSet))
		code[i] = charSet[randomIndex]
	}

	return string(code)
}

func SendSms(phoneNumber, verificationCode string) bool {
	// 替换为你的Access Key ID 和 Access Key Secret
	accessKeyID := "LTAI5tSoZK65qM2NFy2hxRUA"
	accessKeySecret := "edkf66axt61xSKzRgVMnRRQNn5vCpe"

	// 创建短信客户端
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return false
	}

	// 替换为你的签名和模板代码
	signName := "实时到"
	templateCode := "SMS_167527379"

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = `{"code":"` + verificationCode + `"}`

	request.PhoneNumbers = phoneNumber

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Println("Failed to send SMS:", err)
		return false
	}
	global.GVA_LOG.Info("发送信验证码", zap.String("code", response.Code), zap.String("message", response.Message))
	return response.Code == "OK"
}
