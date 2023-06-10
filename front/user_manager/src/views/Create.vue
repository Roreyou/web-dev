<template>
    <div>
      <el-card class="bigbox-card">
        <el-card class="box-card" v-for="item in GPUData" :key="item.index">
          <div slot="header" class="clearfix">
            <span>{{ "GPU型号：" + item.name }}</span>
          </div>
          <div class="text item">
            {{ "显存大小：" + item.size }}
          </div>
          <div>
            <el-radio
              v-model="selectedGPUIndex"
              :label="item.index"
              :disabled="value !== '' && item.index !== value"
              border
            >
              备选项1
            </el-radio>
          </div>
        </el-card>

        <div class="text2">
          选择预装软件框架
          <br>
          <el-select v-model="value" clearable placeholder="请选择">
            <el-option
              v-for="option in selectedGPUFrameList"
              :key="option"
              :label="option"
              :value="option"
            ></el-option>
          </el-select>
        </div>
        

        <div class="text2">
          设置服务器登录密码(弱密码有资料泄露风险 请务必设置复杂密码)  
          <el-input placeholder="请输入密码" v-model="input" show-password></el-input>
        </div>

        <div class="text2">
          <el-checkbox v-model="checked">我已阅读并同意使用须知</el-checkbox>
        </div>

        <el-button type="warning" round>创建</el-button>
      </el-card>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        checked: false,
        input: '',
        value: '',
        selectedGPUIndex: '',
        GPUData: [
          {
            index: '1',
            name: 'RTX3090',
            size: '24G',
            frame_list: ['TensorFlow 2.8.0+ PyTorch 1.11.0+ conda/mamba + opencv 4+ CUDA11.6/cudnn8.4.0+ Python3.9.12'],
          },
          {
            index: '2',
            name: 'PG500-216(V100-32GB)',
            size: '32G',
            frame_list: ['TensorFlow 2.8.0 + PyTorch 1.11.0 + conda/mamba + opency 4 + CUDA11.6/cudnn8.4.0 + Python3.9.12', 'TensorFlow 1.15.5 + PyTorch 1.11.0 + conda/mamba + opencv 4 + CUDA11.6/cudnn8.4.0 + Python3.8'],
          },
        ],
      };
    },
    computed: {
      selectedGPUFrameList() {
        const selectedGPU = this.GPUData.find(item => item.index === this.selectedGPUIndex);
        return selectedGPU ? selectedGPU.frame_list : [];
      },
    },
  };
  </script>
  
  <style>
  .text {
    font-size: 14px;
  }

  .text2 {
    font-size: 13px;
    color: #ffd04b;
    margin-top: 20px;
    margin-bottom: 14px;
  }
  
  .item {
    margin-bottom: 18px;
  }
  
  .clearfix:before,
  .clearfix:after {
    display: table;
    content: '';
  }
  
  .clearfix:after {
    clear: both;
  }
  
  .box-card {
    width: 80%;
    margin-bottom: 20px;
  }
  </style>
  