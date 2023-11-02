<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户编号:" prop="user_id">
          <el-select v-model="formData.user_id" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in userOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="金额类型:" prop="amount_type">
          <el-select v-model="formData.amount_type" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in amount_logOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="金额:" prop="amount">
          <el-input v-model.number="formData.amount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前金额:" prop="cur_amount">
          <el-input v-model.number="formData.cur_amount" :clearable="true" placeholder="请输入" />
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
  createAlog,
  updateAlog,
  findAlog
} from '@/api/alog'

defineOptions({
    name: 'AlogForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const userOptions = ref([])
const amount_logOptions = ref([])
const formData = ref({
            user_id: undefined,
            amount_type: undefined,
            amount: 0,
            cur_amount: 0,
        })
// 验证规则
const rule = reactive({
               user_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               amount_type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               amount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cur_amount : [{
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
      const res = await findAlog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.real
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    userOptions.value = await getDictFunc('user')
    amount_logOptions.value = await getDictFunc('amount_log')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAlog(formData.value)
               break
             case 'update':
               res = await updateAlog(formData.value)
               break
             default:
               res = await createAlog(formData.value)
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
