<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item label="名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="代码" prop="code">
         <el-input v-model="searchInfo.code" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="倍数" prop="multiple">
            
             <el-input v-model.number="searchInfo.multiple" placeholder="搜索条件" />

        </el-form-item>
            <el-form-item label="状态" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="代码" prop="code" width="120" />
        <el-table-column align="left" label="倍数" prop="multiple" width="120" />
        <el-table-column align="left" label="保证金(%)" prop="bond" width="120" />
        <el-table-column align="left" label="止赢点位" prop="point_success" width="120" />
        <el-table-column align="left" label="止赢点位赔付价格" prop="point_success_price" width="120" />
        <el-table-column align="left" label="止损点位" prop="point_fail" width="120" />
        <el-table-column align="left" label="止损点位赔付价格" prop="point_fail_price" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateSymbolFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="代码:"  prop="code" >
              <el-input v-model="formData.code" :clearable="true"  placeholder="请输入代码" />
            </el-form-item>
            <el-form-item label="倍数:"  prop="multiple" >
              <el-input v-model.number="formData.multiple" :clearable="true" placeholder="请输入倍数" />
            </el-form-item>
            <el-form-item label="保证金(%):"  prop="bond" >
              <el-input v-model.number="formData.bond" :clearable="true" placeholder="请输入保证金(%)" />
            </el-form-item>
            <el-form-item label="止赢点位:"  prop="point_success" >
              <el-input v-model.number="formData.point_success" :clearable="true" placeholder="请输入止赢点位" />
            </el-form-item>
            <el-form-item label="止赢价格:"  prop="point_success_price" >
              <el-input v-model.number="formData.point_success_price" :clearable="true" placeholder="请输入止赢点位赔付价格" />
            </el-form-item>
            <el-form-item label="止损点位:"  prop="point_fail" >
              <el-input v-model.number="formData.point_fail" :clearable="true" placeholder="请输入止损点位" />
            </el-form-item>
            <el-form-item label="止损价格:"  prop="point_fail_price" >
              <el-input v-model.number="formData.point_fail_price" :clearable="true" placeholder="请输入止损点位赔付价格" />
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
              <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="开始时间:"  prop="start_at" >
              <el-time-select v-model="formData.start_at" :picker-options="{ start: '07:30', step: '00:15', end: '20:30' }" placeholder="选择时间"></el-time-select>
            </el-form-item>
            <el-form-item label="结束时间:"  prop="end_at" >
              <el-time-select v-model="formData.end_at" :picker-options="{ start: '07:30', step: '00:15', end: '20:30' }" placeholder="选择时间"></el-time-select>
            </el-form-item>
            <el-form-item label="特殊时间:"  prop="days" >
              <el-input type="textarea" v-model="formData.days" rows="3"></el-input>
              <el-alert title="格式为【2023-12-09 09:30:00~2023-12-10 18:30:00】，多个换行" type=info :closable="true"  :show-icon="false"
                        effect="light">
              </el-alert>
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="代码">
                        {{ formData.code }}
                </el-descriptions-item>
                <el-descriptions-item label="倍数">
                        {{ formData.multiple }}
                </el-descriptions-item>
                <el-descriptions-item label="保证金(%)">
                        {{ formData.bond }}
                </el-descriptions-item>
                <el-descriptions-item label="止赢点位">
                        {{ formData.point_success }}
                </el-descriptions-item>
                <el-descriptions-item label="止赢点位赔付价格">
                        {{ formData.point_success_price }}
                </el-descriptions-item>
                <el-descriptions-item label="止损点位">
                        {{ formData.point_fail }}
                </el-descriptions-item>
                <el-descriptions-item label="止损点位赔付价格">
                        {{ formData.point_fail_price }}
                </el-descriptions-item>
                <el-descriptions-item label="状态">
                    {{ formatBoolean(formData.status) }}
                </el-descriptions-item>
          <el-descriptions-item label="下单开始时间">
            <el-time-select v-model="formData.start_at" disabled :picker-options="{ start: '07:30', step: '00:15', end: '20:30' }" placeholder="选择时间"></el-time-select>
          </el-descriptions-item>
          <el-descriptions-item label="下单结束时间">
            <el-time-select v-model="formData.end_at" disabled :picker-options="{ start: '07:30', step: '00:15', end: '20:30' }" placeholder="选择时间"></el-time-select>
          </el-descriptions-item>
          <el-descriptions-item label="特殊时间">
            <el-input type="textarea" v-model="formData.days" disabled rows="3"></el-input>
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createSymbol,
  deleteSymbol,
  deleteSymbolByIds,
  updateSymbol,
  findSymbol,
  getSymbolList
} from '@/api/symbol'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Symbol'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        name: '',
        code: '',
        multiple: 0,
        bond: 0,
        point_success: 0,
        point_success_price: 0,
        point_fail: 0,
        point_fail_price: 0,
        status: false,
        })


// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               code : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               multiple : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               bond : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
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
    pageSize.value = 10
    if (searchInfo.value.status === ""){
        searchInfo.value.status=null
    }
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
  const table = await getSymbolList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteSymbolFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteSymbolByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateSymbolFunc = async(row) => {
    const res = await findSymbol({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.resb
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSymbolFunc = async (row) => {
    const res = await deleteSymbol({ ID: row.ID })
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


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findSymbol({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.resb
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          name: '',
          code: '',
          multiple: 0,
          bond: 0,
          point_success: 0,
          point_success_price: 0,
          point_fail: 0,
          point_fail_price: 0,
          status: false,
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        name: '',
        code: '',
        multiple: 0,
        bond: 0,
        point_success: 0,
        point_success_price: 0,
        point_fail: 0,
        point_fail_price: 0,
        status: false,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createSymbol(formData.value)
                  break
                case 'update':
                  res = await updateSymbol(formData.value)
                  break
                default:
                  res = await createSymbol(formData.value)
                  break
              }
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

</script>

<style>

</style>
