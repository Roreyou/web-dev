package controller

import (
	"back/dao"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ConnectContainer(c *gin.Context) {
	container, is_enter := dao.UseContainer(c)

	if !is_enter { //判断是否能进入
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码不正确",
		})
	} else {
		if container.Container_status == 0 { //容器已被删除
			//重新创建容器
		} else {
			//ssh连接信息
			// config := &ssh.ClientConfig{
			// 	User: "your_username",
			// 	Auth: []ssh.AuthMethod{
			// 		ssh.Password("your_password"),
			// 	},
			// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			// }
			// //连接ssh服务器
			// client, err := ssh.Dial("tcp", "your_server_address:22", config)
			// if err != nil {
			// 	log.Fatal("Failed to dial: ", err)
			// }
			// defer client.Close()
			str := strconv.Itoa(container.Image_ID) //将整数转为字符串
			cmd := "docker attach " + str           //回到已经退出的但是仍然在运行的容器的命令
			// // 创建新的会话
			// session, err := client.NewSession()
			// if err != nil {
			// 	panic(err)
			// }
			// defer session.Close()
			// // 在远程服务器上运行Docker命令，output就是容器信息
			// output, err := session.CombinedOutput(cmd)
			// if err != nil {
			// 	panic(err)
			// }
			// fmt.Println(string(output))
			start_time := time.Now()                //获取现在的时间
			dao.CreateRecord(container, start_time) //创建使用记录
			c.JSON(http.StatusOK, gin.H{            //只负责给出命令
				"cmd": cmd,
			})
		}
	}
}
func ExitContainer(c *gin.Context) { //退出容器但容器仍在运行
	//ssh连接信息
	keywords := c.PostFormArray("keywords") //获取输入数据
	userid_string := keywords[0]
	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型

	// config := &ssh.ClientConfig{
	// 	User: "your_username",
	// 	Auth: []ssh.AuthMethod{
	// 		ssh.Password("your_password"),
	// 	},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }
	// //连接ssh服务器
	// client, err := ssh.Dial("tcp", "your_server_address:22", config)
	// if err != nil {
	// 	log.Fatal("Failed to dial: ", err)
	// }
	// defer client.Close()
	cmd := "exit" //退出容器的命令
	// 创建新的会话
	// session, err := client.NewSession()
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()
	// // 在远程服务器上运行Docker命令，output就是容器信息
	// output, err := session.CombinedOutput(cmd)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(output))
	end_time := time.Now()
	dao.UpdateRecord(user_id, end_time)
	c.JSON(http.StatusOK, gin.H{ //只负责给出命令
		"cmd": cmd,
	})
}
