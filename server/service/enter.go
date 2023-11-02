package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/apis"
	"github.com/flipped-aurora/gin-vue-admin/server/service/article"
	"github.com/flipped-aurora/gin-vue-admin/server/service/article_category"
	"github.com/flipped-aurora/gin-vue-admin/server/service/configs"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/orders"
	"github.com/flipped-aurora/gin-vue-admin/server/service/symbols"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/users"
)

type ServiceGroup struct {
	SystemServiceGroup           system.ServiceGroup
	ExampleServiceGroup          example.ServiceGroup
	UsersServiceGroup            users.ServiceGroup
	ApisServiceGroup             apis.ServiceGroup
	Article_categoryServiceGroup article_category.ServiceGroup
	ArticleServiceGroup          article.ServiceGroup
	OrdersServiceGroup           orders.ServiceGroup
	ConfigsServiceGroup          configs.ServiceGroup
	SymbolsServiceGroup          symbols.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
