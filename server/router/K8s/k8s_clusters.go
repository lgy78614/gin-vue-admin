package K8s

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sClustersRouter struct{}

func (s *K8sClustersRouter) InitK8sClustersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	k8sClustersRouter := Router.Group("k8sClusters").Use(middleware.OperationRecord())
	k8sClustersRouterWithoutRecord := Router.Group("k8sClusters")
	k8sClustersRouterWithoutAuth := PublicRouter.Group("k8sClusters")
	{
		k8sClustersRouter.POST("createK8sClusters", k8sClustersApi.CreateK8sClusters)
		k8sClustersRouter.DELETE("deleteK8sClusters", k8sClustersApi.DeleteK8sClusters)
		k8sClustersRouter.DELETE("deleteK8sClustersByIds", k8sClustersApi.DeleteK8sClustersByIds)
		k8sClustersRouter.PUT("updateK8sClusters", k8sClustersApi.UpdateK8sClusters)
	}
	{
		k8sClustersRouterWithoutRecord.GET("findK8sClusters", k8sClustersApi.FindK8sClusters)
		k8sClustersRouterWithoutRecord.GET("getK8sClustersList", k8sClustersApi.GetK8sClustersList)
	}
	{
		k8sClustersRouterWithoutAuth.GET("getK8sClustersPublic", k8sClustersApi.GetK8sClustersPublic)
		k8sClustersRouterWithoutAuth.GET("k8sclusters", k8sClustersApi.GetK8sClusterInfo)
	}
}
