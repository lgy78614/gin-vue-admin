package K8s

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
	K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s/response"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClustersService struct{}

// CreateK8sClusters 创建k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService) CreateK8sClusters(ctx context.Context, k8sClusters *K8s.K8sClusters) (err error) {
	err = global.GVA_DB.Create(k8sClusters).Error
	return err
}

// DeleteK8sClusters 删除k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService) DeleteK8sClusters(ctx context.Context, clusterId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sClusters{}, "cluster_id = ?", clusterId).Error
	return err
}

// DeleteK8sClustersByIds 批量删除k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService) DeleteK8sClustersByIds(ctx context.Context, clusterIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sClusters{}, "cluster_id in ?", clusterIds).Error
	return err
}

// UpdateK8sClusters 更新k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService) UpdateK8sClusters(ctx context.Context, k8sClusters K8s.K8sClusters) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sClusters{}).Where("cluster_id = ?", k8sClusters.ClusterId).Updates(&k8sClusters).Error
	return err
}

// GetK8sClusters 根据clusterId获取k8sClusters表记录
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService) GetK8sClusters(ctx context.Context, clusterId string) (k8sClusters K8s.K8sClusters, err error) {
	err = global.GVA_DB.Where("cluster_id = ?", clusterId).First(&k8sClusters).Error
	return
}

// GetK8sClustersInfoList 分页获取k8sClusters表记录，并附带实时连接状态
func (k8sClustersService *K8sClustersService) GetK8sClustersInfoList(ctx context.Context, info K8sReq.K8sClustersSearch) (list []response.K8sClusterWithStatus, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&K8s.K8sClusters{})
	var clusters []K8s.K8sClusters

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&clusters).Error
	if err != nil {
		return
	}

	list = make([]response.K8sClusterWithStatus, len(clusters))
	for i, cluster := range clusters {
		status, err := k8sClustersService.getClusterStatus(ctx, &cluster)
		if err != nil {
			global.GVA_LOG.Error("获取集群状态失败:", zap.Error(err))
			continue
		}
		list[i] = *status
	}

	return list, total, nil
}

// func (k8sClustersService *K8sClustersService) enrichClustersStatus(ctx context.Context, clusters *[]K8s.K8sClusters) {
// 	if clusters == nil || len(*clusters) == 0 {
// 		return
// 	}

// 	// 使用协程并发获取状态
// 	var wg sync.WaitGroup
// 	for i := range *clusters {
// 		wg.Add(1)
// 		go func(cluster *K8s.K8sClusters) {
// 			defer wg.Done()

// 			status, err := k8sClustersService.getClusterStatus(ctx, cluster)
// 			if err != nil {
// 				global.GVA_LOG.Warn("获取集群状态失败",
// 					zap.String("cluster", *cluster.ClusterName),
// 					zap.Error(err))
// 			} else {
// 				// 更新返回对象的状态
// 				fmt.Print(*status)
// 			}
// 		}(&(*clusters)[i])
// 	}
// 	wg.Wait()
// }

func (k8sClustersService *K8sClustersService) getClusterStatus(ctx context.Context, cluster *K8s.K8sClusters) (*response.K8sClusterWithStatus, error) {
	if cluster == nil {
		return nil, errors.New("集群对象为空")
	}

	// 解析kubeconfig
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(*cluster.Kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("解析kubeconfig失败: %v", err)
	}

	// 创建超时上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 创建k8s客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("创建Kubernetes客户端失败: %v", err)
	}

	status := &response.K8sClusterWithStatus{
		K8sClusters:     *cluster,
		HealthStatus:    0,
		PodCount:        0,
		NodeCount:       0,
		ActiveNodeCount: 0,
		ServerVersion:   "",
	}

	// 1. 获取服务器版本
	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("获取服务器版本失败: %v", err)
	}
	status.ServerVersion = version.String()

	// 2. 获取节点信息
	nodes, err := clientset.CoreV1().Nodes().List(timeoutCtx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取节点列表失败: %v", err)
	}

	status.NodeCount = len(nodes.Items)
	activeNodes := 0
	for _, node := range nodes.Items {
		for _, condition := range node.Status.Conditions {
			if condition.Type == corev1.NodeReady && condition.Status == corev1.ConditionTrue {
				activeNodes++
				break
			}
		}
	}
	status.ActiveNodeCount = activeNodes

	// 3. 获取Pod总数
	allPods, err := clientset.CoreV1().Pods("").List(timeoutCtx, metav1.ListOptions{})
	if err != nil {
		// 如果获取全部Pod失败，尝试分命名空间获取
		global.GVA_LOG.Warn("获取全部Pod失败，尝试分命名空间获取", zap.Error(err))
		status.PodCount = 0
	} else {
		status.PodCount = len(allPods.Items)
	}

	// 4. 判断集群健康状态
	if activeNodes == 0 {
		status.HealthStatus = 3 // 异常
	} else if float64(activeNodes)/float64(len(nodes.Items)) < 0.8 {
		status.HealthStatus = 2 // 警告
	} else {
		status.HealthStatus = 1 // 健康
	}

	return status, nil
}

func (k8sClustersService *K8sClustersService) GetK8sClustersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetK8sClusterInfo 根据k8s_clusters中的kubeconfig字段，查询K8s集群及资源信息
// Author [yourname](https://github.com/yourname)
func (k8sClustersService *K8sClustersService) GetK8sClusterInfo(ctx context.Context) (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&K8s.K8sClusters{})
	return db.Error
}
