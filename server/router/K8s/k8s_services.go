package K8s

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sServicesRouter struct {}

// InitK8sServicesRouter 初始化 k8sServices表 路由信息
func (s *K8sServicesRouter) InitK8sServicesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	k8sServicesRouter := Router.Group("k8sServices").Use(middleware.OperationRecord())
	k8sServicesRouterWithoutRecord := Router.Group("k8sServices")
	k8sServicesRouterWithoutAuth := PublicRouter.Group("k8sServices")
	{
		k8sServicesRouter.POST("createK8sServices", k8sServicesApi.CreateK8sServices)   // 新建k8sServices表
		k8sServicesRouter.DELETE("deleteK8sServices", k8sServicesApi.DeleteK8sServices) // 删除k8sServices表
		k8sServicesRouter.DELETE("deleteK8sServicesByIds", k8sServicesApi.DeleteK8sServicesByIds) // 批量删除k8sServices表
		k8sServicesRouter.PUT("updateK8sServices", k8sServicesApi.UpdateK8sServices)    // 更新k8sServices表
	}
	{
		k8sServicesRouterWithoutRecord.GET("findK8sServices", k8sServicesApi.FindK8sServices)        // 根据ID获取k8sServices表
		k8sServicesRouterWithoutRecord.GET("getK8sServicesList", k8sServicesApi.GetK8sServicesList)  // 获取k8sServices表列表
	}
	{
	    k8sServicesRouterWithoutAuth.GET("getK8sServicesPublic", k8sServicesApi.GetK8sServicesPublic)  // k8sServices表开放接口
	}
}
