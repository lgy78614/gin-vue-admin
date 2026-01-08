import service from '@/utils/request'
// @Tags K8sClusters
// @Summary 创建k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sClusters true "创建k8sClusters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sClusters/createK8sClusters [post]
export const createK8sClusters = (data) => {
  return service({
    url: '/k8sClusters/createK8sClusters',
    method: 'post',
    data
  })
}

// @Tags K8sClusters
// @Summary 删除k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sClusters true "删除k8sClusters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sClusters/deleteK8sClusters [delete]
export const deleteK8sClusters = (params) => {
  return service({
    url: '/k8sClusters/deleteK8sClusters',
    method: 'delete',
    params
  })
}

// @Tags K8sClusters
// @Summary 批量删除k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除k8sClusters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sClusters/deleteK8sClusters [delete]
export const deleteK8sClustersByIds = (params) => {
  return service({
    url: '/k8sClusters/deleteK8sClustersByIds',
    method: 'delete',
    params
  })
}

// @Tags K8sClusters
// @Summary 更新k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sClusters true "更新k8sClusters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sClusters/updateK8sClusters [put]
export const updateK8sClusters = (data) => {
  return service({
    url: '/k8sClusters/updateK8sClusters',
    method: 'put',
    data
  })
}

// @Tags K8sClusters
// @Summary 用id查询k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.K8sClusters true "用id查询k8sClusters表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sClusters/findK8sClusters [get]
export const findK8sClusters = (params) => {
  return service({
    url: '/k8sClusters/findK8sClusters',
    method: 'get',
    params
  })
}

// @Tags K8sClusters
// @Summary 分页获取k8sClusters表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取k8sClusters表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sClusters/getK8sClustersList [get]
export const getK8sClustersList = (params) => {
  return service({
    url: '/k8sClusters/getK8sClustersList',
    method: 'get',
    params
  })
}

// @Tags K8sClusters
// @Summary 不需要鉴权的k8sClusters表接口
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sClustersSearch true "分页获取k8sClusters表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sClusters/getK8sClustersPublic [get]
export const getK8sClustersPublic = () => {
  return service({
    url: '/k8sClusters/getK8sClustersPublic',
    method: 'get',
  })
}
