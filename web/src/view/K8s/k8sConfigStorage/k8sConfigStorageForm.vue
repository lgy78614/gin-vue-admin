
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="自增主键:" prop="configId">
    <el-input v-model.number="formData.configId" :clearable="true" placeholder="请输入自增主键" />
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
        <el-form-item label="资源名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入资源名称" />
</el-form-item>
        <el-form-item label="类型: ConfigMap/Secret/PersistentVolumeClaim:" prop="type">
    <el-input v-model="formData.type" :clearable="true" placeholder="请输入类型: ConfigMap/Secret/PersistentVolumeClaim" />
</el-form-item>
        <el-form-item label="配置数据或Secret的键值对:" prop="data">
    <el-input v-model="formData.data" :clearable="true" placeholder="请输入配置数据或Secret的键值对" />
</el-form-item>
        <el-form-item label="PVC特有: 关联的Volume名称:" prop="volumeName">
    <el-input v-model="formData.volumeName" :clearable="true" placeholder="请输入PVC特有: 关联的Volume名称" />
</el-form-item>
        <el-form-item label="PVC特有: 存储类:" prop="storageClass">
    <el-input v-model="formData.storageClass" :clearable="true" placeholder="请输入PVC特有: 存储类" />
</el-form-item>
        <el-form-item label="PVC特有: 访问模式:" prop="accessModes">
    <el-input v-model="formData.accessModes" :clearable="true" placeholder="请输入PVC特有: 访问模式" />
</el-form-item>
        <el-form-item label="创建时间:" prop="creationTimestamp">
    <el-date-picker v-model="formData.creationTimestamp" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createK8sConfigStorage,
  updateK8sConfigStorage,
  findK8sConfigStorage
} from '@/api/K8s/k8sConfigStorage'

defineOptions({
    name: 'K8sConfigStorageForm'
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
            configId: undefined,
            clusterId: '',
            namespace: '',
            uid: '',
            name: '',
            type: '',
            data: '',
            volumeName: '',
            storageClass: '',
            accessModes: '',
            creationTimestamp: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findK8sConfigStorage({ ID: route.query.id })
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
               res = await createK8sConfigStorage(formData.value)
               break
             case 'update':
               res = await updateK8sConfigStorage(formData.value)
               break
             default:
               res = await createK8sConfigStorage(formData.value)
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
