
package K8s

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
)

type K8sNodesService struct {}
// CreateK8sNodes 创建k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService) CreateK8sNodes(ctx context.Context, k8sNodes *K8s.K8sNodes) (err error) {
	err = global.GVA_DB.Create(k8sNodes).Error
	return err
}

// DeleteK8sNodes 删除k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService)DeleteK8sNodes(ctx context.Context, nodeId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sNodes{},"node_id = ?",nodeId).Error
	return err
}

// DeleteK8sNodesByIds 批量删除k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService)DeleteK8sNodesByIds(ctx context.Context, nodeIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sNodes{},"node_id in ?",nodeIds).Error
	return err
}

// UpdateK8sNodes 更新k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService)UpdateK8sNodes(ctx context.Context, k8sNodes K8s.K8sNodes) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sNodes{}).Where("node_id = ?",k8sNodes.NodeId).Updates(&k8sNodes).Error
	return err
}

// GetK8sNodes 根据nodeId获取k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService)GetK8sNodes(ctx context.Context, nodeId string) (k8sNodes K8s.K8sNodes, err error) {
	err = global.GVA_DB.Where("node_id = ?", nodeId).First(&k8sNodes).Error
	return
}
// GetK8sNodesInfoList 分页获取k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService)GetK8sNodesInfoList(ctx context.Context, info K8sReq.K8sNodesSearch) (list []K8s.K8sNodes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&K8s.K8sNodes{})
    var k8sNodess []K8s.K8sNodes
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&k8sNodess).Error
	return  k8sNodess, total, err
}
func (k8sNodesService *K8sNodesService)GetK8sNodesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
