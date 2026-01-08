
// 自动生成模板K8sClusters
package K8s
import (
	"time"
)

// k8sClusters表 结构体  K8sClusters
type K8sClusters struct {
  ClusterId  *string `json:"clusterId" form:"clusterId" gorm:"primarykey;comment:集群唯一标识;column:cluster_id;size:64;"`  //集群唯一标识
  ClusterName  *string `json:"clusterName" form:"clusterName" gorm:"comment:集群显示名称;column:cluster_name;size:128;"`  //集群显示名称
  Environment  *string `json:"environment" form:"environment" gorm:"comment:环境: production/staging/development;column:environment;size:32;"`  //环境: production/staging/development
  K8sVersion  *string `json:"k8sVersion" form:"k8sVersion" gorm:"comment:Kubernetes版本;column:k8s_version;size:32;"`  //Kubernetes版本
  CloudProvider  *string `json:"cloudProvider" form:"cloudProvider" gorm:"comment:云厂商: AWS/Azure/GCP/on-premise;column:cloud_provider;size:64;"`  //云厂商: AWS/Azure/GCP/on-premise
  Region  *string `json:"region" form:"region" gorm:"comment:区域;column:region;size:64;"`  //区域
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:记录创建时间;column:created_at;"`  //记录创建时间
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:记录更新时间;column:updated_at;"`  //记录更新时间
}


// TableName k8sClusters表 K8sClusters自定义表名 k8s_clusters
func (K8sClusters) TableName() string {
    return "k8s_clusters"
}





