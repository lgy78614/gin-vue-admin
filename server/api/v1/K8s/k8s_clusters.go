package K8s

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type K8sClustersApi struct {}



// CreateK8sClusters 创建k8sClusters表
// @Tags K8sClusters
// @Summary 创建k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sClusters true "创建k8sClusters表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8sClusters/createK8sClusters [post]
func (k8sClustersApi *K8sClustersApi) CreateK8sClusters(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var k8sClusters K8s.K8sClusters
	err := c.ShouldBindJSON(&k8sClusters)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sClustersService.CreateK8sClusters(ctx,&k8sClusters)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteK8sClusters 删除k8sClusters表
// @Tags K8sClusters
// @Summary 删除k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sClusters true "删除k8sClusters表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8sClusters/deleteK8sClusters [delete]
func (k8sClustersApi *K8sClustersApi) DeleteK8sClusters(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	clusterId := c.Query("clusterId")
	err := k8sClustersService.DeleteK8sClusters(ctx,clusterId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sClustersByIds 批量删除k8sClusters表
// @Tags K8sClusters
// @Summary 批量删除k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /k8sClusters/deleteK8sClustersByIds [delete]
func (k8sClustersApi *K8sClustersApi) DeleteK8sClustersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	clusterIds := c.QueryArray("clusterIds[]")
	err := k8sClustersService.DeleteK8sClustersByIds(ctx,clusterIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sClusters 更新k8sClusters表
// @Tags K8sClusters
// @Summary 更新k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sClusters true "更新k8sClusters表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /k8sClusters/updateK8sClusters [put]
func (k8sClustersApi *K8sClustersApi) UpdateK8sClusters(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var k8sClusters K8s.K8sClusters
	err := c.ShouldBindJSON(&k8sClusters)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sClustersService.UpdateK8sClusters(ctx,k8sClusters)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sClusters 用id查询k8sClusters表
// @Tags K8sClusters
// @Summary 用id查询k8sClusters表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param clusterId query string true "用id查询k8sClusters表"
// @Success 200 {object} response.Response{data=K8s.K8sClusters,msg=string} "查询成功"
// @Router /k8sClusters/findK8sClusters [get]
func (k8sClustersApi *K8sClustersApi) FindK8sClusters(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	clusterId := c.Query("clusterId")
	rek8sClusters, err := k8sClustersService.GetK8sClusters(ctx,clusterId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rek8sClusters, c)
}
// GetK8sClustersList 分页获取k8sClusters表列表
// @Tags K8sClusters
// @Summary 分页获取k8sClusters表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sClustersSearch true "分页获取k8sClusters表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8sClusters/getK8sClustersList [get]
func (k8sClustersApi *K8sClustersApi) GetK8sClustersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo K8sReq.K8sClustersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := k8sClustersService.GetK8sClustersInfoList(ctx,pageInfo)
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

// GetK8sClustersPublic 不需要鉴权的k8sClusters表接口
// @Tags K8sClusters
// @Summary 不需要鉴权的k8sClusters表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sClusters/getK8sClustersPublic [get]
func (k8sClustersApi *K8sClustersApi) GetK8sClustersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    k8sClustersService.GetK8sClustersPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的k8sClusters表接口信息",
    }, "获取成功", c)
}
