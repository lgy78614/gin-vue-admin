
// 自动生成模板K8sConfigStorage
package K8s
import (
	"time"
)

// k8sConfigStorage表 结构体  K8sConfigStorage
type K8sConfigStorage struct {
  ConfigId  *int64 `json:"configId" form:"configId" gorm:"primarykey;comment:自增主键;column:config_id;"`  //自增主键
  ClusterId  *string `json:"clusterId" form:"clusterId" gorm:"comment:所属集群ID;column:cluster_id;size:64;"`  //所属集群ID
  Namespace  *string `json:"namespace" form:"namespace" gorm:"comment:所属命名空间;column:namespace;size:128;"`  //所属命名空间
  Uid  *string `json:"uid" form:"uid" gorm:"comment:K8s UID;column:uid;size:64;"`  //K8s UID
  Name  *string `json:"name" form:"name" gorm:"comment:资源名称;column:name;size:256;"`  //资源名称
  Type  *string `json:"type" form:"type" gorm:"comment:类型: ConfigMap/Secret/PersistentVolumeClaim;column:type;size:32;"`  //类型: ConfigMap/Secret/PersistentVolumeClaim
  Data  *string `json:"data" form:"data" gorm:"comment:配置数据或Secret的键值对;column:data;"`  //配置数据或Secret的键值对
  VolumeName  *string `json:"volumeName" form:"volumeName" gorm:"comment:PVC特有: 关联的Volume名称;column:volume_name;size:256;"`  //PVC特有: 关联的Volume名称
  StorageClass  *string `json:"storageClass" form:"storageClass" gorm:"comment:PVC特有: 存储类;column:storage_class;size:128;"`  //PVC特有: 存储类
  AccessModes  *string `json:"accessModes" form:"accessModes" gorm:"comment:PVC特有: 访问模式;column:access_modes;"`  //PVC特有: 访问模式
  CreationTimestamp  *time.Time `json:"creationTimestamp" form:"creationTimestamp" gorm:"comment:创建时间;column:creation_timestamp;"`  //创建时间
}


// TableName k8sConfigStorage表 K8sConfigStorage自定义表名 k8s_config_storage
func (K8sConfigStorage) TableName() string {
    return "k8s_config_storage"
}





