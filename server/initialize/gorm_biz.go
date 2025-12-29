package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sResource"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(k8sResource.K8sResource{})
	if err != nil {
		return err
	}
	return nil
}
