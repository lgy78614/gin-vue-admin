// server/model/K8s/response/nodes.go
package response

type K8sNodeDetail struct {
	ClusterId      string   `json:"clusterId"`      // 所属集群ID
	Name           string   `json:"name"`           // 节点名称
	Status         string   `json:"status"`         // Ready/NotReady
	Roles          []string `json:"roles"`          // control-plane, worker 等
	OsImage        string   `json:"osImage"`        // 操作系统镜像
	KernelVersion  string   `json:"kernelVersion"`  // 内核版本
	Architecture   string   `json:"architecture"`   // 架构 (amd64/arm64)
	CpuCapacity    string   `json:"cpuCapacity"`    // CPU 核心数（可分配总量，如 "8"）
	MemoryCapacity string   `json:"memoryCapacity"` // 内存容量（字节，如 "16Gi"）
	PodCapacity    string   `json:"podCapacity"`    // Pod 容量上限（如 "110"）
}
