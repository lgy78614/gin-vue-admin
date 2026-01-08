import service from '@/utils/request'
// @Tags K8sServices
// @Summary 创建k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sServices true "创建k8sServices表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sServices/createK8sServices [post]
export const createK8sServices = (data) => {
  return service({
    url: '/k8sServices/createK8sServices',
    method: 'post',
    data
  })
}

// @Tags K8sServices
// @Summary 删除k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sServices true "删除k8sServices表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sServices/deleteK8sServices [delete]
export const deleteK8sServices = (params) => {
  return service({
    url: '/k8sServices/deleteK8sServices',
    method: 'delete',
    params
  })
}

// @Tags K8sServices
// @Summary 批量删除k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除k8sServices表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sServices/deleteK8sServices [delete]
export const deleteK8sServicesByIds = (params) => {
  return service({
    url: '/k8sServices/deleteK8sServicesByIds',
    method: 'delete',
    params
  })
}

// @Tags K8sServices
// @Summary 更新k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sServices true "更新k8sServices表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sServices/updateK8sServices [put]
export const updateK8sServices = (data) => {
  return service({
    url: '/k8sServices/updateK8sServices',
    method: 'put',
    data
  })
}

// @Tags K8sServices
// @Summary 用id查询k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.K8sServices true "用id查询k8sServices表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sServices/findK8sServices [get]
export const findK8sServices = (params) => {
  return service({
    url: '/k8sServices/findK8sServices',
    method: 'get',
    params
  })
}

// @Tags K8sServices
// @Summary 分页获取k8sServices表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取k8sServices表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sServices/getK8sServicesList [get]
export const getK8sServicesList = (params) => {
  return service({
    url: '/k8sServices/getK8sServicesList',
    method: 'get',
    params
  })
}

// @Tags K8sServices
// @Summary 不需要鉴权的k8sServices表接口
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sServicesSearch true "分页获取k8sServices表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sServices/getK8sServicesPublic [get]
export const getK8sServicesPublic = () => {
  return service({
    url: '/k8sServices/getK8sServicesPublic',
    method: 'get',
  })
}
