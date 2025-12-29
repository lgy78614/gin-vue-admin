
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="资源名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入资源名称" />
</el-form-item>
        <el-form-item label="命名空间:" prop="namespace">
    <el-input v-model="formData.namespace" :clearable="false" placeholder="请输入命名空间" />
</el-form-item>
        <el-form-item label="资源类型:" prop="type">
    <el-select v-model="formData.type" placeholder="请选择资源类型" style="width:100%" filterable :clearable="false">
       <el-option v-for="item in ['pod','deployment','service','configmap','secret','ingress','statefulset','daemonset','job','cronjob']" :key="item" :label="item" :value="item" />
    </el-select>
</el-form-item>
        <el-form-item label="资源标签:" prop="labels">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.labels 后端会按照json的类型进行存取
    {{ formData.labels }}
</el-form-item>
        <el-form-item label="资源注解:" prop="annotations">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.annotations 后端会按照json的类型进行存取
    {{ formData.annotations }}
</el-form-item>
        <el-form-item label="资源规格:" prop="spec">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.spec 后端会按照json的类型进行存取
    {{ formData.spec }}
</el-form-item>
        <el-form-item label="资源配置:" prop="config">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.config 后端会按照json的类型进行存取
    {{ formData.config }}
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
  createK8sResource,
  updateK8sResource,
  findK8sResource
} from '@/api/k8sResource/k8sManager'

defineOptions({
    name: 'K8sResourceForm'
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
            name: '',
            namespace: '',
            type: null,
            labels: {},
            annotations: {},
            spec: {},
            config: {},
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               namespace : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findK8sResource({ ID: route.query.id })
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
               res = await createK8sResource(formData.value)
               break
             case 'update':
               res = await updateK8sResource(formData.value)
               break
             default:
               res = await createK8sResource(formData.value)
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
