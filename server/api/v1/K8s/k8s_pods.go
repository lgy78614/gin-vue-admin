package K8s

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type K8sPodsApi struct {}



// CreateK8sPods 创建k8sPods表
// @Tags K8sPods
// @Summary 创建k8sPods表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sPods true "创建k8sPods表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8sPods/createK8sPods [post]
func (k8sPodsApi *K8sPodsApi) CreateK8sPods(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var k8sPods K8s.K8sPods
	err := c.ShouldBindJSON(&k8sPods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sPodsService.CreateK8sPods(ctx,&k8sPods)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteK8sPods 删除k8sPods表
// @Tags K8sPods
// @Summary 删除k8sPods表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sPods true "删除k8sPods表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8sPods/deleteK8sPods [delete]
func (k8sPodsApi *K8sPodsApi) DeleteK8sPods(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	podId := c.Query("podId")
	err := k8sPodsService.DeleteK8sPods(ctx,podId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sPodsByIds 批量删除k8sPods表
// @Tags K8sPods
// @Summary 批量删除k8sPods表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /k8sPods/deleteK8sPodsByIds [delete]
func (k8sPodsApi *K8sPodsApi) DeleteK8sPodsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	podIds := c.QueryArray("podIds[]")
	err := k8sPodsService.DeleteK8sPodsByIds(ctx,podIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sPods 更新k8sPods表
// @Tags K8sPods
// @Summary 更新k8sPods表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sPods true "更新k8sPods表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /k8sPods/updateK8sPods [put]
func (k8sPodsApi *K8sPodsApi) UpdateK8sPods(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var k8sPods K8s.K8sPods
	err := c.ShouldBindJSON(&k8sPods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sPodsService.UpdateK8sPods(ctx,k8sPods)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sPods 用id查询k8sPods表
// @Tags K8sPods
// @Summary 用id查询k8sPods表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param podId query int true "用id查询k8sPods表"
// @Success 200 {object} response.Response{data=K8s.K8sPods,msg=string} "查询成功"
// @Router /k8sPods/findK8sPods [get]
func (k8sPodsApi *K8sPodsApi) FindK8sPods(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	podId := c.Query("podId")
	rek8sPods, err := k8sPodsService.GetK8sPods(ctx,podId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rek8sPods, c)
}
// GetK8sPodsList 分页获取k8sPods表列表
// @Tags K8sPods
// @Summary 分页获取k8sPods表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sPodsSearch true "分页获取k8sPods表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8sPods/getK8sPodsList [get]
func (k8sPodsApi *K8sPodsApi) GetK8sPodsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo K8sReq.K8sPodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := k8sPodsService.GetK8sPodsInfoList(ctx,pageInfo)
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

// GetK8sPodsPublic 不需要鉴权的k8sPods表接口
// @Tags K8sPods
// @Summary 不需要鉴权的k8sPods表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sPods/getK8sPodsPublic [get]
func (k8sPodsApi *K8sPodsApi) GetK8sPodsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    k8sPodsService.GetK8sPodsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的k8sPods表接口信息",
    }, "获取成功", c)
}
