<template>
    <div class="header-container">
        <div class="l-content">
            <el-button @click="handleMenu" icon="el-icon-menu" size="mini"></el-button>
            <span class="text">首页</span>
        </div>
        <div class="r-content">
            <span class="welcome">你好，{{ user_name }}</span>
            <el-dropdown trigger="click" @command="handleClick">
                <span class="el-dropdown-link">
                    下拉菜单<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item icon="el-icon-edit-outline" command="changePW">修改密码</el-dropdown-item>
                    <el-dropdown-item icon="el-icon-back">退出登录</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
        //获取保存的用户名，显示在“欢迎XXX"中
        user_name:sessionStorage.getItem("user_name"),
        }
        
    },
    methods: {
        handleMenu() {
            this.$store.commit('collapseMenu')
        },
        handleClick(command) {
            if (command === 'changePW') {
                this.$router.push('change')
            }
        }
    }
}
</script>
<style lang="less" scoped>
.header-container {
    padding: 0 20px;
    background-color: #333;
    height: 60px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .text {
        color: #fff;
        font-size: 14px;
        margin-left: 10px;
    }
    .el-dropdown-link {
        cursor: pointer;
        color: #fff;
        font-size: 14px;
        margin-left: 10px;
        font-weight: 600;
    }
    .welcome {
        color: #fff;
        font-size: 14px;
        font-weight: 200;
    }
}
</style>