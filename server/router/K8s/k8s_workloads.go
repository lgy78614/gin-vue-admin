package K8s

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sWorkloadsRouter struct {}

// InitK8sWorkloadsRouter 初始化 k8sWorkloads表 路由信息
func (s *K8sWorkloadsRouter) InitK8sWorkloadsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	k8sWorkloadsRouter := Router.Group("k8sWorkloads").Use(middleware.OperationRecord())
	k8sWorkloadsRouterWithoutRecord := Router.Group("k8sWorkloads")
	k8sWorkloadsRouterWithoutAuth := PublicRouter.Group("k8sWorkloads")
	{
		k8sWorkloadsRouter.POST("createK8sWorkloads", k8sWorkloadsApi.CreateK8sWorkloads)   // 新建k8sWorkloads表
		k8sWorkloadsRouter.DELETE("deleteK8sWorkloads", k8sWorkloadsApi.DeleteK8sWorkloads) // 删除k8sWorkloads表
		k8sWorkloadsRouter.DELETE("deleteK8sWorkloadsByIds", k8sWorkloadsApi.DeleteK8sWorkloadsByIds) // 批量删除k8sWorkloads表
		k8sWorkloadsRouter.PUT("updateK8sWorkloads", k8sWorkloadsApi.UpdateK8sWorkloads)    // 更新k8sWorkloads表
	}
	{
		k8sWorkloadsRouterWithoutRecord.GET("findK8sWorkloads", k8sWorkloadsApi.FindK8sWorkloads)        // 根据ID获取k8sWorkloads表
		k8sWorkloadsRouterWithoutRecord.GET("getK8sWorkloadsList", k8sWorkloadsApi.GetK8sWorkloadsList)  // 获取k8sWorkloads表列表
	}
	{
	    k8sWorkloadsRouterWithoutAuth.GET("getK8sWorkloadsPublic", k8sWorkloadsApi.GetK8sWorkloadsPublic)  // k8sWorkloads表开放接口
	}
}
