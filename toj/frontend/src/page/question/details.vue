<template>
    
        <div class="ques-cont">
            <div class="left">
                    <h3>{{detail.title}}</h3>
                     <div class="msg">
                       Category:  <span v-for="mi in detail.problem_categories" :key="mi.id">
                                {{mi.category_basic.name}}
                            </span>
                    </div>
                    <div class="msg">
                        
                        <span>Submit Number: {{detail.submit_num}}</span>
                        <span>Accept Number: {{detail.pass_num}}</span>
                        <span>Maximum Runtime: {{detail.max_runtime}}ms</span>
                        <span>Maximum Runtime: {{detail.max_mem}}Mb</span>
                        
                    </div>
                    <div v-html="detail.content"></div>
                   
            </div>
            <div class="right">
                <Editor></Editor>  
            </div>
        </div>
   
</template>
<script lang="ts" setup>
 import { reactive,ref } from '@vue/reactivity'
 import {useRoute} from 'vue-router'
 import api from '../../api/api.js'
 import Editor from './editor.vue'
 import {
  Edit
} from '@element-plus/icons-vue'
const route=useRoute()
const detail=ref({

})
const getDetail=()=>{
    api.getProblemDetail({
        identity:route.query.identity
    }).then(res=>{
            if(res.data.data){
                detail.value=res.data.data
            }
    })
}
getDetail()
</script>
<style scoped lang="scss">
.ques-cont{
    font-size: 20px;
    display: flex;
    height: 100%;
    .left{
        width: 50%;
        line-height: 2;
        border-right: 10px solid #eee;
        padding: 10px;
        .msg{
            color:black;
            font-size: 15px;
            
            span{
                margin-right: 20px;
                color: #999;
            }
        }
    }
    .right{
        width: 50%;
    }
}
</style>