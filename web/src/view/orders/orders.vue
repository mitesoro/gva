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
        <el-form-item label="账户" prop="account_id">
            
             <el-input v-model.number="searchInfo.account_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="订单号" prop="order_no">
         <el-input v-model="searchInfo.order_no" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="用户编号" prop="user_id">
          <el-select v-model="searchInfo.user_id" clearable placeholder="请选择" @clear="()=>{searchInfo.user_id=undefined}">
            <el-option v-for="(item,key) in userOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="订单状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
            <el-option v-for="(item,key) in orderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="类型" prop="direction">
          <el-select v-model="searchInfo.direction" placeholder="搜索条件">
            <el-option label="买多" value="1"></el-option>
            <el-option label="卖空" value="2"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="价格" prop="price">
            
             <el-input v-model.number="searchInfo.price" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
      <el-divider></el-divider> <!-- 添加分割线 -->
      <el-tag type="danger" style="margin-top: 10px;font-size: large;font-weight: bold;"> <!-- 用 el-tag 创建有背景的方块 -->
        盈亏比：{{success}}：{{fail}}
      </el-tag>
      <el-divider></el-divider> <!-- 添加分割线 -->
    </div>
    <div class="gva-table-box">
<!--        <div class="gva-btn-list">-->
<!--            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
<!--            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">-->
<!--            <p>确定要删除吗？</p>-->
<!--            <div style="text-align: right; margin-top: 8px;">-->
<!--                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>-->
<!--                <el-button type="primary" @click="onDelete">确定</el-button>-->
<!--            </div>-->
<!--            <template #reference>-->
<!--                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>-->
<!--            </template>-->
<!--            </el-popover>-->
<!--        </div>-->
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="开仓日期" width="120">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
          <el-table-column align="left" label="用户昵称" prop="user_id" width="120">
            <template #default="scope">
              {{ scope.row.User.nickname }}
            </template>
          </el-table-column>

          <el-table-column align="left" label="手机号" prop="user_id" width="120">
            <template #default="scope">
              {{ scope.row.User.phone }}
            </template>
          </el-table-column>
<!--        <el-table-column align="left" label="账户" prop="account_id" width="120" />-->
<!--        <el-table-column align="left" label="订单号" prop="order_no" width="120" />-->
        <el-table-column align="left" label="类型"  width="80" >
          <template #default="scope">
            <div>
              <el-tag effect="dark" :type="formatTagType(scope.row.direction)">{{ formatDirection(scope.row.direction) }}</el-tag>
            </div>
          </template>
        </el-table-column>
          <el-table-column align="left" label="合约" prop="symbol_id" width="80" />
          <el-table-column align="left" label="订单状态" prop="status" width="80" >
            <template #default="scope">
              <div>
                <el-tag effect="dark" >{{ formatOrderStatus(scope.row.status) }}</el-tag>
              </div>
            </template>
          </el-table-column>
        <el-table-column align="left" label="手数" prop="volume" width="40" />
        <el-table-column align="left" label="开仓价格" prop="price" :formatter="row => formatCurrency1(row.price)" width="80" />


        <el-table-column align="left" label="平仓价" prop="close_price" :formatter="row => formatCurrency1(row.close_price)" width="80" />
        <el-table-column align="left" label="平仓时间" prop="complete_at" width="120" />
        <el-table-column align="left" label="盈亏" prop="is_win" :formatter="row => formatDirectionWin(row.is_win)" width="80" />
        <el-table-column align="left" label="盈亏金额" prop="win_amount" :formatter="row => formatCurrency(row.win_amount)" width="80" />
        <el-table-column align="left" label="保证金" prop="bond"  :formatter="row => formatCurrency(row.bond)" width="100" />
        <el-table-column align="left" label="手续费" prop="fee" :formatter="row => formatCurrency(row.fee)" width="80" />
<!--        <el-table-column align="left" label="操作">-->
<!--            <template #default="scope">-->
<!--            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">-->
<!--                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>-->
<!--                查看详情-->
<!--            </el-button>-->
<!--&lt;!&ndash;            <el-button type="primary" link icon="edit" class="table-button" @click="updateOrdersFunc(scope.row)">变更</el-button>&ndash;&gt;-->
<!--&lt;!&ndash;            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>&ndash;&gt;-->
<!--            </template>-->
<!--        </el-table-column>-->
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
            <el-form-item label="用户编号:"  prop="user_id" >
              <el-select v-model="formData.user_id" placeholder="请选择用户编号" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
