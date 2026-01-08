package K8s

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sPodsRouter struct {}

// InitK8sPodsRouter 初始化 k8sPods表 路由信息
func (s *K8sPodsRouter) InitK8sPodsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	k8sPodsRouter := Router.Group("k8sPods").Use(middleware.OperationRecord())
	k8sPodsRouterWithoutRecord := Router.Group("k8sPods")
	k8sPodsRouterWithoutAuth := PublicRouter.Group("k8sPods")
	{
		k8sPodsRouter.POST("createK8sPods", k8sPodsApi.CreateK8sPods)   // 新建k8sPods表
		k8sPodsRouter.DELETE("deleteK8sPods", k8sPodsApi.DeleteK8sPods) // 删除k8sPods表
		k8sPodsRouter.DELETE("deleteK8sPodsByIds", k8sPodsApi.DeleteK8sPodsByIds) // 批量删除k8sPods表
		k8sPodsRouter.PUT("updateK8sPods", k8sPodsApi.UpdateK8sPods)    // 更新k8sPods表
	}
	{
		k8sPodsRouterWithoutRecord.GET("findK8sPods", k8sPodsApi.FindK8sPods)        // 根据ID获取k8sPods表
		k8sPodsRouterWithoutRecord.GET("getK8sPodsList", k8sPodsApi.GetK8sPodsList)  // 获取k8sPods表列表
	}
	{
	    k8sPodsRouterWithoutAuth.GET("getK8sPodsPublic", k8sPodsApi.GetK8sPodsPublic)  // k8sPods表开放接口
	}
}
