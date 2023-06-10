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

                            <el-form ref="form" :rules="rules" :model="form" label-width="80px"> //model和定义的form绑定
                                <el-form-item label="账号" prop="account"> <!-- prop设置的是字段名字 用于校验 对于rules中的名字 -->
                                    <el-input placeholder="请输入账号" v-model="form.account"></el-input> <!-- v-model是双向绑定 -->
                                </el-form-item>
                                <el-form-item label="密码" prop="password">
                                    <el-input placeholder="请输入密码" v-model="form.password"></el-input>
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
                                    prop="password"
                                    label="密码">
                                </el-table-column>
                                <el-table-column
                                    label="操作">
                                    <template slot-scope="scope">
                                        <el-button size="mini" @click="handleEdit(scope.row)">重置密码</el-button>
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
                password: ''
            },
            rules: {
                account: [
                    { required: true, message: '请输入姓名' }
                ],
                password: [
                    { required: true, message: '请输入年龄' }
                ]
            },
            tableData: [
                {
                    account: '21311411',
                    password: '88888888'
                },
                {
                    account: '21311391',
                    password: '66666666'
                },
                {
                    account: '21311242',
                    password: '99999999'
                },
                {
                    account: '21311000',
                    password: '00000000'
                },
            ]
        }
    },
    components: {
        ManageAside,
        CommonHeader
    },
    methods: {
        submit() {
            this.$refs.form.validate((valid) => { //this.$refs.form获取表单实例 validate用来校验表单 valid是返回值 bool类型
                // console.log(valid, 'valid')
                if (valid) { //valid为真 校验通过
                    // 后续对表单数据的处理
                    if (this.modalType === 0) { //是否是新增 而不是编辑
                        console.log('新增')
                    } else {
                        console.log('修改')
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
        async handleEdit(row) {
            try {
            const response = await this.$axios.post("/admin/init_password", {
            user_id: this.user_id,
            user_password: this.user_password,
            });
            if (response.data === "fail") {
            this.errorMessage =
                "Login failed. Please check your user ID and password.";
            } else {
            const token = response.data;
            console.log("Login successful. Token:", token);
            // 你可以在这里将token存储到localStorage或者Vuex中，并执行其他登录成功后的操作
            }
        } catch (error) {
            console.error("Error while sending login request:", error);
        }
        },
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