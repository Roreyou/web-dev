<template>
    <div>
        <el-container>
            <el-aside width="auto">
                <manage-aside />
            </el-aside>
            <el-container>
                <el-header>
                    <common-header />
                </el-header>
                <el-main>
                    <div>
                        <el-dialog title="提示" :visible.sync="dialogVisible" width="30%" :before-close="handleClose">

                            <el-form ref="form" :rules="rules" :model="form" label-width="80px">
                                <!-- model是属性不是v-model --> <!-- rules属性是添加校验 对应return中的rules -->
                                <el-form-item label="用户ID" prop="UserID">
                                    <el-input placeholder="请输入用户ID" v-model="form.UserID"></el-input>
                                </el-form-item>

                                <el-form-item label="服务器ID" prop="SeverID"> <!-- prop设置的是字段名字 用于校验 对于rules中的名字 -->
                                    <el-input placeholder="请输入服务器ID" v-model="form.SeverID"></el-input> <!-- v-model是双向绑定 -->
                                </el-form-item>

                                <el-form-item label="镜像ID" prop="ImageID">
                                    <el-input placeholder="请输入镜像ID" v-model="form.ImageID"></el-input>
                                </el-form-item>

                                <el-form-item label="容器密码" prop="Password">
                                    <el-input placeholder="请输入容器密码" v-model="form.Password"></el-input>
                                </el-form-item>


                                <!-- <el-form-item label="使用时间" prop="Usetime">
                                    <el-tag :key="tag" v-model="form.Usetime" v-for="tag in form.Usetime" closable
                                        :disable-transitions="false" @close="tagClose(tag)">
                                        {{ tag }}
                                    </el-tag>
                                    <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue"
                                        ref="saveTagInput" size="small" @keyup.enter.native="tagInputConfirm"
                                        @blur="tagInputConfirm">
                                    </el-input>
                                    <el-button v-else class="button-new-tag" size="small" @click="tagInput">+ New
                                        Tag</el-button>

                                </el-form-item> -->
                            </el-form>

                            <span slot="footer" class="dialog-footer">
                                <el-button @click="cancel">取 消</el-button> <!-- 提交表单是触发的按钮 -->
                                <el-button type="primary" @click="submit">确 定</el-button>
                            </span>

                        </el-dialog>

                        <div class="manage-header">

                            <el-button @click="handleAdd()" type="primary"> <!-- 点击按钮时触发上面dialog的visible.sync属性 -->
                                + 新增
                            </el-button>
                            <!-- form搜索区域 -->
                            <!-- <el-form :inline="true" :model="userForm">
                                <el-form-item>
                                    <el-input placeholder="请输入名称" v-model="userForm.name"></el-input>
                                </el-form-item>
                                <el-form-item>
                                    <el-button type="primary" @click="onSubmit">查询</el-button>
                                </el-form-item>
                            </el-form> -->
                        </div>

                        <el-table stripe height="630px" :data="TableDataa" style="width: 100%">
                            
                            <el-table-column prop="UserID" label="用户ID" width="200">
                            </el-table-column>
                            
                            <el-table-column prop="Password" label="容器密码" width="200">
                            </el-table-column>
                            
                            <el-table-column prop="ID" label="镜像ID" width="150">
                            </el-table-column>

                            <el-table-column prop="IP" label="容器IP" width="200"> <!-- 这个name是数据中的字段名 和上面无关 -->
                            </el-table-column>

                            <el-table-column prop="Port" label="容器端口" width="150">
                            </el-table-column>

                            <el-table-column prop="Status" label="容器状态" width="200">
                            </el-table-column>

                            <!-- <el-table-column label="容器使用时间" width="550">
                                <template slot-scope="scope">
                                    <div v-for="item in scope.row.Usetime" :key="item">{{ item }}</div>
                                </template>
                            </el-table-column> -->

                            <el-table-column label="操作">
                                <template slot-scope="scope">
                                    <el-button type="danger" size="mini" @click="handleDelete(scope.row)">删除</el-button>
                                </template>
                            </el-table-column>
                        </el-table>

                    </div>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>


