package K8s

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sConfigStorageRouter struct {}

// InitK8sConfigStorageRouter 初始化 k8sConfigStorage表 路由信息
func (s *K8sConfigStorageRouter) InitK8sConfigStorageRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	k8sConfigStorageRouter := Router.Group("k8sConfigStorage").Use(middleware.OperationRecord())
	k8sConfigStorageRouterWithoutRecord := Router.Group("k8sConfigStorage")
	k8sConfigStorageRouterWithoutAuth := PublicRouter.Group("k8sConfigStorage")
	{
		k8sConfigStorageRouter.POST("createK8sConfigStorage", k8sConfigStorageApi.CreateK8sConfigStorage)   // 新建k8sConfigStorage表
		k8sConfigStorageRouter.DELETE("deleteK8sConfigStorage", k8sConfigStorageApi.DeleteK8sConfigStorage) // 删除k8sConfigStorage表
		k8sConfigStorageRouter.DELETE("deleteK8sConfigStorageByIds", k8sConfigStorageApi.DeleteK8sConfigStorageByIds) // 批量删除k8sConfigStorage表
		k8sConfigStorageRouter.PUT("updateK8sConfigStorage", k8sConfigStorageApi.UpdateK8sConfigStorage)    // 更新k8sConfigStorage表
	}
	{
		k8sConfigStorageRouterWithoutRecord.GET("findK8sConfigStorage", k8sConfigStorageApi.FindK8sConfigStorage)        // 根据ID获取k8sConfigStorage表
		k8sConfigStorageRouterWithoutRecord.GET("getK8sConfigStorageList", k8sConfigStorageApi.GetK8sConfigStorageList)  // 获取k8sConfigStorage表列表
	}
	{
	    k8sConfigStorageRouterWithoutAuth.GET("getK8sConfigStoragePublic", k8sConfigStorageApi.GetK8sConfigStoragePublic)  // k8sConfigStorage表开放接口
	}
}
