
// 自动生成模板K8sNamespaces
package K8s
import (
	"time"
)

// k8sNamespaces表 结构体  K8sNamespaces
type K8sNamespaces struct {
  NamespaceId  *int64 `json:"namespaceId" form:"namespaceId" gorm:"primarykey;comment:自增主键;column:namespace_id;"`  //自增主键
  ClusterId  *string `json:"clusterId" form:"clusterId" gorm:"comment:所属集群ID;column:cluster_id;size:64;"`  //所属集群ID
  Uid  *string `json:"uid" form:"uid" gorm:"comment:K8s中的唯一UID;column:uid;size:64;"`  //K8s中的唯一UID
  Name  *string `json:"name" form:"name" gorm:"comment:命名空间名称;column:name;size:128;"`  //命名空间名称
  Status  *string `json:"status" form:"status" gorm:"comment:状态: Active/Terminating;column:status;size:32;"`  //状态: Active/Terminating
  Labels  *string `json:"labels" form:"labels" gorm:"comment:标签(JSON格式);column:labels;"`  //标签(JSON格式)
  CreationTimestamp  *time.Time `json:"creationTimestamp" form:"creationTimestamp" gorm:"comment:K8s中的创建时间;column:creation_timestamp;"`  //K8s中的创建时间
  DeletionTimestamp  *time.Time `json:"deletionTimestamp" form:"deletionTimestamp" gorm:"comment:计划删除时间;column:deletion_timestamp;"`  //计划删除时间
}


// TableName k8sNamespaces表 K8sNamespaces自定义表名 k8s_namespaces
func (K8sNamespaces) TableName() string {
    return "k8s_namespaces"
}





