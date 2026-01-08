
// 自动生成模板K8sPods
package K8s
import (
	"time"
)

// k8sPods表 结构体  K8sPods
type K8sPods struct {
  PodId  *int64 `json:"podId" form:"podId" gorm:"primarykey;comment:自增主键;column:pod_id;"`  //自增主键
  ClusterId  *string `json:"clusterId" form:"clusterId" gorm:"comment:所属集群ID;column:cluster_id;size:64;"`  //所属集群ID
  Namespace  *string `json:"namespace" form:"namespace" gorm:"comment:所属命名空间;column:namespace;size:128;"`  //所属命名空间
  Uid  *string `json:"uid" form:"uid" gorm:"comment:K8s UID;column:uid;size:64;"`  //K8s UID
  Name  *string `json:"name" form:"name" gorm:"comment:Pod名称;column:name;size:256;"`  //Pod名称
  NodeName  *string `json:"nodeName" form:"nodeName" gorm:"comment:调度到的节点名称;column:node_name;size:256;"`  //调度到的节点名称
  WorkloadUid  *string `json:"workloadUid" form:"workloadUid" gorm:"comment:所属工作负载的UID;column:workload_uid;size:64;"`  //所属工作负载的UID
  Phase  *string `json:"phase" form:"phase" gorm:"comment:阶段: Pending/Running/Succeeded/Failed/Unknown;column:phase;size:32;"`  //阶段: Pending/Running/Succeeded/Failed/Unknown
  PodIp  *string `json:"podIp" form:"podIp" gorm:"comment:Pod IP地址;column:pod_ip;size:64;"`  //Pod IP地址
  HostIp  *string `json:"hostIp" form:"hostIp" gorm:"comment:主机IP地址;column:host_ip;size:64;"`  //主机IP地址
  Labels  *string `json:"labels" form:"labels" gorm:"comment:标签;column:labels;"`  //标签
  CreationTimestamp  *time.Time `json:"creationTimestamp" form:"creationTimestamp" gorm:"comment:创建时间;column:creation_timestamp;"`  //创建时间
  StartTime  *time.Time `json:"startTime" form:"startTime" gorm:"comment:启动时间(进入Running阶段);column:start_time;"`  //启动时间(进入Running阶段)
  RestartCount  *int32 `json:"restartCount" form:"restartCount" gorm:"comment:重启次数;column:restart_count;"`  //重启次数
  StatusConditions  *string `json:"statusConditions" form:"statusConditions" gorm:"comment:状态条件;column:status_conditions;"`  //状态条件
  LastTerminatedState  *string `json:"lastTerminatedState" form:"lastTerminatedState" gorm:"comment:上一次终止的容器状态;column:last_terminated_state;"`  //上一次终止的容器状态
}


// TableName k8sPods表 K8sPods自定义表名 k8s_pods
func (K8sPods) TableName() string {
    return "k8s_pods"
}





