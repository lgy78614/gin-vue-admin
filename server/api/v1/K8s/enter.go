package K8s

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	K8sClustersApi
	K8sNodesApi
}

var (
	k8sClustersService = service.ServiceGroupApp.K8sServiceGroup.K8sClustersService
	k8sNodesService    = service.ServiceGroupApp.K8sServiceGroup.K8sNodesService
)
