// file: server/model/K8s/response/k8s_clusters.go
package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/K8s"

type K8sClusterWithStatus struct {
	K8s.K8sClusters
	ServerVersion   string `json:"server_version"`
	NodeCount       int    `json:"node_count"`
	ActiveNodeCount int    `json:"active_node_count"`
	PodCount        int    `json:"pod_count"`
	HealthStatus    int    `json:"health_status"` // 1:健康 2:警告 3:异常
}
