<template>
    <el-menu default-active="1-4-1" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose"
      :collapse="isCollapse" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
  
      <h3>{{ isCollapse ? '后台' : 'GPU管理系统' }}</h3>
  
      <el-menu-item @click="clickMenu(item)" v-for="item in menuData" :key="item.name" :index="item.name">
        <i :class="`el-icon-${item.icon}`"></i>
        <span slot="title">{{ item.label }}</span>
      </el-menu-item>
  
    </el-menu>
  </template>
  
    
  <style lang="less" scoped>
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
  }
  
  .el-menu {
    height: 100vh;
    border-right: none;
  
    h3 {
      color: #fff;
      text-align: center;
      line-height: 48px;
      font-size: 19px;
      font-weight: 600;
    }
  }
  </style>
    
  <script>
  export default {
    data() {
      return {
        menuData: [
          {
            path: '/gpumanage',
            name: 'gpumanage',
            label: '服务器管理',
            icon: 'set-up',
            url: 'GPUManage/GPUManage'
          },
          {
            path: '/usermanage',
            name: 'usermanage',
            label: '用户管理',
            icon: 's-custom',
            url: 'UserManage/UserManage'
          },
           {
            path: '/contmanage',
            name: 'contmanage',
            label: '容器管理',
            icon: 's-custom',
            url: 'Contmanage/Contmanage'
          },
        ]
      };
    },
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      },
      // 点击菜单
      clickMenu(item) {
        console.log(item)
        // 当页面的路由与跳转的路由不一致才允许跳转
        if (this.$route.path !== item.path && !(this.$route.path === '/manage' && (item.path === '/'))) {
          this.$router.push(item.path)
        }
        this.$store.commit('selectMenu', item)
      }
    },
    computed: {
      isCollapse() {
        return this.$store.state.tab.isCollapse
      }
    }
  }
  </script>