
package k8sResource

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sResource"
    k8sResourceReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sResource/request"
)

type K8sResourceService struct {}
// CreateK8sResource 创建K8s资源记录
// Author [yourname](https://github.com/yourname)
func (krService *K8sResourceService) CreateK8sResource(ctx context.Context, kr *k8sResource.K8sResource) (err error) {
	err = global.GVA_DB.Create(kr).Error
	return err
}

// DeleteK8sResource 删除K8s资源记录
// Author [yourname](https://github.com/yourname)
func (krService *K8sResourceService)DeleteK8sResource(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&k8sResource.K8sResource{},"id = ?",ID).Error
	return err
}

// DeleteK8sResourceByIds 批量删除K8s资源记录
// Author [yourname](https://github.com/yourname)
func (krService *K8sResourceService)DeleteK8sResourceByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]k8sResource.K8sResource{},"id in ?",IDs).Error
	return err
}

// UpdateK8sResource 更新K8s资源记录
// Author [yourname](https://github.com/yourname)
func (krService *K8sResourceService)UpdateK8sResource(ctx context.Context, kr k8sResource.K8sResource) (err error) {
	err = global.GVA_DB.Model(&k8sResource.K8sResource{}).Where("id = ?",kr.ID).Updates(&kr).Error
	return err
}

// GetK8sResource 根据ID获取K8s资源记录
// Author [yourname](https://github.com/yourname)
func (krService *K8sResourceService)GetK8sResource(ctx context.Context, ID string) (kr k8sResource.K8sResource, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&kr).Error
	return
}
// GetK8sResourceInfoList 分页获取K8s资源记录
// Author [yourname](https://github.com/yourname)
func (krService *K8sResourceService)GetK8sResourceInfoList(ctx context.Context, info k8sResourceReq.K8sResourceSearch) (list []k8sResource.K8sResource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&k8sResource.K8sResource{})
    var krs []k8sResource.K8sResource
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["id"] = true
           orderMap["created_at"] = true
         	orderMap["name"] = true
         	orderMap["namespace"] = true
         	orderMap["type"] = true
         	orderMap["status"] = true
         	orderMap["last_modified"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&krs).Error
	return  krs, total, err
}
func (krService *K8sResourceService)GetK8sResourcePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
