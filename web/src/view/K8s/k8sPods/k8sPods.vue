
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
            <el-form-item label="所属集群ID" prop="clusterId">
  <el-input v-model="searchInfo.clusterId" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="所属命名空间" prop="namespace">
  <el-input v-model="searchInfo.namespace" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="Pod名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="podId"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="自增主键" prop="podId" width="120" />

            <el-table-column align="left" label="所属集群ID" prop="clusterId" width="120" />

            <el-table-column align="left" label="所属命名空间" prop="namespace" width="120" />

            <el-table-column align="left" label="K8s UID" prop="uid" width="120" />

            <el-table-column align="left" label="Pod名称" prop="name" width="120" />

            <el-table-column align="left" label="调度到的节点名称" prop="nodeName" width="120" />

            <el-table-column align="left" label="所属工作负载的UID" prop="workloadUid" width="120" />

            <el-table-column align="left" label="阶段: Pending/Running/Succeeded/Failed/Unknown" prop="phase" width="120" />

            <el-table-column align="left" label="Pod IP地址" prop="podIp" width="120" />

            <el-table-column align="left" label="主机IP地址" prop="hostIp" width="120" />

            <el-table-column align="left" label="标签" prop="labels" width="120" />

            <el-table-column align="left" label="创建时间" prop="creationTimestamp" width="180">
   <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
</el-table-column>
            <el-table-column align="left" label="启动时间(进入Running阶段)" prop="startTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.startTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="重启次数" prop="restartCount" width="120" />

            <el-table-column align="left" label="状态条件" prop="statusConditions" width="120" />

            <el-table-column align="left" label="上一次终止的容器状态" prop="lastTerminatedState" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateK8sPodsFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="自增主键">
    {{ detailForm.podId }}
</el-descriptions-item>
                    <el-descriptions-item label="所属集群ID">
    {{ detailForm.clusterId }}
</el-descriptions-item>
                    <el-descriptions-item label="所属命名空间">
    {{ detailForm.namespace }}
</el-descriptions-item>
                    <el-descriptions-item label="K8s UID">
    {{ detailForm.uid }}
</el-descriptions-item>
                    <el-descriptions-item label="Pod名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="调度到的节点名称">
    {{ detailForm.nodeName }}
</el-descriptions-item>
                    <el-descriptions-item label="所属工作负载的UID">
    {{ detailForm.workloadUid }}
</el-descriptions-item>
                    <el-descriptions-item label="阶段: Pending/Running/Succeeded/Failed/Unknown">
    {{ detailForm.phase }}
</el-descriptions-item>
                    <el-descriptions-item label="Pod IP地址">
    {{ detailForm.podIp }}
</el-descriptions-item>
                    <el-descriptions-item label="主机IP地址">
    {{ detailForm.hostIp }}
</el-descriptions-item>
                    <el-descriptions-item label="标签">
    {{ detailForm.labels }}
</el-descriptions-item>
                    <el-descriptions-item label="创建时间">
    {{ detailForm.creationTimestamp }}
</el-descriptions-item>
                    <el-descriptions-item label="启动时间(进入Running阶段)">
    {{ detailForm.startTime }}
</el-descriptions-item>
                    <el-descriptions-item label="重启次数">
    {{ detailForm.restartCount }}
</el-descriptions-item>
                    <el-descriptions-item label="状态条件">
    {{ detailForm.statusConditions }}
</el-descriptions-item>
                    <el-descriptions-item label="上一次终止的容器状态">
    {{ detailForm.lastTerminatedState }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createK8sPods,
  deleteK8sPods,
  deleteK8sPodsByIds,
  updateK8sPods,
  findK8sPods,
  getK8sPodsList
} from '@/api/K8s/k8sPods'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'K8sPods'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getK8sPodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteK8sPodsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const podIds = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          podIds.push(item.podId)
        })
      const res = await deleteK8sPodsByIds({ podIds })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === podIds.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateK8sPodsFunc = async(row) => {
    const res = await findK8sPods({ podId: row.podId })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteK8sPodsFunc = async (row) => {
    const res = await deleteK8sPods({ podId: row.podId })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findK8sPods({ podId: row.podId })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
