import service from '@/utils/request'
// @Tags K8sConfigStorage
// @Summary 创建k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sConfigStorage true "创建k8sConfigStorage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /k8sConfigStorage/createK8sConfigStorage [post]
export const createK8sConfigStorage = (data) => {
  return service({
    url: '/k8sConfigStorage/createK8sConfigStorage',
    method: 'post',
    data
  })
}

// @Tags K8sConfigStorage
// @Summary 删除k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sConfigStorage true "删除k8sConfigStorage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sConfigStorage/deleteK8sConfigStorage [delete]
export const deleteK8sConfigStorage = (params) => {
  return service({
    url: '/k8sConfigStorage/deleteK8sConfigStorage',
    method: 'delete',
    params
  })
}

// @Tags K8sConfigStorage
// @Summary 批量删除k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除k8sConfigStorage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sConfigStorage/deleteK8sConfigStorage [delete]
export const deleteK8sConfigStorageByIds = (params) => {
  return service({
    url: '/k8sConfigStorage/deleteK8sConfigStorageByIds',
    method: 'delete',
    params
  })
}

// @Tags K8sConfigStorage
// @Summary 更新k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sConfigStorage true "更新k8sConfigStorage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sConfigStorage/updateK8sConfigStorage [put]
export const updateK8sConfigStorage = (data) => {
  return service({
    url: '/k8sConfigStorage/updateK8sConfigStorage',
    method: 'put',
    data
  })
}

// @Tags K8sConfigStorage
// @Summary 用id查询k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.K8sConfigStorage true "用id查询k8sConfigStorage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sConfigStorage/findK8sConfigStorage [get]
export const findK8sConfigStorage = (params) => {
  return service({
    url: '/k8sConfigStorage/findK8sConfigStorage',
    method: 'get',
    params
  })
}

// @Tags K8sConfigStorage
// @Summary 分页获取k8sConfigStorage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取k8sConfigStorage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sConfigStorage/getK8sConfigStorageList [get]
export const getK8sConfigStorageList = (params) => {
  return service({
    url: '/k8sConfigStorage/getK8sConfigStorageList',
    method: 'get',
    params
  })
}

// @Tags K8sConfigStorage
// @Summary 不需要鉴权的k8sConfigStorage表接口
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sConfigStorageSearch true "分页获取k8sConfigStorage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sConfigStorage/getK8sConfigStoragePublic [get]
export const getK8sConfigStoragePublic = () => {
  return service({
    url: '/k8sConfigStorage/getK8sConfigStoragePublic',
    method: 'get',
  })
}
