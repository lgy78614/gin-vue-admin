package K8s

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	K8sClustersApi
	K8sNodesApi
	K8sPodsApi
	K8sServicesApi
}

var (
	k8sClustersService = service.ServiceGroupApp.K8sServiceGroup.K8sClustersService
	k8sNodesService    = service.ServiceGroupApp.K8sServiceGroup.K8sNodesService
	k8sPodsService     = service.ServiceGroupApp.K8sServiceGroup.K8sPodsService
	k8sServicesService = service.ServiceGroupApp.K8sServiceGroup.K8sServicesService
)
