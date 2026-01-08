package K8s

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type K8sNodesApi struct {}



// CreateK8sNodes 创建k8sNodes表
// @Tags K8sNodes
// @Summary 创建k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sNodes true "创建k8sNodes表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8sNodes/createK8sNodes [post]
func (k8sNodesApi *K8sNodesApi) CreateK8sNodes(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var k8sNodes K8s.K8sNodes
	err := c.ShouldBindJSON(&k8sNodes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sNodesService.CreateK8sNodes(ctx,&k8sNodes)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteK8sNodes 删除k8sNodes表
// @Tags K8sNodes
// @Summary 删除k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sNodes true "删除k8sNodes表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8sNodes/deleteK8sNodes [delete]
func (k8sNodesApi *K8sNodesApi) DeleteK8sNodes(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	nodeId := c.Query("nodeId")
	err := k8sNodesService.DeleteK8sNodes(ctx,nodeId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sNodesByIds 批量删除k8sNodes表
// @Tags K8sNodes
// @Summary 批量删除k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /k8sNodes/deleteK8sNodesByIds [delete]
func (k8sNodesApi *K8sNodesApi) DeleteK8sNodesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	nodeIds := c.QueryArray("nodeIds[]")
	err := k8sNodesService.DeleteK8sNodesByIds(ctx,nodeIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sNodes 更新k8sNodes表
// @Tags K8sNodes
// @Summary 更新k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sNodes true "更新k8sNodes表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /k8sNodes/updateK8sNodes [put]
func (k8sNodesApi *K8sNodesApi) UpdateK8sNodes(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var k8sNodes K8s.K8sNodes
	err := c.ShouldBindJSON(&k8sNodes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sNodesService.UpdateK8sNodes(ctx,k8sNodes)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sNodes 用id查询k8sNodes表
// @Tags K8sNodes
// @Summary 用id查询k8sNodes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param nodeId query string true "用id查询k8sNodes表"
// @Success 200 {object} response.Response{data=K8s.K8sNodes,msg=string} "查询成功"
// @Router /k8sNodes/findK8sNodes [get]
func (k8sNodesApi *K8sNodesApi) FindK8sNodes(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	nodeId := c.Query("nodeId")
	rek8sNodes, err := k8sNodesService.GetK8sNodes(ctx,nodeId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rek8sNodes, c)
}
// GetK8sNodesList 分页获取k8sNodes表列表
// @Tags K8sNodes
// @Summary 分页获取k8sNodes表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sNodesSearch true "分页获取k8sNodes表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8sNodes/getK8sNodesList [get]
func (k8sNodesApi *K8sNodesApi) GetK8sNodesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo K8sReq.K8sNodesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := k8sNodesService.GetK8sNodesInfoList(ctx,pageInfo)
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

// GetK8sNodesPublic 不需要鉴权的k8sNodes表接口
// @Tags K8sNodes
// @Summary 不需要鉴权的k8sNodes表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sNodes/getK8sNodesPublic [get]
func (k8sNodesApi *K8sNodesApi) GetK8sNodesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    k8sNodesService.GetK8sNodesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的k8sNodes表接口信息",
    }, "获取成功", c)
}
