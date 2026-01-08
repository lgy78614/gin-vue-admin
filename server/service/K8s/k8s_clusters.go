
package K8s

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
    K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
)

type K8sClustersService struct {}
// CreateK8sClusters 创建k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService) CreateK8sClusters(ctx context.Context, k8sClusters *K8s.K8sClusters) (err error) {
	err = global.GVA_DB.Create(k8sClusters).Error
	return err
}

// DeleteK8sClusters 删除k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService)DeleteK8sClusters(ctx context.Context, clusterId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sClusters{},"cluster_id = ?",clusterId).Error
	return err
}

// DeleteK8sClustersByIds 批量删除k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService)DeleteK8sClustersByIds(ctx context.Context, clusterIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sClusters{},"cluster_id in ?",clusterIds).Error
	return err
}

// UpdateK8sClusters 更新k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService)UpdateK8sClusters(ctx context.Context, k8sClusters K8s.K8sClusters) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sClusters{}).Where("cluster_id = ?",k8sClusters.ClusterId).Updates(&k8sClusters).Error
	return err
}

// GetK8sClusters 根据clusterId获取k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService)GetK8sClusters(ctx context.Context, clusterId string) (k8sClusters K8s.K8sClusters, err error) {
	err = global.GVA_DB.Where("cluster_id = ?", clusterId).First(&k8sClusters).Error
	return
}
// GetK8sClustersInfoList 分页获取k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService)GetK8sClustersInfoList(ctx context.Context, info K8sReq.K8sClustersSearch) (list []K8s.K8sClusters, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&K8s.K8sClusters{})
    var k8sClusterss []K8s.K8sClusters
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&k8sClusterss).Error
	return  k8sClusterss, total, err
}
func (k8sClustersService *K8sClustersService)GetK8sClustersPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
