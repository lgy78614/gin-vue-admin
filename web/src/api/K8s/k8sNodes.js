import service from '@/utils/request'
// @Tags K8sNodes
// @Summary 创建k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sNodes true "创建k8sNodes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sNodes/createK8sNodes [post]
export const createK8sNodes = (data) => {
  return service({
    url: '/k8sNodes/createK8sNodes',
    method: 'post',
    data
  })
}

// @Tags K8sNodes
// @Summary 删除k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sNodes true "删除k8sNodes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNodes/deleteK8sNodes [delete]
export const deleteK8sNodes = (params) => {
  return service({
    url: '/k8sNodes/deleteK8sNodes',
    method: 'delete',
    params
  })
}

// @Tags K8sNodes
// @Summary 批量删除k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除k8sNodes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNodes/deleteK8sNodes [delete]
export const deleteK8sNodesByIds = (params) => {
  return service({
    url: '/k8sNodes/deleteK8sNodesByIds',
    method: 'delete',
    params
  })
}

// @Tags K8sNodes
// @Summary 更新k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sNodes true "更新k8sNodes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sNodes/updateK8sNodes [put]
export const updateK8sNodes = (data) => {
  return service({
    url: '/k8sNodes/updateK8sNodes',
    method: 'put',
    data
  })
}

// @Tags K8sNodes
// @Summary 用id查询k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.K8sNodes true "用id查询k8sNodes表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sNodes/findK8sNodes [get]
export const findK8sNodes = (params) => {
  return service({
    url: '/k8sNodes/findK8sNodes',
    method: 'get',
    params
  })
}

// @Tags K8sNodes
// @Summary 分页获取k8sNodes表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取k8sNodes表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNodes/getK8sNodesList [get]
export const getK8sNodesList = (params) => {
  return service({
    url: '/k8sNodes/getK8sNodesList',
    method: 'get',
    params
  })
}

// @Tags K8sNodes
// @Summary 不需要鉴权的k8sNodes表接口
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sNodesSearch true "分页获取k8sNodes表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sNodes/getK8sNodesPublic [get]
export const getK8sNodesPublic = () => {
  return service({
    url: '/k8sNodes/getK8sNodesPublic',
    method: 'get',
  })
}
