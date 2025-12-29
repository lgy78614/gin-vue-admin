package k8sResource

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sResourceRouter struct {}

// InitK8sResourceRouter 初始化 K8s资源 路由信息
func (s *K8sResourceRouter) InitK8sResourceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	krRouter := Router.Group("kr").Use(middleware.OperationRecord())
	krRouterWithoutRecord := Router.Group("kr")
	krRouterWithoutAuth := PublicRouter.Group("kr")
	{
		krRouter.POST("createK8sResource", krApi.CreateK8sResource)   // 新建K8s资源
		krRouter.DELETE("deleteK8sResource", krApi.DeleteK8sResource) // 删除K8s资源
		krRouter.DELETE("deleteK8sResourceByIds", krApi.DeleteK8sResourceByIds) // 批量删除K8s资源
		krRouter.PUT("updateK8sResource", krApi.UpdateK8sResource)    // 更新K8s资源
	}
	{
		krRouterWithoutRecord.GET("findK8sResource", krApi.FindK8sResource)        // 根据ID获取K8s资源
		krRouterWithoutRecord.GET("getK8sResourceList", krApi.GetK8sResourceList)  // 获取K8s资源列表
	}
	{
	    krRouterWithoutAuth.GET("getK8sResourcePublic", krApi.GetK8sResourcePublic)  // K8s资源开放接口
	}
}
