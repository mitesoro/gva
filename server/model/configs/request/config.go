package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/configs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ConfigSearch struct{
    configs.Config
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
