package K8s

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type K8sNamespacesApi struct {}



// CreateK8sNamespaces 创建k8sNamespaces表
// @Tags K8sNamespaces
// @Summary 创建k8sNamespaces表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sNamespaces true "创建k8sNamespaces表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8sNamespaces/createK8sNamespaces [post]
func (k8sNamespacesApi *K8sNamespacesApi) CreateK8sNamespaces(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var k8sNamespaces K8s.K8sNamespaces
	err := c.ShouldBindJSON(&k8sNamespaces)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sNamespacesService.CreateK8sNamespaces(ctx,&k8sNamespaces)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteK8sNamespaces 删除k8sNamespaces表
// @Tags K8sNamespaces
// @Summary 删除k8sNamespaces表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sNamespaces true "删除k8sNamespaces表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8sNamespaces/deleteK8sNamespaces [delete]
func (k8sNamespacesApi *K8sNamespacesApi) DeleteK8sNamespaces(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	namespaceId := c.Query("namespaceId")
	err := k8sNamespacesService.DeleteK8sNamespaces(ctx,namespaceId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sNamespacesByIds 批量删除k8sNamespaces表
// @Tags K8sNamespaces
// @Summary 批量删除k8sNamespaces表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /k8sNamespaces/deleteK8sNamespacesByIds [delete]
func (k8sNamespacesApi *K8sNamespacesApi) DeleteK8sNamespacesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	namespaceIds := c.QueryArray("namespaceIds[]")
	err := k8sNamespacesService.DeleteK8sNamespacesByIds(ctx,namespaceIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sNamespaces 更新k8sNamespaces表
// @Tags K8sNamespaces
// @Summary 更新k8sNamespaces表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sNamespaces true "更新k8sNamespaces表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /k8sNamespaces/updateK8sNamespaces [put]
func (k8sNamespacesApi *K8sNamespacesApi) UpdateK8sNamespaces(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var k8sNamespaces K8s.K8sNamespaces
	err := c.ShouldBindJSON(&k8sNamespaces)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sNamespacesService.UpdateK8sNamespaces(ctx,k8sNamespaces)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sNamespaces 用id查询k8sNamespaces表
// @Tags K8sNamespaces
// @Summary 用id查询k8sNamespaces表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param namespaceId query int true "用id查询k8sNamespaces表"
// @Success 200 {object} response.Response{data=K8s.K8sNamespaces,msg=string} "查询成功"
// @Router /k8sNamespaces/findK8sNamespaces [get]
func (k8sNamespacesApi *K8sNamespacesApi) FindK8sNamespaces(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	namespaceId := c.Query("namespaceId")
	rek8sNamespaces, err := k8sNamespacesService.GetK8sNamespaces(ctx,namespaceId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rek8sNamespaces, c)
}
// GetK8sNamespacesList 分页获取k8sNamespaces表列表
// @Tags K8sNamespaces
// @Summary 分页获取k8sNamespaces表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sNamespacesSearch true "分页获取k8sNamespaces表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8sNamespaces/getK8sNamespacesList [get]
func (k8sNamespacesApi *K8sNamespacesApi) GetK8sNamespacesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo K8sReq.K8sNamespacesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := k8sNamespacesService.GetK8sNamespacesInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetK8sNamespacesPublic 不需要鉴权的k8sNamespaces表接口
// @Tags K8sNamespaces
// @Summary 不需要鉴权的k8sNamespaces表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sNamespaces/getK8sNamespacesPublic [get]
func (k8sNamespacesApi *K8sNamespacesApi) GetK8sNamespacesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    k8sNamespacesService.GetK8sNamespacesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的k8sNamespaces表接口信息",
    }, "获取成功", c)
}
