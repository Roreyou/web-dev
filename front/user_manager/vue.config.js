const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  devServer: {
      port: 8080, // 定义开发服务器的端口号
    },
 
  
  })
// module.exports = {
//     devServer: {
//       port: 8080, // 定义开发服务器的端口号
//     },
//   };
  
// module.exports = {
//     transpileDependencies: true,
//     lintOnSave: false,
//     devServer: {
//       proxy: {
//         '/': {
//           target: 'http://localhost:8080', // 后端接口地址
//           changeOrigin: true,
//         //   pathRewrite: {
//         //     '^/api': '' // 如果后端接口的路径中有/api前缀，可以将其替换为空字符串
//         //   }
//         }
//       }
//     }
//   };

  
  
  
  
  
  
  






































































































































































































































