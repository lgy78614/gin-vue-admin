
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="自增主键:" prop="workloadId">
    <el-input v-model.number="formData.workloadId" :clearable="true" placeholder="请输入自增主键" />
</el-form-item>
        <el-form-item label="所属集群ID:" prop="clusterId">
    <el-input v-model="formData.clusterId" :clearable="true" placeholder="请输入所属集群ID" />
</el-form-item>
        <el-form-item label="所属命名空间:" prop="namespace">
    <el-input v-model="formData.namespace" :clearable="true" placeholder="请输入所属命名空间" />
</el-form-item>
        <el-form-item label="K8s UID:" prop="uid">
    <el-input v-model="formData.uid" :clearable="true" placeholder="请输入K8s UID" />
</el-form-item>
        <el-form-item label="工作负载名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入工作负载名称" />
</el-form-item>
        <el-form-item label="类型: Deployment/StatefulSet/DaemonSet/Job/CronJob:" prop="type">
    <el-input v-model="formData.type" :clearable="true" placeholder="请输入类型: Deployment/StatefulSet/DaemonSet/Job/CronJob" />
</el-form-item>
        <el-form-item label="期望副本数:" prop="replicasDesired">
    <el-input v-model.number="formData.replicasDesired" :clearable="true" placeholder="请输入期望副本数" />
</el-form-item>
        <el-form-item label="当前副本数:" prop="replicasCurrent">
    <el-input v-model.number="formData.replicasCurrent" :clearable="true" placeholder="请输入当前副本数" />
</el-form-item>
        <el-form-item label="就绪副本数:" prop="replicasReady">
    <el-input v-model.number="formData.replicasReady" :clearable="true" placeholder="请输入就绪副本数" />
</el-form-item>
        <el-form-item label="标签:" prop="labels">
    <el-input v-model="formData.labels" :clearable="true" placeholder="请输入标签" />
</el-form-item>
        <el-form-item label="注解:" prop="annotations">
    <el-input v-model="formData.annotations" :clearable="true" placeholder="请输入注解" />
</el-form-item>
        <el-form-item label="标签选择器(JSON格式):" prop="selector">
    <el-input v-model="formData.selector" :clearable="true" placeholder="请输入标签选择器(JSON格式)" />
</el-form-item>
        <el-form-item label="创建时间:" prop="creationTimestamp">
    <el-date-picker v-model="formData.creationTimestamp" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="最后同步时间:" prop="lastSyncTimestamp">
    <el-date-picker v-model="formData.lastSyncTimestamp" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="状态条件(JSON数组):" prop="statusConditions">
    <el-input v-model="formData.statusConditions" :clearable="true" placeholder="请输入状态条件(JSON数组)" />
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
  createK8sWorkloads,
  updateK8sWorkloads,
  findK8sWorkloads
} from '@/api/K8s/k8sWorkloads'

defineOptions({
    name: 'K8sWorkloadsForm'
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
            workloadId: undefined,
            clusterId: '',
            namespace: '',
            uid: '',
            name: '',
            type: '',
            replicasDesired: undefined,
            replicasCurrent: undefined,
            replicasReady: undefined,
            labels: '',
            annotations: '',
            selector: '',
            creationTimestamp: new Date(),
            lastSyncTimestamp: new Date(),
            statusConditions: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findK8sWorkloads({ ID: route.query.id })
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
               res = await createK8sWorkloads(formData.value)
               break
             case 'update':
               res = await updateK8sWorkloads(formData.value)
               break
             default:
               res = await createK8sWorkloads(formData.value)
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
