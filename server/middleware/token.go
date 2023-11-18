package middleware

import (
	"encoding/base64"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"time"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("auth-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录，请去登录", c)
			c.Abort()
			return
		}
		tokenBase, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录，请去登录", c)
			c.Abort()
			return
		}
		res, err := utils.AESDecodeNormal(string(tokenBase), utils.Sign)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问 aes", c)
			c.Abort()
			return
		}
		ts := cast.ToInt64(res["time"])
		if ts > 0 && ts < time.Now().Unix() {
			response.FailWithDetailed(gin.H{"reload": true}, "已过期，请重新登录", c)
			c.Abort()
			return
		}
		deviceID := c.Request.Header.Get("device_id")
		if deviceID == "" {
			response.FailWithMessageWithCode(10002, "请求错误，设备号错误", c)
			c.Abort()
			return
		}
		if res["device_id"] != deviceID {
			response.FailWithMessageWithCode(10002, "您的账号已在其他设备登录", c)
			c.Abort()
			return
		}
		c.Set("uid", cast.ToInt64(res["uid"]))
		c.Set("phone", cast.ToInt64(res["phone"]))
		c.Next()
	}
}
