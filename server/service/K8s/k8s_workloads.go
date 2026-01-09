
package K8s

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
)

type K8sWorkloadsService struct {}
// CreateK8sWorkloads 创建k8sWorkloads表记录
// Author [yourname](https://github.com/yourname)
func (k8sWorkloadsService *K8sWorkloadsService) CreateK8sWorkloads(ctx context.Context, k8sWorkloads *K8s.K8sWorkloads) (err error) {
	err = global.GVA_DB.Create(k8sWorkloads).Error
	return err
}

// DeleteK8sWorkloads 删除k8sWorkloads表记录
// Author [yourname](https://github.com/yourname)
func (k8sWorkloadsService *K8sWorkloadsService)DeleteK8sWorkloads(ctx context.Context, workloadId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sWorkloads{},"workload_id = ?",workloadId).Error
	return err
}

// DeleteK8sWorkloadsByIds 批量删除k8sWorkloads表记录
// Author [yourname](https://github.com/yourname)
func (k8sWorkloadsService *K8sWorkloadsService)DeleteK8sWorkloadsByIds(ctx context.Context, workloadIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sWorkloads{},"workload_id in ?",workloadIds).Error
	return err
}

// UpdateK8sWorkloads 更新k8sWorkloads表记录
// Author [yourname](https://github.com/yourname)
func (k8sWorkloadsService *K8sWorkloadsService)UpdateK8sWorkloads(ctx context.Context, k8sWorkloads K8s.K8sWorkloads) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sWorkloads{}).Where("workload_id = ?",k8sWorkloads.WorkloadId).Updates(&k8sWorkloads).Error
	return err
}

// GetK8sWorkloads 根据workloadId获取k8sWorkloads表记录
// Author [yourname](https://github.com/yourname)
func (k8sWorkloadsService *K8sWorkloadsService)GetK8sWorkloads(ctx context.Context, workloadId string) (k8sWorkloads K8s.K8sWorkloads, err error) {
	err = global.GVA_DB.Where("workload_id = ?", workloadId).First(&k8sWorkloads).Error
	return
}
// GetK8sWorkloadsInfoList 分页获取k8sWorkloads表记录
// Author [yourname](https://github.com/yourname)
func (k8sWorkloadsService *K8sWorkloadsService)GetK8sWorkloadsInfoList(ctx context.Context, info K8sReq.K8sWorkloadsSearch) (list []K8s.K8sWorkloads, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&K8s.K8sWorkloads{})
    var k8sWorkloadss []K8s.K8sWorkloads
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&k8sWorkloadss).Error
	return  k8sWorkloadss, total, err
}
func (k8sWorkloadsService *K8sWorkloadsService)GetK8sWorkloadsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