<!--            <el-form-item label="账户:"  prop="account_id" >-->
<!--              <el-input v-model.number="formData.account_id" :clearable="true" placeholder="请输入账户" />-->
<!--            </el-form-item>-->
            <el-form-item label="订单号:"  prop="order_no" >
              <el-input v-model="formData.order_no" :clearable="true"  placeholder="请输入订单号" />
            </el-form-item>
            <el-form-item label="类型:"  prop="direction" >
              <el-input v-model.number="formData.direction" :clearable="true" placeholder="请输入类型" />
            </el-form-item>
            <el-form-item label="手:"  prop="volume" >
              <el-input v-model.number="formData.volume" :clearable="true" placeholder="请输入手" />
            </el-form-item>
            <el-form-item label="价格:"  prop="price" >
              <el-input v-model.number="formData.price" :clearable="true" placeholder="请输入价格" />
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
                <el-descriptions-item label="用户编号">
                        {{ filterDict(formData.user_id,intOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="账户">
                        {{ formData.account_id }}
                </el-descriptions-item>
                <el-descriptions-item label="订单号">
                        {{ formData.order_no }}
                </el-descriptions-item>
                <el-descriptions-item label="类型">
                        {{ formData.direction }}
                </el-descriptions-item>
                <el-descriptions-item label="手">
                        {{ formData.volume }}
                </el-descriptions-item>
                <el-descriptions-item label="价格">
                        {{ formatCurrency1(formData.price) }}
                </el-descriptions-item>
          <el-descriptions-item label="保证金">
            {{ formatCurrency(formData.bond) }}
          </el-descriptions-item>
          <el-descriptions-item label="手续费">
            {{ formatCurrency(formData.fee) }}
          </el-descriptions-item>
          <el-descriptions-item label="平仓价">
            {{ formatCurrency1(formData.close_price) }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createOrders,
  deleteOrders,
  deleteOrdersByIds,
  updateOrders,
  findOrders,
  getOrdersList
} from '@/api/orders'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import {useUserStore} from "@/pinia/modules/user";

const userStore = useUserStore()

const userOptions = ref([])
const orderOptions = ref([])
defineOptions({
    name: 'Orders'
})


// 自动化生成的字典（可能为空）以及字段
const intOptions = ref([])
const formData = ref({
        user_id: undefined,
        account_id: 0,
        order_no: '',
        direction: 0,
        volume: 0,
        price: 0,
        })


// 验证规则
const rule = reactive({
               user_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               account_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               order_no : [{
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
               direction : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               volume : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               price : [{
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
const success = ref(0)
const fail = ref(0)

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
  const table = await getOrdersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list.List
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    success.value =  table.data.list.Success
    fail.value =  table.data.list.Fail
  }
}

onMounted(() => {
  // 获取路由参数
  const route = useRoute();
  const userIdFromRoute = route.query.id;

  // 判断id是否存在且不为空
  if (userIdFromRoute) {
    searchInfo.value.user_id = parseInt(userIdFromRoute, 10);
    // 在这里可以使用 searchInfo.value.user_id 进行进一步的操作
  }
  if (userStore.userInfo.authority.authorityName === "合伙人") {
    searchInfo.value.admin_id = userStore.userInfo.ID;
    success.value = 1;
  }
  getTableData();
});




// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    intOptions.value = await getDictFunc('user')
  userOptions.value = await getDictFunc('user')
  orderOptions.value = await getDictFunc('order_status')
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
            deleteOrdersFunc(row)
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
      const res = await deleteOrdersByIds({ ids })
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
const updateOrdersFunc = async(row) => {
    const res = await findOrders({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reos
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOrdersFunc = async (row) => {
    const res = await deleteOrders({ ID: row.ID })
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
  const res = await findOrders({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reos
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          user_id: undefined,
          account_id: 0,
          order_no: '',
          direction: 0,
          volume: 0,
          price: 0,
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
        user_id: undefined,
        account_id: 0,
        order_no: '',
        direction: 0,
        volume: 0,
        price: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createOrders(formData.value)
                  break
                case 'update':
                  res = await updateOrders(formData.value)
                  break
                default:
                  res = await createOrders(formData.value)
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
const formatDirection = (direction) => {
  if (direction === 2) {
    return "卖空";
  }
  return "买多" ;
}

const formatDirectionWin = (direction) => {
  if (direction === 2) {
    return "亏";
  }
  if (direction === 1) {
    return "赢";
  }
  return "--" ;
}

const formatOrderStatus= (direction) => {
  if (direction === 1) {
    return "成交";
  }
  if (direction === 2) {
    return "取消";
  }
  if (direction === 3) {
    return "失败";
  }
  if (direction === 5) {
    return "平仓";
  }
  return "失败" ;
}
const formatTagType = (direction) => {
  return direction === 1 ? 'success' : 'danger';
}

const formatCurrency = (amount) => {
  return (amount / 100).toFixed(2);
}

const formatCurrency1 = (amount) => {
  return (amount).toFixed(2);
}

const formatTime = (amount) => {
  console.log(amount);
  return amount;
}

</script>

<style>

</style>
