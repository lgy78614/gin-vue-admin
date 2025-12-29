package k8sResource

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ K8sResourceRouter }

var krApi = api.ApiGroupApp.K8sResourceApiGroup.K8sResourceApi
