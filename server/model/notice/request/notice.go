package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/notice"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type NoticeSearch struct{
    notice.Notice
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
