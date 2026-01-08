package K8s

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	K8sClustersRouter
	K8sNodesRouter
}

var (
	k8sClustersApi = api.ApiGroupApp.K8sApiGroup.K8sClustersApi
	k8sNodesApi    = api.ApiGroupApp.K8sApiGroup.K8sNodesApi
)
