package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/K8s"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	K8s     K8s.RouterGroup
}
