
package K8s

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
)

type K8sNamespacesService struct {}
// CreateK8sNamespaces 创建k8sNamespaces表记录
// Author [yourname](https://github.com/yourname)
func (k8sNamespacesService *K8sNamespacesService) CreateK8sNamespaces(ctx context.Context, k8sNamespaces *K8s.K8sNamespaces) (err error) {
	err = global.GVA_DB.Create(k8sNamespaces).Error
	return err
}

// DeleteK8sNamespaces 删除k8sNamespaces表记录
// Author [yourname](https://github.com/yourname)
func (k8sNamespacesService *K8sNamespacesService)DeleteK8sNamespaces(ctx context.Context, namespaceId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sNamespaces{},"namespace_id = ?",namespaceId).Error
	return err
}

// DeleteK8sNamespacesByIds 批量删除k8sNamespaces表记录
// Author [yourname](https://github.com/yourname)
func (k8sNamespacesService *K8sNamespacesService)DeleteK8sNamespacesByIds(ctx context.Context, namespaceIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sNamespaces{},"namespace_id in ?",namespaceIds).Error
	return err
}

// UpdateK8sNamespaces 更新k8sNamespaces表记录
// Author [yourname](https://github.com/yourname)
func (k8sNamespacesService *K8sNamespacesService)UpdateK8sNamespaces(ctx context.Context, k8sNamespaces K8s.K8sNamespaces) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sNamespaces{}).Where("namespace_id = ?",k8sNamespaces.NamespaceId).Updates(&k8sNamespaces).Error
	return err
}

// GetK8sNamespaces 根据namespaceId获取k8sNamespaces表记录
// Author [yourname](https://github.com/yourname)
func (k8sNamespacesService *K8sNamespacesService)GetK8sNamespaces(ctx context.Context, namespaceId string) (k8sNamespaces K8s.K8sNamespaces, err error) {
	err = global.GVA_DB.Where("namespace_id = ?", namespaceId).First(&k8sNamespaces).Error
	return
}
// GetK8sNamespacesInfoList 分页获取k8sNamespaces表记录
// Author [yourname](https://github.com/yourname)
func (k8sNamespacesService *K8sNamespacesService)GetK8sNamespacesInfoList(ctx context.Context, info K8sReq.K8sNamespacesSearch) (list []K8s.K8sNamespaces, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&K8s.K8sNamespaces{})
    var k8sNamespacess []K8s.K8sNamespaces
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&k8sNamespacess).Error
	return  k8sNamespacess, total, err
}
func (k8sNamespacesService *K8sNamespacesService)GetK8sNamespacesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
