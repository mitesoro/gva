package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/alog"
	"github.com/flipped-aurora/gin-vue-admin/server/router/api"
	"github.com/flipped-aurora/gin-vue-admin/server/router/article"
	"github.com/flipped-aurora/gin-vue-admin/server/router/article_category"
	"github.com/flipped-aurora/gin-vue-admin/server/router/configs"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/orders"
	"github.com/flipped-aurora/gin-vue-admin/server/router/recharge"
	"github.com/flipped-aurora/gin-vue-admin/server/router/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/users"
)

type RouterGroup struct {
	System           system.RouterGroup
	Example          example.RouterGroup
	Users            users.RouterGroup
	Api              api.RouterGroup
	Article_category article_category.RouterGroup
	Article          article.RouterGroup
	Orders           orders.RouterGroup
	Configs          configs.RouterGroup
	Symbols          symbols.RouterGroup
	Recharge         recharge.RouterGroup
	Alog             alog.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
