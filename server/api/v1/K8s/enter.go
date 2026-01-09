package K8s

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	K8sClustersApi
	K8sNodesApi
	K8sPodsApi
	K8sServicesApi
	K8sNamespacesApi
	K8sWorkloadsApi
	K8sConfigStorageApi
}

var (
	k8sClustersService      = service.ServiceGroupApp.K8sServiceGroup.K8sClustersService
	k8sNodesService         = service.ServiceGroupApp.K8sServiceGroup.K8sNodesService
	k8sPodsService          = service.ServiceGroupApp.K8sServiceGroup.K8sPodsService
	k8sServicesService      = service.ServiceGroupApp.K8sServiceGroup.K8sServicesService
	k8sNamespacesService    = service.ServiceGroupApp.K8sServiceGroup.K8sNamespacesService
	k8sWorkloadsService     = service.ServiceGroupApp.K8sServiceGroup.K8sWorkloadsService
	k8sConfigStorageService = service.ServiceGroupApp.K8sServiceGroup.K8sConfigStorageService
)
