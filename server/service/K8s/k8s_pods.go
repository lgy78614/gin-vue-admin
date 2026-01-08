
package K8s

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
)

type K8sPodsService struct {}
// CreateK8sPods 创建k8sPods表记录
// Author [yourname](https://github.com/yourname)
func (k8sPodsService *K8sPodsService) CreateK8sPods(ctx context.Context, k8sPods *K8s.K8sPods) (err error) {
	err = global.GVA_DB.Create(k8sPods).Error
	return err
}

// DeleteK8sPods 删除k8sPods表记录
// Author [yourname](https://github.com/yourname)
func (k8sPodsService *K8sPodsService)DeleteK8sPods(ctx context.Context, podId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sPods{},"pod_id = ?",podId).Error
	return err
}

// DeleteK8sPodsByIds 批量删除k8sPods表记录
// Author [yourname](https://github.com/yourname)
func (k8sPodsService *K8sPodsService)DeleteK8sPodsByIds(ctx context.Context, podIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sPods{},"pod_id in ?",podIds).Error
	return err
}

// UpdateK8sPods 更新k8sPods表记录
// Author [yourname](https://github.com/yourname)
func (k8sPodsService *K8sPodsService)UpdateK8sPods(ctx context.Context, k8sPods K8s.K8sPods) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sPods{}).Where("pod_id = ?",k8sPods.PodId).Updates(&k8sPods).Error
	return err
}

// GetK8sPods 根据podId获取k8sPods表记录
// Author [yourname](https://github.com/yourname)
func (k8sPodsService *K8sPodsService)GetK8sPods(ctx context.Context, podId string) (k8sPods K8s.K8sPods, err error) {
	err = global.GVA_DB.Where("pod_id = ?", podId).First(&k8sPods).Error
	return
}
// GetK8sPodsInfoList 分页获取k8sPods表记录
// Author [yourname](https://github.com/yourname)
func (k8sPodsService *K8sPodsService)GetK8sPodsInfoList(ctx context.Context, info K8sReq.K8sPodsSearch) (list []K8s.K8sPods, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&K8s.K8sPods{})
    var k8sPodss []K8s.K8sPods
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.ClusterId != nil && *info.ClusterId != "" {
        db = db.Where("cluster_id = ?", *info.ClusterId)
    }
    if info.Namespace != nil && *info.Namespace != "" {
        db = db.Where("namespace = ?", *info.Namespace)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&k8sPodss).Error
	return  k8sPodss, total, err
}
func (k8sPodsService *K8sPodsService)GetK8sPodsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
