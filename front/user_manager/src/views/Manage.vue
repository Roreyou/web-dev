<template>
    <el-card class="box-card">
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
            tableData:[{
            ID: '163',
            Type: 'RTX3080',
            SSH: '50501',
            State: true
          }, {
            ID: '164',
            Type: 'RTX3081',
            SSH: '50502',
            State: true
          }, {
            ID: '165',
            Type: 'RTX3082',
            SSH: '50503',
            State: true
          }, {
            ID: '166',
            Type: 'RTX3083',
            SSH: '50504',
            State: true
          }]
        }
    },
    mounted() {
    // 在进入页面后立即调用的操作
    this.getUseGPUData();
    },
    methods: {
      getUseGPUData() {
        // 调用接口，获取使用记录
        $.ajax({
            type: 'POST',
            url: 'http://127.0.0.1:8081/show',
            data:{
              user_id:sessionStorage.getItem('user_id')
            },
            success: (response) => {
            // 请求成功的处理逻辑
            console.log(response);
            // 渲染数据到 el-table
            //this.renderTable(response);
            },
            error: (xhr, status, error) => {
            // 请求失败的处理逻辑
            console.log('Error:' + error);
            },
        });
      },
      handleDelete(index, row) {
        console.log(index, row);
        alert(index, row)
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