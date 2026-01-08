
package K8s

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
)

type K8sServicesService struct {}
// CreateK8sServices 创建k8sServices表记录
// Author [yourname](https://github.com/yourname)
func (k8sServicesService *K8sServicesService) CreateK8sServices(ctx context.Context, k8sServices *K8s.K8sServices) (err error) {
	err = global.GVA_DB.Create(k8sServices).Error
	return err
}

// DeleteK8sServices 删除k8sServices表记录
// Author [yourname](https://github.com/yourname)
func (k8sServicesService *K8sServicesService)DeleteK8sServices(ctx context.Context, serviceId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sServices{},"service_id = ?",serviceId).Error
	return err
}

// DeleteK8sServicesByIds 批量删除k8sServices表记录
// Author [yourname](https://github.com/yourname)
func (k8sServicesService *K8sServicesService)DeleteK8sServicesByIds(ctx context.Context, serviceIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sServices{},"service_id in ?",serviceIds).Error
	return err
}

// UpdateK8sServices 更新k8sServices表记录
// Author [yourname](https://github.com/yourname)
func (k8sServicesService *K8sServicesService)UpdateK8sServices(ctx context.Context, k8sServices K8s.K8sServices) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sServices{}).Where("service_id = ?",k8sServices.ServiceId).Updates(&k8sServices).Error
	return err
}

// GetK8sServices 根据serviceId获取k8sServices表记录
// Author [yourname](https://github.com/yourname)
func (k8sServicesService *K8sServicesService)GetK8sServices(ctx context.Context, serviceId string) (k8sServices K8s.K8sServices, err error) {
	err = global.GVA_DB.Where("service_id = ?", serviceId).First(&k8sServices).Error
	return
}
// GetK8sServicesInfoList 分页获取k8sServices表记录
// Author [yourname](https://github.com/yourname)
func (k8sServicesService *K8sServicesService)GetK8sServicesInfoList(ctx context.Context, info K8sReq.K8sServicesSearch) (list []K8s.K8sServices, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&K8s.K8sServices{})
    var k8sServicess []K8s.K8sServices
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.ServiceId != nil {
        db = db.Where("service_id = ?", *info.ServiceId)
    }
    if info.ClusterId != nil && *info.ClusterId != "" {
        db = db.Where("cluster_id = ?", *info.ClusterId)
    }
    if info.Namespace != nil && *info.Namespace != "" {
        db = db.Where("namespace = ?", *info.Namespace)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?", *info.Name)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&k8sServicess).Error
	return  k8sServicess, total, err
}
func (k8sServicesService *K8sServicesService)GetK8sServicesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
