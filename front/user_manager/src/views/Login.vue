<template>
  <el-form
    ref="form"
    label-width="70px"
    :inline="true"
    class="login-container"
    :model="form"
    :rules="rules"
  >
    <h3 class="login_title">GPU管理系统登录</h3>
    <el-form-item label="学号" prop="user_id">
      <el-input v-model="form.user_id" placeholder="请输入学号"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="user_password">
      <el-input
        type="password"
        v-model="form.user_password"
        placeholder="请输入密码"
      ></el-input>
    </el-form-item>
    <el-form-item>
      <el-button
        @click="submit"
        style="margin-left: 105px; margin-top: 10px"
        type="primary"
        >登录</el-button
      >
    </el-form-item>
  </el-form>
</template>

<script>

import $ from 'jquery'
import VueRouter from 'vue-router'
import Manage from '../views/Manage.vue'

export default {
  data() {
    
    return {
      user_id:"",
      user_name:"",
      real_name:"",
      token:"",
      user_password:"",
      form: {
        user_id: "",
        user_password: "",
      },
      rules: {
        user_id: [
          { required: true, trigger: "blur", message: "该栏不能为空" },
        ],
        user_password: [
          { required: true, trigger: "blur", message: "该栏不能为空" },
        ],
      },
    };
  },
  mounted(){
    //this.ifPush()
    
  },

  methods: {
  //   ifPush(){
  //     let if_token =sessionStorage.getItem('token')
  //     console.log(if_token)
  //     if(if_token !== null)
  //     {
  //         this.$router.push('/manage')//用户已登录时，使其不再访问登录页面
  //     }
  //     },


    submit() {
       // 发送 POST 请求验证用户身份
       if(this.form.user_id === "admin" && this.form.user_password === "888"){
        const token = "1";
        sessionStorage.setItem("token", token);
        this.$router.push('/usermanage');
       }
       else{
        $.ajax({
        type: "POST",
        url: "http://127.0.0.1:8081/user",
        data: {
            user_id: this.form.user_id,
            user_password: this.form.user_password,
            },
            success: (response) => {
        console.log(response);
        // 获取用户信息
        const token = response.token;
        const user_id = response.UserInfo.user_id;
        const user_name = response.UserInfo.user_name;
        const real_name = response.UserInfo.real_name;
        const remainder = response.UserInfo.remainder;
        const user_password = response.UserInfo.user_password;
        // 存储用户信息
        sessionStorage.setItem("token", token);
        sessionStorage.setItem("user_id", user_id);
        sessionStorage.setItem("user_name", user_name);
        sessionStorage.setItem("real_name", real_name);
        sessionStorage.setItem("remainder", remainder);
        sessionStorage.setItem("user_password", user_password);
        // 将用户重定向到首页或指定页面
        this.$router.push("/manage");
        },
        error: function(xhr, status, error) {
        // 请求失败的处理逻辑
        console.log("Error:", error);
        alert("登录失败，请检查用户名和密码是否正确");
        }
        });
      }
       
   
      },
  },
  
};
</script>
<style lang="less" scoped>
.login-container {
  width: 350px;
  border: 1px solid #eaeaea;
  margin: 180px auto;
  padding: 35px 35px 15px 35px;
  background-color: #fff;
  border-radius: 15px;
  box-shadow: 0 0 25px #cac6c6;
  box-sizing: border-box;
  .login_title {
    text-align: center;
    margin-bottom: 40px;
    color: #505458;
  }
  .el-input {
    width: 198px;
  }
}
</style>