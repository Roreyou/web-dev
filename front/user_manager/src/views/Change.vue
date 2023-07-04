<template>
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        
      <el-form-item label="用户名">
            <el-input class="board" v-model="ruleForm.user_name" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>
        
      <el-form-item label="真实姓名">
            <el-input class="board" v-model="ruleForm.real_name" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>
      
      <el-form-item label="账号">
            <el-input class="board" v-model="ruleForm.account" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>

        <el-form-item label="旧密码" prop="oldPass">
            <el-input class="board" type="password" v-model="ruleForm.oldPass" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="密码" prop="pass">
            <el-input class="board" type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="确认密码" prop="checkPass">
            <el-input class="board" type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
    </el-form>
</template>

<style>
.board {
    width: 400px;
}
</style>

<script>
import $ from 'jquery'
  export default {
    data() {
      var validoldPass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入旧密码'));
        } else {
            callback();
        }
      }
      var validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {
          if (this.ruleForm.checkPass !== '') {
            this.$refs.ruleForm.validateField('checkPass');
          }
          callback();
        }
      };
      var validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.ruleForm.pass) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
      return {
        ruleForm: {
          //把从login保存的东西拿来
          user_name:sessionStorage.getItem("user_name"),
          real_name:sessionStorage.getItem("real_name"),
          account:sessionStorage.getItem("user_id"),
          oldPass: '',
          pass: '',
          checkPass: ''
        },
        rules: {//验证规则
          oldPass: [
            { validator: validoldPass, trigger: 'blur' }
          ],
          pass: [
            { validator: validatePass, trigger: 'blur' }
          ],
          checkPass: [
            { validator: validatePass2, trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
   submitForm(formName) {
            $.ajax({
          type: "POST",
          url: "http://127.0.0.1:8081/user/change_password",
          data: {
            // 请求的数据
            user_id: sessionStorage.getItem("user_id"),
            user_password: this.ruleForm.oldPass,
            new_password:this.ruleForm.checkPass
          },
          success: function(response) {
            // 请求成功的处理逻辑
            console.log(response);
          },
          error: function(xhr, status, error) {
            // 请求失败的处理逻辑
            console.log("Error:", error);
          }
        });

      // //ref 加在普通的元素上，用this.$refs.（ref值） 获取到的是dom元素
      // //ref 加在子组件上，用this.$refs.（ref值） 获取到的是组件实例，
      //   //可以使用组件的所有方法。在使用方法的时候直接this.$refs.（ref值）.方法（） 就可以使用了。
      // //API (web 或 XML 页面) = DOM + JS (脚本语言),文档对象模型 (DOM) 是 HTML 和 XML 文档的编程接口。
      // this.$refs[formName].validate(async (valid) => {
      //   if (!valid) return;

      //   const payload = {
      //     user_id: this.ruleForm.account,
      //     user_password: this.ruleForm.oldPass,
      //     new_password: this.ruleForm.pass,
      //   };

      //   try {
      //     const response = await this.axios.post(
      //       "http://127.0.0.1:8080/user/change_password",
      //       payload
      //     );

      //     const data = response.data;
      //     if (data.Success === "修改成功！") {
      //       this.$message.success("重置密码成功，正在跳转登录页面");
      //       this.$router.push("/Login").catch((err) => {
      //         console.error(err);
      //       });
      //     } else if (data.fail === "原密码错误") {
      //       this.$message.error("原密码错误");
      //     } else {
      //       this.$message.error("重置密码失败！");
      //     }
      //   } catch (error) {
      //     console.error(error);
      //     this.$message.error("重置密码失败！");
      //   }
      // });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  },

  }
</script>