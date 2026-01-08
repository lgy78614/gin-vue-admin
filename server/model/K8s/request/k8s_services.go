
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type K8sServicesSearch struct{
      ServiceId  *int `json:"serviceId" form:"serviceId"` 
      ClusterId  *string `json:"clusterId" form:"clusterId"` 
      Namespace  *string `json:"namespace" form:"namespace"` 
      Name  *string `json:"name" form:"name"` 
    request.PageInfo
}
