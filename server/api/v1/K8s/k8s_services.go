package K8s

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type K8sServicesApi struct {}



// CreateK8sServices 创建k8sServices表
// @Tags K8sServices
// @Summary 创建k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sServices true "创建k8sServices表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8sServices/createK8sServices [post]
func (k8sServicesApi *K8sServicesApi) CreateK8sServices(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var k8sServices K8s.K8sServices
	err := c.ShouldBindJSON(&k8sServices)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sServicesService.CreateK8sServices(ctx,&k8sServices)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteK8sServices 删除k8sServices表
// @Tags K8sServices
// @Summary 删除k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sServices true "删除k8sServices表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8sServices/deleteK8sServices [delete]
func (k8sServicesApi *K8sServicesApi) DeleteK8sServices(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	serviceId := c.Query("serviceId")
	err := k8sServicesService.DeleteK8sServices(ctx,serviceId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sServicesByIds 批量删除k8sServices表
// @Tags K8sServices
// @Summary 批量删除k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /k8sServices/deleteK8sServicesByIds [delete]
func (k8sServicesApi *K8sServicesApi) DeleteK8sServicesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	serviceIds := c.QueryArray("serviceIds[]")
	err := k8sServicesService.DeleteK8sServicesByIds(ctx,serviceIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sServices 更新k8sServices表
// @Tags K8sServices
// @Summary 更新k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sServices true "更新k8sServices表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /k8sServices/updateK8sServices [put]
func (k8sServicesApi *K8sServicesApi) UpdateK8sServices(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var k8sServices K8s.K8sServices
	err := c.ShouldBindJSON(&k8sServices)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sServicesService.UpdateK8sServices(ctx,k8sServices)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sServices 用id查询k8sServices表
// @Tags K8sServices
// @Summary 用id查询k8sServices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param serviceId query int true "用id查询k8sServices表"
// @Success 200 {object} response.Response{data=K8s.K8sServices,msg=string} "查询成功"
// @Router /k8sServices/findK8sServices [get]
func (k8sServicesApi *K8sServicesApi) FindK8sServices(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	serviceId := c.Query("serviceId")
	rek8sServices, err := k8sServicesService.GetK8sServices(ctx,serviceId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rek8sServices, c)
}
// GetK8sServicesList 分页获取k8sServices表列表
// @Tags K8sServices
// @Summary 分页获取k8sServices表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sServicesSearch true "分页获取k8sServices表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8sServices/getK8sServicesList [get]
func (k8sServicesApi *K8sServicesApi) GetK8sServicesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo K8sReq.K8sServicesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := k8sServicesService.GetK8sServicesInfoList(ctx,pageInfo)
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

// GetK8sServicesPublic 不需要鉴权的k8sServices表接口
// @Tags K8sServices
// @Summary 不需要鉴权的k8sServices表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sServices/getK8sServicesPublic [get]
func (k8sServicesApi *K8sServicesApi) GetK8sServicesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    k8sServicesService.GetK8sServicesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的k8sServices表接口信息",
    }, "获取成功", c)
}
