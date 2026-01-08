package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/K8s"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(K8s.K8sClusters{}, K8s.K8sNodes{}, K8s.K8sPods{}, K8s.K8sServices{})
	if err != nil {
		return err
	}
	return nil
}
