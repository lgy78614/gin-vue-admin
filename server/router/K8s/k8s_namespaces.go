package K8s

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sNamespacesRouter struct {}

// InitK8sNamespacesRouter 初始化 k8sNamespaces表 路由信息
func (s *K8sNamespacesRouter) InitK8sNamespacesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	k8sNamespacesRouter := Router.Group("k8sNamespaces").Use(middleware.OperationRecord())
	k8sNamespacesRouterWithoutRecord := Router.Group("k8sNamespaces")
	k8sNamespacesRouterWithoutAuth := PublicRouter.Group("k8sNamespaces")
	{
		k8sNamespacesRouter.POST("createK8sNamespaces", k8sNamespacesApi.CreateK8sNamespaces)   // 新建k8sNamespaces表
		k8sNamespacesRouter.DELETE("deleteK8sNamespaces", k8sNamespacesApi.DeleteK8sNamespaces) // 删除k8sNamespaces表
		k8sNamespacesRouter.DELETE("deleteK8sNamespacesByIds", k8sNamespacesApi.DeleteK8sNamespacesByIds) // 批量删除k8sNamespaces表
		k8sNamespacesRouter.PUT("updateK8sNamespaces", k8sNamespacesApi.UpdateK8sNamespaces)    // 更新k8sNamespaces表
	}
	{
		k8sNamespacesRouterWithoutRecord.GET("findK8sNamespaces", k8sNamespacesApi.FindK8sNamespaces)        // 根据ID获取k8sNamespaces表
		k8sNamespacesRouterWithoutRecord.GET("getK8sNamespacesList", k8sNamespacesApi.GetK8sNamespacesList)  // 获取k8sNamespaces表列表
	}
	{
	    k8sNamespacesRouterWithoutAuth.GET("getK8sNamespacesPublic", k8sNamespacesApi.GetK8sNamespacesPublic)  // k8sNamespaces表开放接口
	}
}
