<template>
    <el-card class="box-card">
        <div slot="header" class="clearfix">
            <span>使用记录</span>
        </div>
        <div class="text item">
            <el-upload
                class="upload-demo"
                drag
                action="https://jsonplaceholder.typicode.com/posts/"
                multiple>
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过50G</div>
            </el-upload>
        </div>
    </el-card>
</template>

<script>
export default {
    
}
</script>