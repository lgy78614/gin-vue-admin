
// 自动生成模板K8sServices
package K8s
import (
	"time"
)

// k8sServices表 结构体  K8sServices
type K8sServices struct {
  ServiceId  *int64 `json:"serviceId" form:"serviceId" gorm:"primarykey;comment:自增主键;column:service_id;"`  //自增主键
  ClusterId  *string `json:"clusterId" form:"clusterId" gorm:"comment:所属集群ID;column:cluster_id;size:64;"`  //所属集群ID
  Namespace  *string `json:"namespace" form:"namespace" gorm:"comment:所属命名空间;column:namespace;size:128;"`  //命名空间
  Uid  *string `json:"uid" form:"uid" gorm:"comment:K8s UID;column:uid;size:64;"`  //K8s UID
  Name  *string `json:"name" form:"name" gorm:"comment:服务名称;column:name;size:256;"`  //服务名称
  Type  *string `json:"type" form:"type" gorm:"comment:类型: Service/Ingress;column:type;size:32;"`  //类型
  ServiceType  *string `json:"serviceType" form:"serviceType" gorm:"comment:Service特有: ClusterIP/NodePort/LoadBalancer;column:service_type;size:32;"`  //Service
  ClusterIp  *string `json:"clusterIp" form:"clusterIp" gorm:"comment:Service的Cluster IP;column:cluster_ip;size:64;"`  //Cluster IP
  Ports  *int64 `json:"ports" form:"ports" gorm:"comment:端口配置(JSON数组);column:ports;"`  //端口配置
  Selector  *string `json:"selector" form:"selector" gorm:"comment:后端Pod选择器;column:selector;"`  //Pod选择器
  IngressHosts  *string `json:"ingressHosts" form:"ingressHosts" gorm:"comment:Ingress特有: 关联的Host和Path规则;column:ingress_hosts;"`  //Ingress特有
  CreationTimestamp  *time.Time `json:"creationTimestamp" form:"creationTimestamp" gorm:"comment:创建时间;column:creation_timestamp;"`  //创建时间
}


// TableName k8sServices表 K8sServices自定义表名 k8s_services
func (K8sServices) TableName() string {
    return "k8s_services"
}





