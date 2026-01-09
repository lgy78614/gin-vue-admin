import service from '@/utils/request'
// @Tags K8sWorkloads
// @Summary 创建k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sWorkloads true "创建k8sWorkloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sWorkloads/createK8sWorkloads [post]
export const createK8sWorkloads = (data) => {
  return service({
    url: '/k8sWorkloads/createK8sWorkloads',
    method: 'post',
    data
  })
}

// @Tags K8sWorkloads
// @Summary 删除k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sWorkloads true "删除k8sWorkloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sWorkloads/deleteK8sWorkloads [delete]
export const deleteK8sWorkloads = (params) => {
  return service({
    url: '/k8sWorkloads/deleteK8sWorkloads',
    method: 'delete',
    params
  })
}

// @Tags K8sWorkloads
// @Summary 批量删除k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除k8sWorkloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sWorkloads/deleteK8sWorkloads [delete]
export const deleteK8sWorkloadsByIds = (params) => {
  return service({
    url: '/k8sWorkloads/deleteK8sWorkloadsByIds',
    method: 'delete',
    params
  })
}

// @Tags K8sWorkloads
// @Summary 更新k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sWorkloads true "更新k8sWorkloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sWorkloads/updateK8sWorkloads [put]
export const updateK8sWorkloads = (data) => {
  return service({
    url: '/k8sWorkloads/updateK8sWorkloads',
    method: 'put',
    data
  })
}

// @Tags K8sWorkloads
// @Summary 用id查询k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.K8sWorkloads true "用id查询k8sWorkloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sWorkloads/findK8sWorkloads [get]
export const findK8sWorkloads = (params) => {
  return service({
    url: '/k8sWorkloads/findK8sWorkloads',
    method: 'get',
    params
  })
}

// @Tags K8sWorkloads
// @Summary 分页获取k8sWorkloads表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取k8sWorkloads表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sWorkloads/getK8sWorkloadsList [get]
export const getK8sWorkloadsList = (params) => {
  return service({
    url: '/k8sWorkloads/getK8sWorkloadsList',
    method: 'get',
    params
  })
}

// @Tags K8sWorkloads
// @Summary 不需要鉴权的k8sWorkloads表接口
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sWorkloadsSearch true "分页获取k8sWorkloads表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sWorkloads/getK8sWorkloadsPublic [get]
export const getK8sWorkloadsPublic = () => {
  return service({
    url: '/k8sWorkloads/getK8sWorkloadsPublic',
    method: 'get',
  })
}
