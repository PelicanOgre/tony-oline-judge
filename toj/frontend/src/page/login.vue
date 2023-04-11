<template>
  <div class="login-page">
    <el-tabs
        v-model="activeName"
        type="card"
        class="demo-tabs"
        @tab-click="handleClick"
    >
      <el-tab-pane label="Login" name="first">
        <el-form
            ref="ruleFormRef"
            :model="ruleForm"
            :rules="rules"
            label-width="100px"
            class="login-form"
            :size="formSize"
        >
          <el-form-item label="Email" prop="mail">
            <el-input placeholder="please enter email" v-model="ruleForm.mail"/>
          </el-form-item>

          <el-form-item label="Password" prop="password">
            <el-input v-model="ruleForm.password" show-password/>
          </el-form-item>
          <div style="text-align: center">
            <el-button type="primary" @click="submitForm(ruleFormRef)"
            >Login
            </el-button
            >
          </div>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="Register" name="second">

        <el-form
            ref="registFormRef"
            :model="registForm"
            :rules="rules"
            label-width="100px"
            class="login-form"
            :size="formSize"
        >
          <el-form-item label="Username" prop="name">
            <el-input v-model="registForm.name"/>
          </el-form-item>

          <el-form-item label="Password" prop="password">
            <el-input v-model="registForm.password" show-password/>
          </el-form-item>
          <el-form-item label="Email" prop="mail">
            <el-input v-model="registForm.mail"/>
          </el-form-item>
          <el-form-item label="Code" prop="code">
            <el-row :gutter="20">
              <el-col :span="12">
                <el-input v-model="registForm.code"/>
              </el-col>
              <el-col :span="12" style="text-align:right;">
                <el-button disabled v-if="remainTime>0&&remainTime<60">{{ remainTime }}s</el-button>
                <el-button @click="sendCode" v-else type="primary">Send Code</el-button>
              </el-col>
            </el-row>


          </el-form-item>

          <div style="text-align: center">
            <el-button type="primary" @click="subRegister(registFormRef)"
            >Register
            </el-button
            >
          </div>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="Forget Password" name="third">

      <el-form
          ref="resetPasswordFormRef"
          :model="resetPasswordForm"
          :rules="rules"
          label-width="100px"
          class="login-form"
          :size="formSize"
      >
      
    
        <el-form-item label="Email" prop="mail">
          <el-input v-model="resetPasswordForm.mail"/>
        </el-form-item>
        <el-form-item label="Code" prop="code">
          <el-row :gutter="20">
            <el-col :span="12">
              <el-input v-model="resetPasswordForm.code"/>
            </el-col>
            <el-col :span="12" style="text-align:right;">
              <el-button disabled v-if="remainTime>0&&remainTime<60">{{ remainTime }}s</el-button>
              <el-button @click="sendCode1" v-else type="primary">Send Code</el-button>
            </el-col>
          </el-row>


        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input v-model="resetPasswordForm.password" show-password/>
        </el-form-item>

        <div style="text-align: center">
          <el-button type="primary" @click="subReset(resetPasswordFormRef)"
          >Reset Password
          </el-button
          >
        </div>
      </el-form>
      </el-tab-pane>

    </el-tabs>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import {useStore} from "vuex";
import {useRouter} from "vue-router";
import type {FormInstance, TabsPaneContext} from "element-plus";
import {ElMessage} from "element-plus"
import api from '../api/api.js'

const activeName = ref("first");

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};
const emits = defineEmits(["loginSucc"]);
const formSize = ref("default");
const store = useStore();
const ruleFormRef = ref<FormInstance>();
const registFormRef = ref<FormInstance>()
const resetPasswordFormRef = ref<FormInstance>()
const remainTime = ref(60)
const ruleForm = reactive({
  mail: "",
  password: "",
});
const registForm = reactive({
  name: "",
  password: "",
  mail: '',
  code: ''
});

const resetPasswordForm = reactive({
  password: "",
  mail: '',
  code: ''
  
});

