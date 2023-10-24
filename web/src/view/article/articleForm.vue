<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="内容:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item label="作者:" prop="author">
          <el-input v-model="formData.author" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章分类:" prop="article_category">
          <el-select v-model="formData.article_category" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createArticle,
  updateArticle,
  findArticle
} from '@/api/article'

defineOptions({
    name: 'ArticleForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const genderOptions = ref([])
const formData = ref({
            title: '',
            content: '',
            author: '',
            article_category: undefined,
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               article_category : [{
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
      const res = await findArticle({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rea
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('article_category')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createArticle(formData.value)
               break
             case 'update':
               res = await updateArticle(formData.value)
               break
             default:
               res = await createArticle(formData.value)
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
