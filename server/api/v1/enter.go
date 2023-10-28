package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/apis"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/article"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/article_category"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/orders"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/users"
)

type ApiGroup struct {
	SystemApiGroup           system.ApiGroup
	ExampleApiGroup          example.ApiGroup
	UsersApiGroup            users.ApiGroup
	ApisApiGroup             apis.ApiGroup
	Article_categoryApiGroup article_category.ApiGroup
	ArticleApiGroup          article.ApiGroup
	OrdersApiGroup           orders.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
