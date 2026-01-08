
// 自动生成模板K8sNodes
package K8s
import (
	"time"
)

// k8sNodes表 结构体  K8sNodes
type K8sNodes struct {
  NodeId  *string `json:"nodeId" form:"nodeId" gorm:"primarykey;comment:节点唯一标识;column:node_id;size:64;"`  //节点唯一标识
  ClusterId  *string `json:"clusterId" form:"clusterId" gorm:"comment:所属集群ID;column:cluster_id;size:64;"`  //所属集群ID
  Name  *string `json:"name" form:"name" gorm:"comment:节点名称;column:name;size:256;"`  //节点名称
  Status  *string `json:"status" form:"status" gorm:"comment:状态: Ready/NotReady/Unknown;column:status;size:32;"`  //状态: Ready/NotReady/Unknown
  Role  *string `json:"role" form:"role" gorm:"comment:角色: control-plane/worker;column:role;size:32;"`  //角色: control-plane/worker
  OsImage  *string `json:"osImage" form:"osImage" gorm:"comment:操作系统镜像;column:os_image;size:128;"`  //操作系统镜像
  KernelVersion  *string `json:"kernelVersion" form:"kernelVersion" gorm:"comment:内核版本;column:kernel_version;size:128;"`  //内核版本
  Architecture  *string `json:"architecture" form:"architecture" gorm:"comment:架构: amd64/arm64;column:architecture;size:16;"`  //架构: amd64/arm64
  CpuCapacity  *float64 `json:"cpuCapacity" form:"cpuCapacity" gorm:"comment:CPU核心数(可分配总量);column:cpu_capacity;size:10;"`  //CPU核心数(可分配总量)
  MemoryCapacityBytes  *int64 `json:"memoryCapacityBytes" form:"memoryCapacityBytes" gorm:"comment:内存容量(字节);column:memory_capacity_bytes;"`  //内存容量(字节)
  PodCapacity  *int32 `json:"podCapacity" form:"podCapacity" gorm:"comment:Pod容量上限;column:pod_capacity;"`  //Pod容量上限
  CreationTimestamp  *time.Time `json:"creationTimestamp" form:"creationTimestamp" gorm:"comment:K8s中的创建时间;column:creation_timestamp;"`  //K8s中的创建时间
  LastHeartbeat  *time.Time `json:"lastHeartbeat" form:"lastHeartbeat" gorm:"comment:最后一次心跳/同步时间;column:last_heartbeat;"`  //最后一次心跳/同步时间
}


// TableName k8sNodes表 K8sNodes自定义表名 k8s_nodes
func (K8sNodes) TableName() string {
    return "k8s_nodes"
}





