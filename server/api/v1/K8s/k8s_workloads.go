package K8s

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type K8sWorkloadsApi struct {}



// CreateK8sWorkloads 创建k8sWorkloads表
// @Tags K8sWorkloads
// @Summary 创建k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sWorkloads true "创建k8sWorkloads表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8sWorkloads/createK8sWorkloads [post]
func (k8sWorkloadsApi *K8sWorkloadsApi) CreateK8sWorkloads(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var k8sWorkloads K8s.K8sWorkloads
	err := c.ShouldBindJSON(&k8sWorkloads)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sWorkloadsService.CreateK8sWorkloads(ctx,&k8sWorkloads)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteK8sWorkloads 删除k8sWorkloads表
// @Tags K8sWorkloads
// @Summary 删除k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sWorkloads true "删除k8sWorkloads表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8sWorkloads/deleteK8sWorkloads [delete]
func (k8sWorkloadsApi *K8sWorkloadsApi) DeleteK8sWorkloads(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	workloadId := c.Query("workloadId")
	err := k8sWorkloadsService.DeleteK8sWorkloads(ctx,workloadId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sWorkloadsByIds 批量删除k8sWorkloads表
// @Tags K8sWorkloads
// @Summary 批量删除k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /k8sWorkloads/deleteK8sWorkloadsByIds [delete]
func (k8sWorkloadsApi *K8sWorkloadsApi) DeleteK8sWorkloadsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	workloadIds := c.QueryArray("workloadIds[]")
	err := k8sWorkloadsService.DeleteK8sWorkloadsByIds(ctx,workloadIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sWorkloads 更新k8sWorkloads表
// @Tags K8sWorkloads
// @Summary 更新k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sWorkloads true "更新k8sWorkloads表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /k8sWorkloads/updateK8sWorkloads [put]
func (k8sWorkloadsApi *K8sWorkloadsApi) UpdateK8sWorkloads(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var k8sWorkloads K8s.K8sWorkloads
	err := c.ShouldBindJSON(&k8sWorkloads)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sWorkloadsService.UpdateK8sWorkloads(ctx,k8sWorkloads)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sWorkloads 用id查询k8sWorkloads表
// @Tags K8sWorkloads
// @Summary 用id查询k8sWorkloads表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param workloadId query int true "用id查询k8sWorkloads表"
// @Success 200 {object} response.Response{data=K8s.K8sWorkloads,msg=string} "查询成功"
// @Router /k8sWorkloads/findK8sWorkloads [get]
func (k8sWorkloadsApi *K8sWorkloadsApi) FindK8sWorkloads(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	workloadId := c.Query("workloadId")
	rek8sWorkloads, err := k8sWorkloadsService.GetK8sWorkloads(ctx,workloadId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rek8sWorkloads, c)
}
// GetK8sWorkloadsList 分页获取k8sWorkloads表列表
// @Tags K8sWorkloads
// @Summary 分页获取k8sWorkloads表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sWorkloadsSearch true "分页获取k8sWorkloads表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8sWorkloads/getK8sWorkloadsList [get]
func (k8sWorkloadsApi *K8sWorkloadsApi) GetK8sWorkloadsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo K8sReq.K8sWorkloadsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := k8sWorkloadsService.GetK8sWorkloadsInfoList(ctx,pageInfo)
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

// GetK8sWorkloadsPublic 不需要鉴权的k8sWorkloads表接口
// @Tags K8sWorkloads
// @Summary 不需要鉴权的k8sWorkloads表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sWorkloads/getK8sWorkloadsPublic [get]
func (k8sWorkloadsApi *K8sWorkloadsApi) GetK8sWorkloadsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    k8sWorkloadsService.GetK8sWorkloadsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的k8sWorkloads表接口信息",
    }, "获取成功", c)
}
