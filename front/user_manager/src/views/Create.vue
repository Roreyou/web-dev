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
              border
            >
              选择
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

        <el-button type="warning" @click="createGPU" round>创建</el-button>
      </el-card>
    </div>
  </template>
  
  <script>
    import $ from 'jquery'
  export default {
    data() {
      return {
        checked: false,
        input: '',
        value: '',
        selectedGPUIndex: '',
        GPUData: [],
        selectedGPUFrameList:['tensorflow2.1.0+pythorch1.11+python3.6.9','tensorflow2.6.0+pythorch1.11+python3.9.0'],
      };
    },
    mounted() {
    // 在进入页面后立即调用的操作
    this.getAccessGPUData();
    },
    methods:{
      getAccessGPUData() {
        // 调用接口，获取使用记录
        $.ajax({
            type: 'GET',
            url: 'http://127.0.0.1:8081/user/get_item',
            success: (response) => {
            // 请求成功的处理逻辑
            console.log(response);
            // 渲染数据到 el-table
            this.renderTable(response);
            },
            error: (xhr, status, error) => {
            // 请求失败的处理逻辑
            console.log('Error:');
            },
        });
      },
      renderTable(data) {
        // 将返回的数据格式转换为 el-table 所需的格式
        this.GPUData = data.map((item) => ({
            'index': item['server_id'],
            'name': item['sever_type'],
            'size': item['server_size'],
        }));
      },
      createGPU(){
        if(this.input!==''&&this.value!==''&&this.selectedGPUIndex!==''&&this.checked===true){
          console.log("byebye!")
          var new_value="0"
          if(this.value==='tensorflow2.1.0+pythorch1.11+python3.6.9'){
            new_value="1"
          }
          else{
            new_value="2"
          }
          console.log("user_id="+sessionStorage.getItem('user_id'))
          console.log("server_id="+this.selectedGPUIndex)
          console.log("image_id="+new_value)
          console.log("user_password="+this.input)
          $.ajax({
                type: 'POST',
                url: 'http://127.0.0.1:8081/user/create_gpu',
                data: {
                     // 请求的参数
                    user_id: sessionStorage.getItem('user_id'),
                    server_id:this.selectedGPUIndex,
                    image_id:new_value,
                    user_password:this.input,
                },
                success: (response) => {
                // 请求成功的处理逻辑
                console.log(response);
                alert("恭喜您创建GPU成功！")
                },
                error: (xhr, status, error) => {
                // 请求失败的处理逻辑
                console.log('Error:', error);
                alert("创建GPU失败！")
                },
           });
        }
        else{
          alert("请完善相关数据！")
        }
      }
    },
    computed: {
      // selectedGPUFrameList() {
      //   const selectedGPU = this.GPUData.find(item => item.index === this.selectedGPUIndex);
      //   return selectedGPU ? selectedGPU.frame_list : [];
      // },
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
  