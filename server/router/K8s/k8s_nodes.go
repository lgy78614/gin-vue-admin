package K8s

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sNodesRouter struct {}

// InitK8sNodesRouter 初始化 k8sNodes表 路由信息
func (s *K8sNodesRouter) InitK8sNodesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	k8sNodesRouter := Router.Group("k8sNodes").Use(middleware.OperationRecord())
	k8sNodesRouterWithoutRecord := Router.Group("k8sNodes")
	k8sNodesRouterWithoutAuth := PublicRouter.Group("k8sNodes")
	{
		k8sNodesRouter.POST("createK8sNodes", k8sNodesApi.CreateK8sNodes)   // 新建k8sNodes表
		k8sNodesRouter.DELETE("deleteK8sNodes", k8sNodesApi.DeleteK8sNodes) // 删除k8sNodes表
		k8sNodesRouter.DELETE("deleteK8sNodesByIds", k8sNodesApi.DeleteK8sNodesByIds) // 批量删除k8sNodes表
		k8sNodesRouter.PUT("updateK8sNodes", k8sNodesApi.UpdateK8sNodes)    // 更新k8sNodes表
	}
	{
		k8sNodesRouterWithoutRecord.GET("findK8sNodes", k8sNodesApi.FindK8sNodes)        // 根据ID获取k8sNodes表
		k8sNodesRouterWithoutRecord.GET("getK8sNodesList", k8sNodesApi.GetK8sNodesList)  // 获取k8sNodes表列表
	}
	{
	    k8sNodesRouterWithoutAuth.GET("getK8sNodesPublic", k8sNodesApi.GetK8sNodesPublic)  // k8sNodes表开放接口
	}
}
