<template>
    <el-card class="box-card">
        <div slot="header" class="clearfix">
            <span>使用记录</span>
        </div>
        <div class="text item">
            <el-table
                :data="tableData"
                style="width: 100%">
                <el-table-column
                    prop="租用的机器ID"
                    label="机器ID"
                    width="200px">
                </el-table-column>
                <el-table-column
                    prop="机器开始使用时间"
                    label="开始时间"
                    width="200px">
                </el-table-column>
                <el-table-column
                    prop="机器结束使用时间"
                    label="结束时间"
                    width="200px">
                </el-table-column>
                <el-table-column
                    prop="本服务器累计租用时长"
                    label="本次租用总时长"
                    width="200px">
                </el-table-column>
                <el-table-column
                    prop="总共租用服务器时间"
                    label="累计租用总时长"
                    width="200px">
                </el-table-column>
            </el-table>
        </div>
    </el-card>
</template>

<script>

import $ from 'jquery'

export default {
    data() {
        return {
            tableData:[]
        }
    },
    mounted() {
    // 在进入页面后立即调用的操作
    this.getGPUdata();
  },
  methods: {
    getGPUdata() {
      // 调用接口，获取使用记录
      $.ajax({
        type: 'POST',
        url: 'http://127.0.0.1:8081/user/recording',
        data: {
          // 请求的参数
          id: sessionStorage.getItem('user_id'),
        },
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
        '使用者ID': item['使用者ID'],
        '租用的机器ID': item['租用的机器ID'],
        '机器开始使用时间': item['机器开始使用时间'],
        '机器结束使用时间': item['结束租用服务器时间'],
        '本次租用服务器的时间': item['本次租用服务器的时间'],
        '总共租用服务器时间': item['总共租用服务器时间'],
      }));
    },
  },
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
