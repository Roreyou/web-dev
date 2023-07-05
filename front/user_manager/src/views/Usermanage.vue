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
                    <div class="usermanage">
                        <el-dialog
                            title="提示"
                            :visible.sync="dialogVisible"
                            width="35%"
                            :before-close="handleClose">

                            <el-form ref="form" :rules="rules" :model="form" label-width="80px">
                                <el-form-item label="账号" prop="account"> <!-- prop设置的是字段名字 用于校验 对于rules中的名字 -->
                                    <el-input placeholder="请输入账号" v-model="form.account"></el-input> <!-- v-model是双向绑定 -->
                                </el-form-item>
                                <el-form-item label="用户名" prop="name">
                                    <el-input placeholder="请输入用户名" v-model="form.name"></el-input>
                                </el-form-item>
                                <el-form-item label="真实姓名" prop="real_name">
                                    <el-input placeholder="请输入真实姓名" v-model="form.real_name"></el-input>
                                </el-form-item>
                            </el-form>

                            <span slot="footer" class="dialog-footer">
                                <el-button @click="cancel">取 消</el-button> <!-- 提交表单是触发的按钮 -->
                                <el-button type="primary" @click="submit">确 定</el-button>
                            </span>

                        </el-dialog>

                        <div class="usermanage-header">
                            <el-button @click="handleAdd" type="primary"> <!-- 点击按钮时触发上面dialog的visible.sync属性 -->
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

                        <div class="common-tabel">
                            <el-table
                                stripe
                                height="90%"
                                :data="tableData"
                                style="width: 100%">
                                <el-table-column
                                    prop="account"
                                    label="账号">
                                </el-table-column>
                                <el-table-column
                                    prop="name"
                                    label="用户名">
                                </el-table-column>
                                <el-table-column
                                    prop="real_name"
                                    label="真实姓名">
                                </el-table-column>
                                <el-table-column
                                    prop="remainder"
                                    label="剩余时长">
                                </el-table-column>
                                <el-table-column
                                    label="操作"
                                    width="280">
                                    <template slot-scope="scope">
                                        <el-button size="mini" @click="handleEdit(scope.row)">重置密码</el-button>
                                        <el-button size="mini" @click="editRemainder(scope.row)">重置额度</el-button>
                                        <el-button type="danger" size="mini" @click="handleDelete(scope.row)">删除</el-button>
                                    </template>
                                </el-table-column>
                            </el-table>
                        </div>
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
            form: {
                account: '',
                name: '',
                real_name: ''
            },
            rules: {
                account: [
                    { required: true, message: '请输入账号' }
                ],
                name: [
                    { required: true, message: '请输入用户名' }
                ],
                real_name: [
                    { required: true, message: '请输入真实姓名' }
                ]
            },
            tableData: []
        }
    },
    components: {
        ManageAside,
        CommonHeader
    },
    mounted() {
    // 在进入页面后立即调用的操作
    this.getUserData();
    },
    methods: {
        getUserData() {
        // 调用接口，获取使用记录
        $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/admin/show_user',
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
        // 将返回的数据格式转换为 el-table 所需的格式
        this.tableData = data.map((item) => ({
            'account': item['user_id'],
            'name': item['user_name'],
            'remainder': item['remainder'],
            'real_name': item['real_name'],
        }));
        },
        submit() {
            this.$refs.form.validate((valid) => { //this.$refs.form获取表单实例 validate用来校验表单 valid是返回值 bool类型
                // console.log(valid, 'valid')
                if (valid) { //valid为真 校验通过
                    // 后续对表单数据的处理
                    $.ajax({
                        type: 'POST',
                        url: 'http://127.0.0.1:8081/admin/add_user',
                        data: {
                            user_id: this.form.account,
                            user_name:this.form.name,
                            real_name:this.form.real_name,
                        },
                        success: (response) => {
                        // 请求成功的处理逻辑
                        console.log(response);
                        alert("恭喜您添加账户成功，密码初始化为123456！")
                        },
                        error: (xhr, status, error) => {
                        // 请求失败的处理逻辑
                        console.log('Error:', error);
                        alert("添加账户失败！")
                        },

                    });
                    location.reload()
                    //关闭窗口
                    handleClose()
                }
            })
        },
        // 弹窗关闭的时候
        handleClose() {
            this.$nextTick(() => {
                this.form = {
                account: '',
                password: ''
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
        editRemainder(row){
            console.log(row)
            $.ajax({
                type: 'POST',
                url: 'http://127.0.0.1:8081/admin/change_remainder',
                data: {
                    user_id: row.account
                },
                success: (response) => {
                // 请求成功的处理逻辑
                this.getUserData();
                console.log(response);
                alert("恭喜您！该账户剩余额度已重置为10h！")
                },
                error: (xhr, status, error) => {
                // 请求失败的处理逻辑
                this.getUserData();
                console.log('Error:', error);
                alert("重置额度失败！")
                },
            });
            location.reload();
        },
        handleEdit(row) {
            console.log(row)
            $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/admin/init_password',
            data: {
                user_id: row.account
            },
            success: (response) => {
            // 请求成功的处理逻辑
            console.log(response);
            alert("恭喜您！该账户密码已重置为123456！")
            },
            error: (xhr, status, error) => {
            // 请求失败的处理逻辑
            console.log('Error:', error);
            alert("重置密码失败！")
            },

        });
        },
        handleDelete(row){
            console.log(row)

            this.$confirm('此操作将永久删除该用户, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
            $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/admin/delete_user',
            data: {
                user_id: row.account
            },
            success: (response) => {
            // 请求成功的处理逻辑
            console.log(response);
            alert("删除成功！")
            location.reload()
            },
            error: (xhr, status, error) => {
            // 请求失败的处理逻辑
            console.log('Error:', error);
            alert("删除失败！")
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
       } 
    }
}
</script>

<style lang="less" scoped>
.el-header {
    padding: 0px;
}
.usermanage {
    height: 90%;
    .usermanage-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    .common-tabel {
        position: relative;
        height: calc(100% - 62px);
    }
}
</style>