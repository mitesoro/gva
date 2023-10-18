package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/api"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/users"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Users   users.RouterGroup
	Api     api.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
