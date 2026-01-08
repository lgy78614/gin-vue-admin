
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="自增主键:" prop="podId">
    <el-input v-model.number="formData.podId" :clearable="true" placeholder="请输入自增主键" />
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
        <el-form-item label="Pod名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入Pod名称" />
</el-form-item>
        <el-form-item label="调度到的节点名称:" prop="nodeName">
    <el-input v-model="formData.nodeName" :clearable="true" placeholder="请输入调度到的节点名称" />
</el-form-item>
        <el-form-item label="所属工作负载的UID:" prop="workloadUid">
    <el-input v-model="formData.workloadUid" :clearable="true" placeholder="请输入所属工作负载的UID" />
</el-form-item>
        <el-form-item label="阶段: Pending/Running/Succeeded/Failed/Unknown:" prop="phase">
    <el-input v-model="formData.phase" :clearable="true" placeholder="请输入阶段: Pending/Running/Succeeded/Failed/Unknown" />
</el-form-item>
        <el-form-item label="Pod IP地址:" prop="podIp">
    <el-input v-model="formData.podIp" :clearable="true" placeholder="请输入Pod IP地址" />
</el-form-item>
        <el-form-item label="主机IP地址:" prop="hostIp">
    <el-input v-model="formData.hostIp" :clearable="true" placeholder="请输入主机IP地址" />
</el-form-item>
        <el-form-item label="标签:" prop="labels">
    <el-input v-model="formData.labels" :clearable="true" placeholder="请输入标签" />
</el-form-item>
        <el-form-item label="创建时间:" prop="creationTimestamp">
    <el-date-picker v-model="formData.creationTimestamp" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="启动时间(进入Running阶段):" prop="startTime">
    <el-date-picker v-model="formData.startTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="重启次数:" prop="restartCount">
    <el-input v-model.number="formData.restartCount" :clearable="true" placeholder="请输入重启次数" />
</el-form-item>
        <el-form-item label="状态条件:" prop="statusConditions">
    <el-input v-model="formData.statusConditions" :clearable="true" placeholder="请输入状态条件" />
</el-form-item>
        <el-form-item label="上一次终止的容器状态:" prop="lastTerminatedState">
    <el-input v-model="formData.lastTerminatedState" :clearable="true" placeholder="请输入上一次终止的容器状态" />
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
  createK8sPods,
  updateK8sPods,
  findK8sPods
} from '@/api/K8s/k8sPods'

defineOptions({
    name: 'K8sPodsForm'
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
            podId: undefined,
            clusterId: '',
            namespace: '',
            uid: '',
            name: '',
            nodeName: '',
            workloadUid: '',
            phase: '',
            podIp: '',
            hostIp: '',
            labels: '',
            creationTimestamp: new Date(),
            startTime: new Date(),
            restartCount: undefined,
            statusConditions: '',
            lastTerminatedState: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findK8sPods({ ID: route.query.id })
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
               res = await createK8sPods(formData.value)
               break
             case 'update':
               res = await updateK8sPods(formData.value)
               break
             default:
               res = await createK8sPods(formData.value)
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
