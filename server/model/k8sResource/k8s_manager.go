
// 自动生成模板K8sResource
package k8sResource
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// K8s资源 结构体  K8sResource
type K8sResource struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:K8s资源名称;column:name;" binding:"required"`  //资源名称
  Namespace  *string `json:"namespace" form:"namespace" gorm:"comment:K8s资源所属命名空间;column:namespace;" binding:"required"`  //命名空间
  Type  string `json:"type" form:"type" gorm:"comment:K8s资源类型;column:type;type:enum('pod','deployment','service','configmap','secret','ingress','statefulset','daemonset','job','cronjob');" binding:"required"`  //资源类型
  Labels  datatypes.JSON `json:"labels" form:"labels" gorm:"comment:K8s资源标签;column:labels;" swaggertype:"object"`  //资源标签
  Annotations  datatypes.JSON `json:"annotations" form:"annotations" gorm:"comment:K8s资源注解;column:annotations;" swaggertype:"object"`  //资源注解
  Status  *string `json:"status" form:"status" gorm:"comment:K8s资源状态;column:status;"`  //资源状态
  Spec  datatypes.JSON `json:"spec" form:"spec" gorm:"comment:K8s资源规格;column:spec;" swaggertype:"object"`  //资源规格
  Config  datatypes.JSON `json:"config" form:"config" gorm:"comment:K8s资源配置;column:config;" swaggertype:"object"`  //资源配置
  LastModified  *time.Time `json:"lastModified" form:"lastModified" gorm:"comment:K8s资源最后修改时间;column:last_modified;"`  //最后修改
}


// TableName K8s资源 K8sResource自定义表名 k8s_resources
func (K8sResource) TableName() string {
    return "k8s_resources"
}