<script>
import $ from 'jquery'
import ManageAside from '@/components/ManageAside.vue';
import CommonHeader from '@/components/CommonHeader.vue';
export default {
    data() {
        return {
            dialogVisible: false,
            modalType: 0,
            inputVisible: false,
            inputValue: '',
            getStatus: -1,
            form: {
                UserID: '',
                SeverID: '',
                ImageID: '',
                Password: '',
                //Usetime: []
            },
            rules: {
                UserID: [
                    { required: true, message: '请输入用户ID' } //message是错误时出现的信息 每一个{}代表一条规则
                ],
                SeverID: [
                    { required: true, message: '请输入服务器ID' }
                ],
                ImageID: [
                    { required: true, message: '请输入镜像ID' }
                ],
                Password: [
                    { required: true, message: '请输入容器密码' }
                ],
                // Usetime: [
                //     { required: true, message: '请输入容器使用时间' }
                // ]
            },
            TableDataa: []
        }
    },
    mounted() {
    // 在进入页面后立即调用的操作
    this.getContdata();
  },
    components: {
        ManageAside,
        CommonHeader
    },
    methods: {
        getContdata() {
    // 调用接口，获取使用记录
    $.ajax({
      type: 'POST',
      url: 'http://127.0.0.1:8081/admin/get_container',
      data: {},
      success: (response) => {
        // 请求成功的处理逻辑
        console.log(response);
        // 渲染数据到 el-table
        this.renderTable(response);
      },
      error: (xhr, status, error) => {
        // 请求失败的处理逻辑
        console.log('Error:', error);
      },
    });
  },
  renderTable(data) {
  if (Array.isArray(data)) {
    this.TableDataa = data.reduce((acc, item) => {
      if (item['容器的状态'] === 2 || item['容器的状态'] === 3) {
        acc.push({
          'UserID': item['创建容器的用户'],
          'Password': item['容器密码'],
          'ID': item['镜像ID'],
          'IP': item['容器所在ip'],
          'Port': item['容器端口'],
          'Status': this.getStatusText(item['容器的状态']),
        });
      }
      return acc;
    }, []);
  } else {
    this.TableDataa = []; // Assign an empty array if data is not an array
  }
},
  getStatusText(statusCode) {
    if (statusCode === 2) {
      return '已创建且未使用';
    } else if (statusCode === 3) {
      return '正在使用';
    } else {
      return '';
    }
  },
  

    // judge(int num){
    //     if(num==1){
    //         return 
    //     }
    // },
        // 弹窗关闭的时候
        handleClose() {
            this.$nextTick(() => {
                this.form = {
                    ID: '',
                    IP: '',
                    Port: '',
                    Status: '',
                    //Usetime: []
                }
                this.$refs.form.resetFields()
            })
            this.dialogVisible = false
        },
        cancel() { //点击取消时调用点击关闭同样的函数
            this.handleClose()
        },
        handleDelete(row){
            console.log(row)
            $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/admin/delete_container',
            data: {
                user_id:row.UserID,
                container_password:row.Password

            },
            success: (response) => {
            // 请求成功的处理逻辑
            console.log(response);
            alert("删除成功！")
            location.reload()
            },
            
        });
        this.getContdata();
       },
        // 提交用户表单
        submit() {
            this.$refs.form.validate((valid) => { //this.$refs.form获取表单实例 validate用来校验表单 valid是返回值 bool类型
                // console.log(valid, 'valid')
                if (valid) { //valid为真 校验通过
                    // 后续对表单数据的处理
                    if (this.modalType === 0) { //是否是新增 而不是编辑
                        $.ajax({
                            type: 'POST',
                            url: 'http://127.0.0.1:8081/admin/add_container',
                            data: {
                                user_id:this.form.UserID,
                                server_id: this.form.SeverID,
                                image_id: this.form.ImageID,
                                user_password: this.form.Password,
                            },
                            success: (response) => {
                            // 请求成功的处理逻辑
                            console.log(response);
                            // 渲染数据到 el-table
                            this.renderTable(response);
                            }
                        });
                        
                    } else {
                        console.log('修改')
                    }
                    this.getContdata();
                    this.handleClose();
                }
            })
        },
        handleAdd() {
            this.modalType = 0
            this.dialogVisible = true;
        },
    }
}
</script>

<style lang="less" scoped>
.el-header {
    padding: 0px;
}
</style>