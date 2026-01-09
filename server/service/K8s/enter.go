package K8s

type ServiceGroup struct {
	K8sClustersService
	K8sNodesService
	K8sPodsService
	K8sServicesService
	K8sNamespacesService
	K8sWorkloadsService
	K8sConfigStorageService
}