const rules = reactive({
  username: [
    {required: true, message: "Please Enter Email", trigger: "blur"},

  ],
  name: [
    {required: true, message: "Please Enter Username", trigger: "blur"},

  ],
  code: [
    {required: true, message: "Please Enter Verification Code", trigger: "blur"},

  ],
  mail: [
    {required: true, message: "Please Enter Email", trigger: "blur"},

  ],


  password: [
    {required: true, message: "Please Enter Password", trigger: "blur"},

  ],
});
const router = useRouter();
console.log(router);
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      api.login(ruleForm).then((res: any) => {
        console.log(res);
        if (res.data.code == 200) {
          ElMessage.success('Login successfully!')

          localStorage.setItem("token", res.data.data.token);
          store.commit("loginSucc", res.data.data.token);
          store.commit("setUser", {username: res.data.data.name, is_admin: res.data.data.is_admin});
          localStorage.setItem('is_admin', res.data.data.is_admin)

          localStorage.setItem('username', res.data.data.name)
          localStorage.setItem('is_login', "1")
          emits("loginSucc");
          var s=window.location.href
          window.location.replace(s.split("/")[0])
        } else {
          ElMessage.error(res.data.msg)
        }

      })

      // router.push('/index')
    } else {
      console.log("error submit!", fields);
    }
  });
};
const subRegister = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      api.register(registForm).then((res:any) => {
        if (res.data.code == 200) {
          ElMessage.success('Register Successfully!')
          localStorage.setItem("token", res.data.data.token);
          store.commit("loginSucc", res.data.data.token);
          store.commit("setUser", {username: registForm.name, is_admin: res.data.data.is_admin});

          localStorage.setItem('username', registForm.name)
          localStorage.setItem('is_admin', res.data.data.is_admin)
          localStorage.setItem('is_login', "1")

          emits("loginSucc");
          var s=window.location.href
          window.location.replace(s.split("/")[0])
        } else {
          ElMessage.error(res.data.msg)

        }
      })
      // router.push('/index')
    } else {
      console.log("error submit!", fields);
    }
  });
};

const subReset = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      api.forgetPassword(resetPasswordForm).then((res:any) => {
        console.log(res);
        if (res.data.code == 200) {
          console.log(res)
          ElMessage.success('Password Reset Successfully!')
          localStorage.setItem("token", res.data.data.token);
          store.commit("loginSucc", res.data.data.token);
          store.commit("setUser", {username: res.data.data.name, is_admin: res.data.data.is_admin});

          localStorage.setItem('username', res.data.data.name)
          localStorage.setItem('is_admin', res.data.data.is_admin)
          localStorage.setItem('is_login', "1")

          emits("loginSucc");
          var s=window.location.href
          window.location.replace(s.split("/")[0])
        } else {
          ElMessage.error(res.data.msg)

        }
      })
      // router.push('/index')
    } else {
      console.log("error submit!", fields);
    }
  });
};


const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
const startRemain = () => {
  if (remainTime.value > 0) {
    remainTime.value--
    setTimeout(startRemain, 1000)
  } else {
    remainTime.value = 60
  }
}
const sendCode = () => {
  const re = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
  if (re.test(registForm.mail)) {
    startRemain()
    api.sendCode({
      email: registForm.mail
    }).then((res: any) => {
      if (res.data.code == 200) {
        ElMessage.success(res.data.msg)
      } else {
        ElMessage.error(res.data.msg)
      }
    })
  } else {
    ElMessage('please Enter Correct Email Address!')
  }
}
const sendCode1 = () => {
  const re = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
  if (re.test(resetPasswordForm.mail)) {
    startRemain()
    console.log(resetPasswordForm.mail)
    api.sendCode({
      email: resetPasswordForm.mail
    }).then((res: any) => {
      if (res.data.code == 200) {
        ElMessage.success(res.data.msg)
      } else {
        ElMessage.error(res.data.msg)
      }
    })
  } else {
    ElMessage('please Enter Correct Email Address!')
  }
}
</script>
<style lang="scss" scoped>
.login-page {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-form {
  /*widows: 300px;*/
  border: 3px solid #eee;
  padding: 50px 90px 40px 90px;
  border-radius: 15px;
}
</style>