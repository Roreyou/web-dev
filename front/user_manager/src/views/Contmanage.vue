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
                                <el-form-item label="容器ID" prop="ID">
                                    <el-input placeholder="请输入容器ID" v-model="form.ID"></el-input>
                                </el-form-item>

                                <el-form-item label="容器IP" prop="IP"> <!-- prop设置的是字段名字 用于校验 对于rules中的名字 -->
                                    <el-input placeholder="请输入容器IP" v-model="form.IP"></el-input> <!-- v-model是双向绑定 -->
                                </el-form-item>

                                <el-form-item label="端口" prop="Port">
                                    <el-input placeholder="请输入容器端口号" v-model="form.Port"></el-input>
                                </el-form-item>

                                <el-form-item label="状态" prop="Status">
                                    <template>
                                        <el-select v-model="form.Status" placeholder="请选择容器状态">
                                            <el-option v-for="item in options" :key="item.value" :label="item.value"
                                                :value="item.value">
                                            </el-option>
                                        </el-select>
                                    </template>
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

                        <el-table stripe height="90%" :data="TableDataa" style="width: 100%">
                            <el-table-column prop="ID" label="容器ID" width="150">
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
                                    <el-button size="mini" @click="handleEdit(scope.row)">编辑</el-button>
                                    <!-- (scope.row)是传参 当前的行数 -->
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
                IP: '',
                Port: '',
                Status: '',
                //Usetime: []
            },
            rules: {
                ID: [
                    { required: true, message: '请输入容器ID' } //message是错误时出现的信息 每一个{}代表一条规则
                ],
                IP: [
                    { required: true, message: '请输入容器IP' }
                ],
                Port: [
                    { required: true, message: '请输入容器端口号' }
                ],
                Status: [
                    { required: true, message: '请输入容器状态' }
                ],
                // Usetime: [
                //     { required: true, message: '请输入容器使用时间' }
                // ]
            },
            options: [{ value: '创建' }, { value: '正在使用' }, { value: '删除' }],
            TableDataa: [{
                ID: '39',
                IP: '192.168.0.1',
                Port: '3306',
                Status: '创建',
                //Usetime: []
            }, {
                ID: '45',
                IP: '192.168.0.3',
                Port: '3355',
                Status: '正在使用',
                //Usetime: ['2022.12.15-2022.12.19', '2023.4.5-2023.4.8']
            }, {
                ID: '66',
                IP: '192.168.3.2',
                Port: '3746',
                Status: '删除',
                //Usetime: ['2022.12.25-2022.12.26']
            }, {
                ID: '88',
                IP: '192.168.4.5',
                Port: '3852',
                Status: '正在使用',
                //Usetime: ['2022.12.6-2022.12.9', '2023.2.5-2023.2.9']
            }]
        }
    },
    components: {
        ManageAside,
        CommonHeader
    },
    methods: {
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

                    this.handleClose();
                }
            })
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

        // tagClose(tag) {
        //     this.form.Usetime.splice(this.form.Usetime.indexOf(tag), 1);
        // },

        // tagInput() {
        //     this.inputVisible = true;
        //     this.$nextTick(_ => {
        //         this.$refs.saveTagInput.$refs.input.focus();
        //     });
        // },

        // tagInputConfirm() {
        //     let inputValue = this.inputValue;
        //     if (inputValue) {
        //         this.form.Usetime.push(inputValue);
        //     }
        //     this.inputVisible = false;
        //     this.inputValue = '';
        // }
    }
}
</script>

<style lang="less" scoped>
.el-header {
    padding: 0px;
}
// .el-tag+.el-tag {
//     margin-left: 10px;
// }

// .button-new-tag {
//     margin-left: 10px;
//     height: 32px;
//     line-height: 30px;
//     padding-top: 0;
//     padding-bottom: 0;
// }

// .input-new-tag {
//     width: 90px;
//     margin-left: 10px;
//     vertical-align: bottom;
// }
</style>