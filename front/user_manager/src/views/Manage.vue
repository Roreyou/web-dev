<template>
    <el-card v-loading="loading" class="box-card">
      <el-dialog
          title="提示"
          :visible.sync="dialogVisible"
          width="35%"
          :before-close="handleClose">

          <el-form ref="form" :rules="rules" :model="form" label-width="80px">
              <el-form-item label="密码" prop="password"> <!-- prop设置的是字段名字 用于校验 对于rules中的名字 -->
                  <el-input placeholder="请输入服务器密码" v-model="form.password"></el-input> <!-- v-model是双向绑定 -->
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
          loading: false,
          modalType: 0, //0开机 1删除
          dialogVisible: false,
          form: {
            password:'',
            },
            rules: {
                password: [
                    { required: false,  trigger: "blur", message: '请输入密码' }
                ],
            },
            tableData:[]
        }
    },
    mounted() {
    // 在进入页面后立即调用的操作
    this.getUseGPUData();
    },
    methods: {
      getUseGPUData() {
        // 调用接口，获取使用记录
        console.log(sessionStorage.getItem('user_id'))
        $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/user/show',
            data:{
              user_id:sessionStorage.getItem('user_id')
            },
            success: (response) => {
            // 请求成功的处理逻辑
            console.log(response);
            // 渲染数据到 el-table
            this.renderTable(response);
            },
            error: (xhr, status, error) => {
            // 请求失败的处理逻辑
            console.log('Error:' + error);
            },
        });
      },
      renderTable(item) {
        this.tableData.splice(0); // 清空tableData数组
        if(item['container_status']===2||item['container_status']===3){
          this.tableData.push({
          'ID': item['machine_id'],
          'Type': item['server_type'],
          'SSH': item['SSH'],
          //'State': false,
          'State': this.getStatusText(item['container_status'])
        });
        }       
      },
  getStatusText(statusCode) {
    if (statusCode === 2) {
      return false;//处于关机状态
    } else if (statusCode === 3) {
      return true;//处于开机状态
    } else {
      return '';
    }
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
            this.getUseGPUData();
            this.dialogVisible = false
      },
      cancel() { //点击取消时调用点击关闭同样的函数
          this.handleClose()
      },
      submit(){
        if(this.modalType === 0){
          //开机功能的提交
          this.loading=true
          console.log(sessionStorage.getItem('user_id'))
          console.log(this.form.password)
          $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/user/work',
            data: {
              user_id: sessionStorage.getItem('user_id'),
              container_password:this.form.password,
            },
            success: (response) => {
            // 请求成功的处理逻辑
            this.loading=false
            alert("已成功开机！")
            console.log("你是成功的人士");
            console.log(response);
            // 更改开关状态至绿色
            this.getUseGPUData();
            },
            error: (xhr, status, error) => {
            // 请求失败的处理逻辑
            this.loading=false
            if(error.msg==="密码不正确"){
              alert("密码不正确,开机失败！")
            }
            else if(error.msg==="时间已超额"){
              alert("时间已超额,开机失败！")
            }
            else{
              alert("该容器正在被使用，无法开机！")
            }
            console.log("你是嗯嗯的人士");
            console.log('Error:', error);
            },
        });
        }
        else if(this.modalType ===1){
          //删除功能的提交
          this.loading=true
          $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/user/close',
            data: {
              user_id: sessionStorage.getItem('user_id'),
              container_password:this.form.password,
            },
            success: (response) => {
            // 请求成功的处理逻辑
            this.loading=false
            alert("您已成功删除容器！")
            console.log("你是成功的人士");
            console.log(response);
            // 更改开关状态至绿色
            this.getUseGPUData();
            },
            error: (xhr, status, error) => {
            // 请求失败的处理逻辑
            this.loading=false
            alert("删除容器失败，请检查密码或者容器状态！")
            console.log("你是嗯嗯的人士");
            console.log('Error:', error);
            },
        });
        }

        this.handleClose();

      },
      changeSwitch(value){
        if(value){
          //开关从关到开
          this.modalType = 0 //开机功能
          this.dialogVisible = true;
        }
        else{
          //开关从开到关
          this.$confirm('是否停止使用服务器?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
          }).then(() => {

            this.loading=true
            $.ajax({
              type: 'POST',
              url: 'http://127.0.0.1:8081/user/exit',
              data: {
                user_id: sessionStorage.getItem('user_id'),
              },
              success: (response) => {
              // 请求成功的处理逻辑
              this.loading=false
              console.log("你是成功的人士");
              console.log(response);
              if(response.warning==="容器使用超出额度"){
                alert("已成功关机，鉴于您已超出容器使用额度，您本月将无法使用服务器！")
              }
              else{
                alert("已成功关机！")
              }
              // 更改开关状态至绿色
              this.getUseGPUData();
              },
              error: (xhr, status, error) => {
              // 请求失败的处理逻辑
              this.loading=false
              console.log("你是嗯嗯的人士");
              console.log('Error:', error);
              },
            });        
          })
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