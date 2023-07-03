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
                                    <el-button v-else class="button-new-tag" size="small" @click="tagInput">+ New Tag</el-button>

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
                            height="90%"
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
                SSH: '',
                Frame: []
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
                SSH: [
                    { required: true, message: '请输入SSH命令' }
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
    methods: {
        // 提交用户表单
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
                ID: '',
                Type: '',
                Size: '',
                SSH: '',
                Frame: []
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