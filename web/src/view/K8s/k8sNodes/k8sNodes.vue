
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

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
        row-key="nodeId"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="节点唯一标识" prop="nodeId" width="120" />

            <el-table-column align="left" label="所属集群ID" prop="clusterId" width="120" />

            <el-table-column align="left" label="节点名称" prop="name" width="120" />

            <el-table-column align="left" label="状态: Ready/NotReady/Unknown" prop="status" width="120" />

            <el-table-column align="left" label="角色: control-plane/worker" prop="role" width="120" />

            <el-table-column align="left" label="操作系统镜像" prop="osImage" width="120" />

            <el-table-column align="left" label="内核版本" prop="kernelVersion" width="120" />

            <el-table-column align="left" label="架构: amd64/arm64" prop="architecture" width="120" />

            <el-table-column align="left" label="CPU核心数(可分配总量)" prop="cpuCapacity" width="120" />

            <el-table-column align="left" label="内存容量(字节)" prop="memoryCapacityBytes" width="120" />

            <el-table-column align="left" label="Pod容量上限" prop="podCapacity" width="120" />

            <el-table-column align="left" label="K8s中的创建时间" prop="creationTimestamp" width="180">
   <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
</el-table-column>
            <el-table-column align="left" label="最后一次心跳/同步时间" prop="lastHeartbeat" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastHeartbeat) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateK8sNodesFunc(scope.row)">编辑</el-button>
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="节点唯一标识">
    {{ detailForm.nodeId }}
</el-descriptions-item>
                    <el-descriptions-item label="所属集群ID">
    {{ detailForm.clusterId }}
</el-descriptions-item>
                    <el-descriptions-item label="节点名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="状态: Ready/NotReady/Unknown">
    {{ detailForm.status }}
</el-descriptions-item>
                    <el-descriptions-item label="角色: control-plane/worker">
    {{ detailForm.role }}
</el-descriptions-item>
                    <el-descriptions-item label="操作系统镜像">
    {{ detailForm.osImage }}
</el-descriptions-item>
                    <el-descriptions-item label="内核版本">
    {{ detailForm.kernelVersion }}
</el-descriptions-item>
                    <el-descriptions-item label="架构: amd64/arm64">
    {{ detailForm.architecture }}
</el-descriptions-item>
                    <el-descriptions-item label="CPU核心数(可分配总量)">
    {{ detailForm.cpuCapacity }}
</el-descriptions-item>
                    <el-descriptions-item label="内存容量(字节)">
    {{ detailForm.memoryCapacityBytes }}
</el-descriptions-item>
                    <el-descriptions-item label="Pod容量上限">
    {{ detailForm.podCapacity }}
</el-descriptions-item>
                    <el-descriptions-item label="K8s中的创建时间">
    {{ detailForm.creationTimestamp }}
</el-descriptions-item>
                    <el-descriptions-item label="最后一次心跳/同步时间">
    {{ detailForm.lastHeartbeat }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createK8sNodes,
  deleteK8sNodes,
  deleteK8sNodesByIds,
  updateK8sNodes,
  findK8sNodes,
  getK8sNodesList
} from '@/api/K8s/k8sNodes'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'K8sNodes'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
  const table = await getK8sNodesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteK8sNodesFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const nodeIds = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          nodeIds.push(item.nodeId)
        })
      const res = await deleteK8sNodesByIds({ nodeIds })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === nodeIds.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateK8sNodesFunc = async(row) => {
    const res = await findK8sNodes({ nodeId: row.nodeId })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteK8sNodesFunc = async (row) => {
    const res = await deleteK8sNodes({ nodeId: row.nodeId })
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
  const res = await findK8sNodes({ nodeId: row.nodeId })
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
