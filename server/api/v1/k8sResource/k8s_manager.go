package k8sResource

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/k8sResource"
    k8sResourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sResource/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type K8sResourceApi struct {}



// CreateK8sResource 创建K8s资源
// @Tags K8sResource
// @Summary 创建K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body k8sResource.K8sResource true "创建K8s资源"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /kr/createK8sResource [post]
func (krApi *K8sResourceApi) CreateK8sResource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var kr k8sResource.K8sResource
	err := c.ShouldBindJSON(&kr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = krService.CreateK8sResource(ctx,&kr)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteK8sResource 删除K8s资源
// @Tags K8sResource
// @Summary 删除K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body k8sResource.K8sResource true "删除K8s资源"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /kr/deleteK8sResource [delete]
func (krApi *K8sResourceApi) DeleteK8sResource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := krService.DeleteK8sResource(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sResourceByIds 批量删除K8s资源
// @Tags K8sResource
// @Summary 批量删除K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /kr/deleteK8sResourceByIds [delete]
func (krApi *K8sResourceApi) DeleteK8sResourceByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := krService.DeleteK8sResourceByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sResource 更新K8s资源
// @Tags K8sResource
// @Summary 更新K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body k8sResource.K8sResource true "更新K8s资源"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /kr/updateK8sResource [put]
func (krApi *K8sResourceApi) UpdateK8sResource(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var kr k8sResource.K8sResource
	err := c.ShouldBindJSON(&kr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = krService.UpdateK8sResource(ctx,kr)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sResource 用id查询K8s资源
// @Tags K8sResource
// @Summary 用id查询K8s资源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询K8s资源"
// @Success 200 {object} response.Response{data=k8sResource.K8sResource,msg=string} "查询成功"
// @Router /kr/findK8sResource [get]
func (krApi *K8sResourceApi) FindK8sResource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rekr, err := krService.GetK8sResource(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rekr, c)
}
// GetK8sResourceList 分页获取K8s资源列表
// @Tags K8sResource
// @Summary 分页获取K8s资源列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query k8sResourceReq.K8sResourceSearch true "分页获取K8s资源列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /kr/getK8sResourceList [get]
func (krApi *K8sResourceApi) GetK8sResourceList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo k8sResourceReq.K8sResourceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := krService.GetK8sResourceInfoList(ctx,pageInfo)
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

// GetK8sResourcePublic 不需要鉴权的K8s资源接口
// @Tags K8sResource
// @Summary 不需要鉴权的K8s资源接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /kr/getK8sResourcePublic [get]
func (krApi *K8sResourceApi) GetK8sResourcePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    krService.GetK8sResourcePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的K8s资源接口信息",
    }, "获取成功", c)
}
