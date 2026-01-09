
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="自增主键:" prop="namespaceId">
    <el-input v-model.number="formData.namespaceId" :clearable="true" placeholder="请输入自增主键" />
</el-form-item>
        <el-form-item label="所属集群ID:" prop="clusterId">
    <el-input v-model="formData.clusterId" :clearable="true" placeholder="请输入所属集群ID" />
</el-form-item>
        <el-form-item label="K8s中的唯一UID:" prop="uid">
    <el-input v-model="formData.uid" :clearable="true" placeholder="请输入K8s中的唯一UID" />
</el-form-item>
        <el-form-item label="命名空间名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入命名空间名称" />
</el-form-item>
        <el-form-item label="状态: Active/Terminating:" prop="status">
    <el-input v-model="formData.status" :clearable="true" placeholder="请输入状态: Active/Terminating" />
</el-form-item>
        <el-form-item label="标签(JSON格式):" prop="labels">
    <el-input v-model="formData.labels" :clearable="true" placeholder="请输入标签(JSON格式)" />
</el-form-item>
        <el-form-item label="K8s中的创建时间:" prop="creationTimestamp">
    <el-date-picker v-model="formData.creationTimestamp" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="计划删除时间:" prop="deletionTimestamp">
    <el-date-picker v-model="formData.deletionTimestamp" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createK8sNamespaces,
  updateK8sNamespaces,
  findK8sNamespaces
} from '@/api/K8s/k8sNamespaces'

defineOptions({
    name: 'K8sNamespacesForm'
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
            namespaceId: undefined,
            clusterId: '',
            uid: '',
            name: '',
            status: '',
            labels: '',
            creationTimestamp: new Date(),
            deletionTimestamp: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findK8sNamespaces({ ID: route.query.id })
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
               res = await createK8sNamespaces(formData.value)
               break
             case 'update':
               res = await updateK8sNamespaces(formData.value)
               break
             default:
               res = await createK8sNamespaces(formData.value)
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
