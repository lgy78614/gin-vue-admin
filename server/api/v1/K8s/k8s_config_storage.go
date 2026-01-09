package K8s

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type K8sConfigStorageApi struct {}



// CreateK8sConfigStorage 创建k8sConfigStorage表
// @Tags K8sConfigStorage
// @Summary 创建k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sConfigStorage true "创建k8sConfigStorage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /k8sConfigStorage/createK8sConfigStorage [post]
func (k8sConfigStorageApi *K8sConfigStorageApi) CreateK8sConfigStorage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var k8sConfigStorage K8s.K8sConfigStorage
	err := c.ShouldBindJSON(&k8sConfigStorage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sConfigStorageService.CreateK8sConfigStorage(ctx,&k8sConfigStorage)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteK8sConfigStorage 删除k8sConfigStorage表
// @Tags K8sConfigStorage
// @Summary 删除k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sConfigStorage true "删除k8sConfigStorage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /k8sConfigStorage/deleteK8sConfigStorage [delete]
func (k8sConfigStorageApi *K8sConfigStorageApi) DeleteK8sConfigStorage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	configId := c.Query("configId")
	err := k8sConfigStorageService.DeleteK8sConfigStorage(ctx,configId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sConfigStorageByIds 批量删除k8sConfigStorage表
// @Tags K8sConfigStorage
// @Summary 批量删除k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /k8sConfigStorage/deleteK8sConfigStorageByIds [delete]
func (k8sConfigStorageApi *K8sConfigStorageApi) DeleteK8sConfigStorageByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	configIds := c.QueryArray("configIds[]")
	err := k8sConfigStorageService.DeleteK8sConfigStorageByIds(ctx,configIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sConfigStorage 更新k8sConfigStorage表
// @Tags K8sConfigStorage
// @Summary 更新k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body K8s.K8sConfigStorage true "更新k8sConfigStorage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /k8sConfigStorage/updateK8sConfigStorage [put]
func (k8sConfigStorageApi *K8sConfigStorageApi) UpdateK8sConfigStorage(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var k8sConfigStorage K8s.K8sConfigStorage
	err := c.ShouldBindJSON(&k8sConfigStorage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = k8sConfigStorageService.UpdateK8sConfigStorage(ctx,k8sConfigStorage)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sConfigStorage 用id查询k8sConfigStorage表
// @Tags K8sConfigStorage
// @Summary 用id查询k8sConfigStorage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param configId query int true "用id查询k8sConfigStorage表"
// @Success 200 {object} response.Response{data=K8s.K8sConfigStorage,msg=string} "查询成功"
// @Router /k8sConfigStorage/findK8sConfigStorage [get]
func (k8sConfigStorageApi *K8sConfigStorageApi) FindK8sConfigStorage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	configId := c.Query("configId")
	rek8sConfigStorage, err := k8sConfigStorageService.GetK8sConfigStorage(ctx,configId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rek8sConfigStorage, c)
}
// GetK8sConfigStorageList 分页获取k8sConfigStorage表列表
// @Tags K8sConfigStorage
// @Summary 分页获取k8sConfigStorage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query K8sReq.K8sConfigStorageSearch true "分页获取k8sConfigStorage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /k8sConfigStorage/getK8sConfigStorageList [get]
func (k8sConfigStorageApi *K8sConfigStorageApi) GetK8sConfigStorageList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo K8sReq.K8sConfigStorageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := k8sConfigStorageService.GetK8sConfigStorageInfoList(ctx,pageInfo)
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

// GetK8sConfigStoragePublic 不需要鉴权的k8sConfigStorage表接口
// @Tags K8sConfigStorage
// @Summary 不需要鉴权的k8sConfigStorage表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /k8sConfigStorage/getK8sConfigStoragePublic [get]
func (k8sConfigStorageApi *K8sConfigStorageApi) GetK8sConfigStoragePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    k8sConfigStorageService.GetK8sConfigStoragePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的k8sConfigStorage表接口信息",
    }, "获取成功", c)
}
