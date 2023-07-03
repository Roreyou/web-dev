<template>
    <el-card class="box-card">
      <el-dialog
          title="提示"
          :visible.sync="dialogVisible"
          width="35%"
          :before-close="handleClose">

          <el-form ref="form" :rules="rules" :model="form" label-width="80px">
              <el-form-item label="密码" prop="password"> <!-- prop设置的是字段名字 用于校验 对于rules中的名字 -->
                  <el-input placeholder="请输入服务器密码" v-model="form.account"></el-input> <!-- v-model是双向绑定 -->
              </el-form-item>
          </el-form>

          <span slot="footer" class="dialog-footer">
              <el-button @click="cancel">取 消</el-button> <!-- 提交表单是触发的按钮 -->
              <el-button type="primary" @click="submit">确 定</el-button>
          </span>
      </el-dialog>

        <div slot="header" class="clearfix">
            <span>服务器列表</span>
        </div>

        <div class="text item">
            <el-table
                stripe
                :data="tableData"
                style="width: 100%">
                <el-table-column
                    prop="ID"
                    label="机器ID"
                    width="180">
                </el-table-column>
                <el-table-column
                    prop="Type"
                    label="GPU类型"
                    width="180">
                </el-table-column>
                <el-table-column
                    prop="SSH"
                    label="SSH命令"
                    width="280">
                </el-table-column>
                <el-table-column
                    label="状态"
                    width="180">
                    <template slot-scope="scope">
                        <el-switch
                            @change="changeSwitch"
                            v-model="scope.row.State"
                            active-color="#13ce66"
                            inactive-color="#ff4949">
                        </el-switch>
                    </template>
                </el-table-column>
                <el-table-column
                    label="操作"
                    width="180">
                    <template slot-scope="scope">
                        <el-button
                            size="mini"
                            type="danger"
                            @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </el-card>
</template>

<script>
//import { Scope } from 'eslint-scope';
import $ from 'jquery'

export default {
    data() {
        return {
          modalType: 0, //0开机 1删除
          dialogVisible: false,
          form: {
            password:'',
            },
            rules: {
                password: [
                    { required: true, message: '请输入密码' }
                ],
            },
            tableData:[{
            ID: '163',
            Type: 'RTX3080',
            SSH: '50501',
            State: true
          }, {
            ID: '164',
            Type: 'RTX3081',
            SSH: '50502',
            State: true
          }, {
            ID: '165',
            Type: 'RTX3082',
            SSH: '50503',
            State: true
          }, {
            ID: '166',
            Type: 'RTX3083',
            SSH: '50504',
            State: true
          }]
        }
    },
    mounted() {
    // 在进入页面后立即调用的操作
    this.getUseGPUData();
    },
    methods: {
      getUseGPUData() {
        // 调用接口，获取使用记录
        $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/show',
            data:{
              user_id:sessionStorage.getItem('user_id')
            },
            success: (response) => {
            // 请求成功的处理逻辑
            console.log(response);
            // 渲染数据到 el-table
            //this.renderTable(response);
            },
            error: (xhr, status, error) => {
            // 请求失败的处理逻辑
            console.log('Error:' + error);
            },
        });
      },
      handleDelete(index, row) {
        // console.log(index, row);
        // alert(index, row)
        this.modalType = 1 //删除功能
        this.dialogVisible = true;

      },
      handleClose() {
            this.$nextTick(() => {
                this.form = {
                password: ''
                }
                this.$refs.form.resetFields()
            })
            this.dialogVisible = false
      },
      cancel() { //点击取消时调用点击关闭同样的函数
          this.handleClose()
      },
      submit(){
        if(this.modalType === 0){
          //开机功能的提交
          alert("开机")
        }
        else if(this.modalType ===1){
          //删除功能的提交
          alert("删除")
        }
      },
      changeSwitch(value){
        if(value){
          //开关从关到开
          this.modalType = 0 //开机功能
          this.dialogVisible = true;
        }
        else{
          //开关从开到关
        }
      }
    },
    mouted(){
      getData().then((data) => {
        console.log(data)
      })
    }
}
</script>

<style>
  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }

  .box-card {
    width: 80%;
  }
</style>