
// 自动生成模板K8sWorkloads
package K8s
import (
	"time"
)

// k8sWorkloads表 结构体  K8sWorkloads
type K8sWorkloads struct {
  WorkloadId  *int64 `json:"workloadId" form:"workloadId" gorm:"primarykey;comment:自增主键;column:workload_id;"`  //自增主键
  ClusterId  *string `json:"clusterId" form:"clusterId" gorm:"comment:所属集群ID;column:cluster_id;size:64;"`  //所属集群ID
  Namespace  *string `json:"namespace" form:"namespace" gorm:"comment:所属命名空间;column:namespace;size:128;"`  //所属命名空间
  Uid  *string `json:"uid" form:"uid" gorm:"comment:K8s UID;column:uid;size:64;"`  //K8s UID
  Name  *string `json:"name" form:"name" gorm:"comment:工作负载名称;column:name;size:256;"`  //工作负载名称
  Type  *string `json:"type" form:"type" gorm:"comment:类型: Deployment/StatefulSet/DaemonSet/Job/CronJob;column:type;size:32;"`  //类型: Deployment/StatefulSet/DaemonSet/Job/CronJob
  ReplicasDesired  *int32 `json:"replicasDesired" form:"replicasDesired" gorm:"comment:期望副本数;column:replicas_desired;"`  //期望副本数
  ReplicasCurrent  *int32 `json:"replicasCurrent" form:"replicasCurrent" gorm:"comment:当前副本数;column:replicas_current;"`  //当前副本数
  ReplicasReady  *int32 `json:"replicasReady" form:"replicasReady" gorm:"comment:就绪副本数;column:replicas_ready;"`  //就绪副本数
  Labels  *string `json:"labels" form:"labels" gorm:"comment:标签;column:labels;"`  //标签
  Annotations  *string `json:"annotations" form:"annotations" gorm:"comment:注解;column:annotations;"`  //注解
  Selector  *string `json:"selector" form:"selector" gorm:"comment:标签选择器(JSON格式);column:selector;"`  //标签选择器(JSON格式)
  CreationTimestamp  *time.Time `json:"creationTimestamp" form:"creationTimestamp" gorm:"comment:创建时间;column:creation_timestamp;"`  //创建时间
  LastSyncTimestamp  *time.Time `json:"lastSyncTimestamp" form:"lastSyncTimestamp" gorm:"comment:最后同步时间;column:last_sync_timestamp;"`  //最后同步时间
  StatusConditions  *string `json:"statusConditions" form:"statusConditions" gorm:"comment:状态条件(JSON数组);column:status_conditions;"`  //状态条件(JSON数组)
}


// TableName k8sWorkloads表 K8sWorkloads自定义表名 k8s_workloads
func (K8sWorkloads) TableName() string {
    return "k8s_workloads"
}





