
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="集群唯一标识:" prop="clusterId">
    <el-input v-model="formData.clusterId" :clearable="true" placeholder="请输入集群唯一标识" />
</el-form-item>
        <el-form-item label="集群显示名称:" prop="clusterName">
    <el-input v-model="formData.clusterName" :clearable="true" placeholder="请输入集群显示名称" />
</el-form-item>
        <el-form-item label="环境: production/staging/development:" prop="environment">
    <el-input v-model="formData.environment" :clearable="true" placeholder="请输入环境: production/staging/development" />
</el-form-item>
        <el-form-item label="Kubernetes版本:" prop="k8sVersion">
    <el-input v-model="formData.k8sVersion" :clearable="true" placeholder="请输入Kubernetes版本" />
</el-form-item>
        <el-form-item label="云厂商: AWS/Azure/GCP/on-premise:" prop="cloudProvider">
    <el-input v-model="formData.cloudProvider" :clearable="true" placeholder="请输入云厂商: AWS/Azure/GCP/on-premise" />
</el-form-item>
        <el-form-item label="区域:" prop="region">
    <el-input v-model="formData.region" :clearable="true" placeholder="请输入区域" />
</el-form-item>
        <el-form-item label="记录创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="记录更新时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createK8sClusters,
  updateK8sClusters,
  findK8sClusters
} from '@/api/K8s/k8sClusters'

defineOptions({
    name: 'K8sClustersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            clusterId: '',
            clusterName: '',
            environment: '',
            k8sVersion: '',
            cloudProvider: '',
            region: '',
            createdAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findK8sClusters({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createK8sClusters(formData.value)
               break
             case 'update':
               res = await updateK8sClusters(formData.value)
               break
             default:
               res = await createK8sClusters(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
