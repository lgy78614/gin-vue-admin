package K8s

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sClustersRouter struct {}

// InitK8sClustersRouter 初始化 k8sClusters表 路由信息
func (s *K8sClustersRouter) InitK8sClustersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	k8sClustersRouter := Router.Group("k8sClusters").Use(middleware.OperationRecord())
	k8sClustersRouterWithoutRecord := Router.Group("k8sClusters")
	k8sClustersRouterWithoutAuth := PublicRouter.Group("k8sClusters")
	{
		k8sClustersRouter.POST("createK8sClusters", k8sClustersApi.CreateK8sClusters)   // 新建k8sClusters表
		k8sClustersRouter.DELETE("deleteK8sClusters", k8sClustersApi.DeleteK8sClusters) // 删除k8sClusters表
		k8sClustersRouter.DELETE("deleteK8sClustersByIds", k8sClustersApi.DeleteK8sClustersByIds) // 批量删除k8sClusters表
		k8sClustersRouter.PUT("updateK8sClusters", k8sClustersApi.UpdateK8sClusters)    // 更新k8sClusters表
	}
	{
		k8sClustersRouterWithoutRecord.GET("findK8sClusters", k8sClustersApi.FindK8sClusters)        // 根据ID获取k8sClusters表
		k8sClustersRouterWithoutRecord.GET("getK8sClustersList", k8sClustersApi.GetK8sClustersList)  // 获取k8sClusters表列表
	}
	{
	    k8sClustersRouterWithoutAuth.GET("getK8sClustersPublic", k8sClustersApi.GetK8sClustersPublic)  // k8sClusters表开放接口
	}
}
