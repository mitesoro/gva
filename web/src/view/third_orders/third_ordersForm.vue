<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户编号:" prop="user_id">
          <el-select v-model="formData.user_id" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="账户:" prop="account_id">
          <el-input v-model.number="formData.account_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单号:" prop="order_no">
          <el-input v-model="formData.order_no" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:" prop="direction">
          <el-input v-model.number="formData.direction" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手:" prop="volume">
          <el-input v-model.number="formData.volume" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model.number="formData.price" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createOrders,
  updateOrders,
  findOrders
} from '@/api/orders'

defineOptions({
    name: 'OrdersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
               }],
               account_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               order_no : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               direction : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               volume : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               price : [{
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
      const res = await findOrders({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reos
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    intOptions.value = await getDictFunc('int')
}

init()
// 保存按钮
const save = async() => {
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
