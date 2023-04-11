<script setup lang="ts">
import { computed, ref } from 'vue'
import {useStore} from 'vuex'
import LoginPage from '../page/login.vue'
import {ElMessage} from 'element-plus'
import {useRouter} from 'vue-router'
import {
  Fold,Expand,Setting
} from '@element-plus/icons-vue'
 const  circleUrl='https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
const store=useStore()
const router=useRouter()
const collapse=computed(()=>store.state.collapse)
const isLogin=computed(()=>store.state.isLogin)
var is_admin=computed(()=>store.state.is_admin)
const is_admin_login=localStorage.getItem("is_admin")
const username=computed(()=>store.state.username)
function changeMenu(){
  store.commit('changeCollapse',!collapse.value)
}
const handleCommand = (command: string | number | object) => {
  if(command=='a'){
    localStorage.clear()
    // location.reload()
    

    store.commit('logout')

    var s=window.location.href
    window.location.replace(s.split("/")[0])

  }else if(command=='b'){//分类
    console.log(is_admin_login)

    if (is_admin){
      router.push('/questionManage')
    }
    else{
      ElMessage.error("Administrator Rights Required")
    }
    
  }else if(command=='c'){//忘记密码
    router.push('/forgetPassword')

  }
}
 const showLogin=ref(false)
 const loginSucc=()=>{
   showLogin.value=false
 }
</script>

<template>
 
 
  <div>
      <el-dropdown @command="handleCommand" v-if="isLogin">

        <div class="el-dropdown-link">
          
          <span>Hello! {{username}}</span>
           
           <el-icon><setting /></el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            
            <el-dropdown-item :icon="Plus" command="a">Logout</el-dropdown-item>
  
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <span class="login" @click="showLogin=true" v-else>Login</span>
    <el-dialog
    v-model="showLogin"
    title="Login/Register"
    width="500px"
    :before-close="handleClose"
  >
    <LoginPage @loginSucc="loginSucc" /> 
    <!-- <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false"
          >Confirm</el-button
        >
      </span>
    </template> -->
  </el-dialog>
  </div>

  
   <!-- 管理员
   <div class="setting">
     <el-icon><setting /></el-icon>
   </div> -->
 
</template>

<style scoped lang="scss">
.fold{
  padding-left: 20px;
}
 /* .user-msg{
   display: flex;
   justify-content: space-between;
   align-items: center;
   width: 20%;
   max-width: 250px;
   padding: 0 10px;
 } */
 .el-dropdown-link{
  color: white;
   display: flex;
   font-size: 20px;
   justify-content: space-between;
   align-items: center;
    padding-right: 30px;
    span{
      padding: 0 10px;
    }
 }
 .login{
   color: white;
   padding: 0 30px;
   cursor: pointer;
   font-size: 20px;
 }
</style>
