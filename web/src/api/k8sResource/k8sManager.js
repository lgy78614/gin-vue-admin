import service from '@/utils/request'
// @Tags K8sResource
// @Summary 创建K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sResource true "创建K8s资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /kr/createK8sResource [post]
export const createK8sResource = (data) => {
  return service({
    url: '/kr/createK8sResource',
    method: 'post',
    data
  })
}

// @Tags K8sResource
// @Summary 删除K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sResource true "删除K8s资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /kr/deleteK8sResource [delete]
export const deleteK8sResource = (params) => {
  return service({
    url: '/kr/deleteK8sResource',
    method: 'delete',
    params
  })
}

// @Tags K8sResource
// @Summary 批量删除K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8s资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /kr/deleteK8sResource [delete]
export const deleteK8sResourceByIds = (params) => {
  return service({
    url: '/kr/deleteK8sResourceByIds',
    method: 'delete',
    params
  })
}

// @Tags K8sResource
// @Summary 更新K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.K8sResource true "更新K8s资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /kr/updateK8sResource [put]
export const updateK8sResource = (data) => {
  return service({
    url: '/kr/updateK8sResource',
    method: 'put',
    data
  })
}

// @Tags K8sResource
// @Summary 用id查询K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.K8sResource true "用id查询K8s资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /kr/findK8sResource [get]
export const findK8sResource = (params) => {
  return service({
    url: '/kr/findK8sResource',
    method: 'get',
    params
  })
}

// @Tags K8sResource
// @Summary 分页获取K8s资源列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取K8s资源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kr/getK8sResourceList [get]
export const getK8sResourceList = (params) => {
  return service({
    url: '/kr/getK8sResourceList',
    method: 'get',
    params
  })
}

// @Tags K8sResource
// @Summary 不需要鉴权的K8s资源接口
// @Accept application/json
// @Produce application/json
// @Param data query k8sResourceReq.K8sResourceSearch true "分页获取K8s资源列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /kr/getK8sResourcePublic [get]
export const getK8sResourcePublic = () => {
  return service({
    url: '/kr/getK8sResourcePublic',
    method: 'get',
  })
}
