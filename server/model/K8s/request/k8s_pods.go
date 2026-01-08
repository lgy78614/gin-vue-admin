
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type K8sPodsSearch struct{
      ClusterId  *string `json:"clusterId" form:"clusterId"` 
      Namespace  *string `json:"namespace" form:"namespace"` 
      Name  *string `json:"name" form:"name"` 
    request.PageInfo
}
