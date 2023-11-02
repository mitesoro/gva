package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/alog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AlogSearch struct{
    alog.Alog
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
