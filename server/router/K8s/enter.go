package K8s

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	K8sClustersRouter
	K8sNodesRouter
	K8sPodsRouter
	K8sServicesRouter
	K8sNamespacesRouter
	K8sWorkloadsRouter
	K8sConfigStorageRouter
}

var (
	k8sClustersApi      = api.ApiGroupApp.K8sApiGroup.K8sClustersApi
	k8sNodesApi         = api.ApiGroupApp.K8sApiGroup.K8sNodesApi
	k8sPodsApi          = api.ApiGroupApp.K8sApiGroup.K8sPodsApi
	k8sServicesApi      = api.ApiGroupApp.K8sApiGroup.K8sServicesApi
	k8sNamespacesApi    = api.ApiGroupApp.K8sApiGroup.K8sNamespacesApi
	k8sWorkloadsApi     = api.ApiGroupApp.K8sApiGroup.K8sWorkloadsApi
	k8sConfigStorageApi = api.ApiGroupApp.K8sApiGroup.K8sConfigStorageApi
)
