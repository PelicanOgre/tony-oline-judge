
<template>
  <div>
    
   
   

  </div>
</template>
<script lang="ts" setup>
import {
  DArrowRight
} from '@element-plus/icons-vue'
import { ref, onMounted } from "vue";
//  按需引入 echarts
import * as echarts from "echarts";
const arr=ref([])
const brr=[
  {name:'Number of Exams Notified',value:12,color:'#16CEB9'},
  {name:'Number of people who have completed the exam',value:23,color:'#6648FF'},
  {name:'Number of people who have not completed the test',value:14,color:'#399'},
  {name:'The number of overdue exams',value:32,color:'#399'}
]
const charts=(el:HTMLElement)=>{
  (arr.value as Array<HTMLElement>).push(el)
}
 
 
interface chartOption{
  el:HTMLElement;
  name:string;
  num:number,
  color:string;

}
 
</script>

<style scoped lang="scss">
.card{
    padding: 10px;
    box-sizing: border-box;
}
.bar {
  display: flex;
  justify-content: space-between;
  background-color: #f5f5f5;
  line-height: 30px;
  align-items: center;
  padding: 10px 20px;
    span:nth-child(2){
        color: #999;
        font-size: 14px;
    }
}
.charts{
    display: flex;
    justify-content: space-between;
    height: 200px;
    &>div{
        width: 25%;
    }
}
h3{
    border-bottom: 1px solid #eee;
    padding: 20px;
    margin-bottom: 10px;
}
</style>
