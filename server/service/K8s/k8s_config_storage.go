
package K8s

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
)

type K8sConfigStorageService struct {}
// CreateK8sConfigStorage 创建k8sConfigStorage表记录
// Author [yourname](https://github.com/yourname)
func (k8sConfigStorageService *K8sConfigStorageService) CreateK8sConfigStorage(ctx context.Context, k8sConfigStorage *K8s.K8sConfigStorage) (err error) {
	err = global.GVA_DB.Create(k8sConfigStorage).Error
	return err
}

// DeleteK8sConfigStorage 删除k8sConfigStorage表记录
// Author [yourname](https://github.com/yourname)
func (k8sConfigStorageService *K8sConfigStorageService)DeleteK8sConfigStorage(ctx context.Context, configId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sConfigStorage{},"config_id = ?",configId).Error
	return err
}

// DeleteK8sConfigStorageByIds 批量删除k8sConfigStorage表记录
// Author [yourname](https://github.com/yourname)
func (k8sConfigStorageService *K8sConfigStorageService)DeleteK8sConfigStorageByIds(ctx context.Context, configIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sConfigStorage{},"config_id in ?",configIds).Error
	return err
}

// UpdateK8sConfigStorage 更新k8sConfigStorage表记录
// Author [yourname](https://github.com/yourname)
func (k8sConfigStorageService *K8sConfigStorageService)UpdateK8sConfigStorage(ctx context.Context, k8sConfigStorage K8s.K8sConfigStorage) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sConfigStorage{}).Where("config_id = ?",k8sConfigStorage.ConfigId).Updates(&k8sConfigStorage).Error
	return err
}

// GetK8sConfigStorage 根据configId获取k8sConfigStorage表记录
// Author [yourname](https://github.com/yourname)
func (k8sConfigStorageService *K8sConfigStorageService)GetK8sConfigStorage(ctx context.Context, configId string) (k8sConfigStorage K8s.K8sConfigStorage, err error) {
	err = global.GVA_DB.Where("config_id = ?", configId).First(&k8sConfigStorage).Error
	return
}
// GetK8sConfigStorageInfoList 分页获取k8sConfigStorage表记录
// Author [yourname](https://github.com/yourname)
func (k8sConfigStorageService *K8sConfigStorageService)GetK8sConfigStorageInfoList(ctx context.Context, info K8sReq.K8sConfigStorageSearch) (list []K8s.K8sConfigStorage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&K8s.K8sConfigStorage{})
    var k8sConfigStorages []K8s.K8sConfigStorage
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&k8sConfigStorages).Error
	return  k8sConfigStorages, total, err
}
func (k8sConfigStorageService *K8sConfigStorageService)GetK8sConfigStoragePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
