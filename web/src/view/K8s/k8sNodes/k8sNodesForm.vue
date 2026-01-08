
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="节点唯一标识:" prop="nodeId">
    <el-input v-model="formData.nodeId" :clearable="true" placeholder="请输入节点唯一标识" />
</el-form-item>
        <el-form-item label="所属集群ID:" prop="clusterId">
    <el-input v-model="formData.clusterId" :clearable="true" placeholder="请输入所属集群ID" />
</el-form-item>
        <el-form-item label="节点名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入节点名称" />
</el-form-item>
        <el-form-item label="状态: Ready/NotReady/Unknown:" prop="status">
    <el-input v-model="formData.status" :clearable="true" placeholder="请输入状态: Ready/NotReady/Unknown" />
</el-form-item>
        <el-form-item label="角色: control-plane/worker:" prop="role">
    <el-input v-model="formData.role" :clearable="true" placeholder="请输入角色: control-plane/worker" />
</el-form-item>
        <el-form-item label="操作系统镜像:" prop="osImage">
    <el-input v-model="formData.osImage" :clearable="true" placeholder="请输入操作系统镜像" />
</el-form-item>
        <el-form-item label="内核版本:" prop="kernelVersion">
    <el-input v-model="formData.kernelVersion" :clearable="true" placeholder="请输入内核版本" />
</el-form-item>
        <el-form-item label="架构: amd64/arm64:" prop="architecture">
    <el-input v-model="formData.architecture" :clearable="true" placeholder="请输入架构: amd64/arm64" />
</el-form-item>
        <el-form-item label="CPU核心数(可分配总量):" prop="cpuCapacity">
    <el-input-number v-model="formData.cpuCapacity" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="内存容量(字节):" prop="memoryCapacityBytes">
    <el-input v-model.number="formData.memoryCapacityBytes" :clearable="true" placeholder="请输入内存容量(字节)" />
</el-form-item>
        <el-form-item label="Pod容量上限:" prop="podCapacity">
    <el-input v-model.number="formData.podCapacity" :clearable="true" placeholder="请输入Pod容量上限" />
</el-form-item>
        <el-form-item label="K8s中的创建时间:" prop="creationTimestamp">
    <el-date-picker v-model="formData.creationTimestamp" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="最后一次心跳/同步时间:" prop="lastHeartbeat">
    <el-date-picker v-model="formData.lastHeartbeat" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createK8sNodes,
  updateK8sNodes,
  findK8sNodes
} from '@/api/K8s/k8sNodes'

defineOptions({
    name: 'K8sNodesForm'
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
            nodeId: '',
            clusterId: '',
            name: '',
            status: '',
            role: '',
            osImage: '',
            kernelVersion: '',
            architecture: '',
            cpuCapacity: 0,
            memoryCapacityBytes: undefined,
            podCapacity: undefined,
            creationTimestamp: new Date(),
            lastHeartbeat: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findK8sNodes({ ID: route.query.id })
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
               res = await createK8sNodes(formData.value)
               break
             case 'update':
               res = await updateK8sNodes(formData.value)
               break
             default:
               res = await createK8sNodes(formData.value)
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
