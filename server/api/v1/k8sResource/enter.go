package k8sResource

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ K8sResourceApi }

var krService = service.ServiceGroupApp.K8sResourceServiceGroup.K8sResourceService
