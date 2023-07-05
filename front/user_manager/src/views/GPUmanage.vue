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
                        <el-dialog
                            title="提示"
                            :visible.sync="dialogVisible"
                            width="50%"
                            :before-close="handleClose">

                            <el-form ref="form" :rules="rules" :model="form" label-width="80px"> 
                                <!-- model是属性不是v-model --> <!-- rules属性是添加校验 对应return中的rules -->
                                <el-form-item label="机器ID" prop="ID">
                                    <el-input placeholder="请输入机器ID" v-model="form.ID"></el-input>
                                </el-form-item>
                                
                                <el-form-item label="型号" prop="Type"> <!-- prop设置的是字段名字 用于校验 对于rules中的名字 -->
                                    <el-input placeholder="请输入GPU型号" v-model="form.Type"></el-input> <!-- v-model是双向绑定 -->
                                </el-form-item>

                                <el-form-item label="显存" prop="Size">
                                    <el-input placeholder="请输入GPU显存大小" v-model="form.Size"></el-input>
                                </el-form-item>

                                <el-form-item label="框架" prop="Frame">
                                    <el-tag
                                    :key="tag"
                                    v-model="form.Frame"
                                    v-for="tag in form.Frame"
                                    closable
                                    :disable-transitions="false"
                                    @close="tagClose(tag)">
                                    {{tag}}
                                    </el-tag>
                                    <el-input
                                    class="input-new-tag"
                                    v-if="inputVisible"
                                    v-model="inputValue"
                                    ref="saveTagInput"
                                    size="small"
                                    @keyup.enter.native="tagInputConfirm"
                                    @blur="tagInputConfirm"
                                    >
                                    </el-input>
                                    

                                </el-form-item>
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

                        <el-table
                            stripe
                            height="630px"
                            :data="TableDataa"
                            style="width: 100%">
                            <el-table-column
                                prop="ID" 
                                label="机器ID"
                                width="150">
                            </el-table-column>
                            <el-table-column
                                prop="Type" 
                                label="GPU型号"
                                width="200"> <!-- 这个name是数据中的字段名 和上面无关 -->
                            </el-table-column>
                            <el-table-column
                                prop="Size" 
                                label="显存大小"
                                width="150">
                            </el-table-column>
                            <el-table-column
                                label="预装软件框架"
                                width="550">
                                <template slot-scope="scope">
                                    <div v-for="item in scope.row.Frame">{{ item }}</div>
                                </template>
                            </el-table-column>
                            <el-table-column
                                label="操作">
                                <template slot-scope="scope">
                                    <el-button size="mini" @click="handleEdit(scope.row)">编辑</el-button> <!-- (scope.row)是传参 当前的行数 -->
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
            form: {
                ID: '',
                Type: '',
                Size: '',
                Frame: ['tensorflow2.1.0+pythorch1.11+python3.6.9', 'tensorflow2.6.0+pythorch1.11+python3.9.0']
            },
            rules: {
                ID: [
                    { required: true, message: '请输入机器ID' } //message是错误时出现的信息 每一个{}代表一条规则
                ],
                Type: [
                    { required: true, message: '请输入GPU型号' } 
                ],
                Size: [
                    { required: true, message: '请输入GPU显存大小' }
                ],
                Frame: [
                    { required: true, message: '请输入预装软件框架' }
                ]
            },
            TableDataa:[]
        }
    },
    components: {
        ManageAside,
        CommonHeader
    },
    mounted() {
    // 在进入页面后立即调用的操作
    this.getGPUData();
    },
    methods: {
        getGPUData() {
        // 调用接口，获取使用记录
        $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/admin/get_server',
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
        // 检查数据是否为数组
        if (Array.isArray(data)) {
            this.TableDataa = data.map((item) => ({
            'ID': item['server_id'],
            'Type': item['sever_type'],
            'Size': item['server_size'],
            'Frame': ['tensorflow2.1.0+pythorch1.11+python3.6.9', 'tensorflow2.6.0+pythorch1.11+python3.9.0']
            }));
        } else {
            this.TableDataa = []; // 如果数据不是数组，将 TableDataa 设置为空数组
        }
    },

    handleDelete(row){
            console.log(row)

            this.$confirm('此操作将永久删除该服务器, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
            $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/admin/delete_server',
            data: {
                server_id:row.ID,
            },
            success: (response) => {
            // 请求成功的处理逻辑
            console.log(response);
            alert("删除GPU成功！")
            location.reload()
            },       
        });
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });          
        });
        this.getUserData();
       },
        // 提交用户表单
        submit() {
            this.$refs.form.validate((valid) => { //this.$refs.form获取表单实例 validate用来校验表单 valid是返回值 bool类型
                // console.log(valid, 'valid')
                if (valid) { //valid为真 校验通过
                    // 后续对表单数据的处理
                    if (this.modalType === 0) { //是否是新增 而不是编辑
                        console.log('新增')
                        //新增的逻辑
                        $.ajax({
                            type: 'POST',
                            url: 'http://127.0.0.1:8081/admin/add_server',
                            data: {
                                server_id: this.form.ID,
                                server_type: this.form.Type,
                                server_size: this.form.Size,
                            },
                            success: (response) => {
                            // 请求成功的处理逻辑
                            console.log(response);
                            // 渲染数据到 el-table
                            alert("新增GPU成功！")
                            this.renderTable(response);
                            }
                        });
                        this.getGPUData();
                    } else {
                        console.log('修改')
                        //编辑按钮的接口实现
                        $.ajax({
                            type: 'POST',
                            url: 'http://127.0.0.1:8081/admin/update_server',
                            data: {
                                server_id: this.form.ID,
                                server_type: this.form.Type,
                                server_size: this.form.Size,
                            },
                            success: (response) => {
                            // 请求成功的处理逻辑
                            console.log(response);
                            // 渲染数据到 el-table
                            alert("修改GPU成功！")
                            this.renderTable(response);
                            }
                        });
                        this.getGPUData();
                    }

                    // 清空表单的数据
                    this.$refs.form.resetFields()
                    // 关闭弹窗
                    this.dialogVisible = false
                }
            })
        },
        // 弹窗关闭的时候
        handleClose() {
            this.$nextTick(() => {
                this.form = {
                ID: '',
                Type: '',
                Size: '',
                SSH: '',
                Frame: ['tensorflow2.1.0+pythorch1.11+python3.6.9', 'tensorflow2.6.0+pythorch1.11+python3.9.0']
                }
                this.$refs.form.resetFields()
            })
            this.dialogVisible = false
        },
        cancel() { //点击取消时调用点击关闭同样的函数
            this.handleClose()
        },
        handleAdd() {
            this.modalType = 0
            this.dialogVisible = true;
        },
        handleEdit(row) {
            this.modalType = 1
            this.dialogVisible = true
            // 注意需要对当前行数据进行深拷贝，否则会出错
            console.log(row)
            this.form = JSON.parse(JSON.stringify(row))
        },
        tagClose(tag) {
            this.form.Frame.splice(this.form.Frame.indexOf(tag), 1);
        },

        tagInput() {
            this.inputVisible = true;
            this.$nextTick(_ => {
            this.$refs.saveTagInput.$refs.input.focus();
            });
        },

        tagInputConfirm() {
            let inputValue = this.inputValue;
            if (inputValue) {
            this.form.Frame.push(inputValue);
            }
            this.inputVisible = false;
            this.inputValue = '';
        }
    }
}
</script>

<style lang="less" scoped>
.el-header {
    padding: 0px;
}
  .el-tag + .el-tag {
    margin-left: 10px;
  }
  .button-new-tag {
    margin-left: 10px;
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }
  .input-new-tag {
    width: 90px;
    margin-left: 10px;
    vertical-align: bottom;
  }
</style>