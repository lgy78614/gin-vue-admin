package K8s

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
	K8sReq "github.com/flipped-aurora/gin-vue-admin/server/model/K8s/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s/response"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sNodesService struct{}

// CreateK8sNodes 创建k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService) CreateK8sNodes(ctx context.Context, k8sNodes *K8s.K8sNodes) (err error) {
	err = global.GVA_DB.Create(k8sNodes).Error
	return err
}

// DeleteK8sNodes 删除k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService) DeleteK8sNodes(ctx context.Context, nodeId string) (err error) {
	err = global.GVA_DB.Delete(&K8s.K8sNodes{}, "node_id = ?", nodeId).Error
	return err
}

// DeleteK8sNodesByIds 批量删除k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService) DeleteK8sNodesByIds(ctx context.Context, nodeIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]K8s.K8sNodes{}, "node_id in ?", nodeIds).Error
	return err
}

// UpdateK8sNodes 更新k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService) UpdateK8sNodes(ctx context.Context, k8sNodes K8s.K8sNodes) (err error) {
	err = global.GVA_DB.Model(&K8s.K8sNodes{}).Where("node_id = ?", k8sNodes.NodeId).Updates(&k8sNodes).Error
	return err
}

// GetK8sNodes 根据nodeId获取k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService) GetK8sNodes(ctx context.Context, nodeId string) (k8sNodes K8s.K8sNodes, err error) {
	err = global.GVA_DB.Where("node_id = ?", nodeId).First(&k8sNodes).Error
	return
}

// GetK8sNodesInfoList 分页获取k8sNodes表记录
// Author [yourname](https://github.com/yourname)
func (k8sNodesService *K8sNodesService) GetK8sNodesInfoList(ctx context.Context, info K8sReq.K8sNodesSearch) (list []K8s.K8sNodes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&K8s.K8sNodes{})
	var k8sNodess []K8s.K8sNodes
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&k8sNodess).Error

	return k8sNodess, total, err
}

// GetK8sNodesDetailsFromCluster 从指定集群的 API Server 获取节点详细信息
func (k8sNodesService *K8sNodesService) GetK8sNodesDetailsFromCluster(ctx context.Context, clusterId string) ([]response.K8sNodeDetail, error) {
	// 1. 从数据库获取集群 kubeconfig
	var cluster K8s.K8sClusters
	if err := global.GVA_DB.Where("cluster_id = ?", clusterId).First(&cluster).Error; err != nil {
		return nil, fmt.Errorf("集群不存在: %v", err)
	}

	if cluster.Kubeconfig == nil || *cluster.Kubeconfig == "" {
		return nil, fmt.Errorf("集群 kubeconfig 为空")
	}

	// 2. 构建 clientset
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(*cluster.Kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("解析 kubeconfig 失败: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("创建 Kubernetes 客户端失败: %v", err)
	}

	// 3. 获取节点列表
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	nodeList, err := clientset.CoreV1().Nodes().List(timeoutCtx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取节点列表失败: %v", err)
	}

	// 4. 转换为响应结构
	var result []response.K8sNodeDetail
	for _, node := range nodeList.Items {
		// 提取角色（通过 label 判断）
		roles := []string{}
		labels := node.GetLabels()
		if labels["node-role.kubernetes.io/control-plane"] == "true" {
			roles = append(roles, "control-plane")
		}
		if labels["node-role.kubernetes.io/master"] == "true" {
			roles = append(roles, "master")
		}
		if len(roles) == 0 {
			roles = append(roles, "worker")
		}

		// 提取状态
		status := "Unknown"
		for _, condition := range node.Status.Conditions {
			if condition.Type == corev1.NodeReady {
				status = string(condition.Status)
				break
			}
		}

		result = append(result, response.K8sNodeDetail{
			ClusterId:      clusterId,
			Name:           node.Name,
			Status:         status,
			Roles:          roles,
			OsImage:        node.Status.NodeInfo.OSImage,
			KernelVersion:  node.Status.NodeInfo.KernelVersion,
			Architecture:   node.Status.NodeInfo.Architecture,
			CpuCapacity:    node.Status.Capacity.Cpu().String(),
			MemoryCapacity: node.Status.Capacity.Memory().String(),
			PodCapacity:    node.Status.Capacity.Pods().String(),
		})
	}

	return result, nil
}
func (k8sNodesService *K8sNodesService) GetK8sNodesPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
