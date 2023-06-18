package controller

import (
	"back/dao"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

// 容器状态：
// 0表示容器未创建
// 1表示容器已删除
// 2表示容器已创建且未被使用
// 3表示容器正在使用
func ConnectContainer(c *gin.Context) {
	container, is_enter := dao.UseContainer(c) //密码输入是否正确
	if !is_enter {                             //判断是否能进入
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码不正确",
		})
	} else {
		if container.Container_status == 1 { //容器被删除
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "时间已超额",
			})
		} else if flag := dao.IsContainerUsing(); flag == true { //容器正在使用
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"warining": "该容器正在被使用",
			})
		} else { //容器存在
			//ssh连接信息
			config := &ssh.ClientConfig{
				User: "zhangn279", //服务器的账号
				Auth: []ssh.AuthMethod{
					ssh.Password("ssezhangneng@972"), //服务器密码
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			}
			//连接ssh服务器
			client, err := ssh.Dial("tcp", "172.16.108.78:2022", config)
			if err != nil {
				log.Fatal("Failed to dial: ", err)
			}
			defer client.Close()
			//str := strconv.Itoa(container.Image_ID) //将整数转为字符串
			str := container.Container_id
			//cmd := "docker attach " + str           //回到已经退出的但是仍然在运行的容器的命令
			cmd := "docker start " + str //启动容器
			// 创建新的会话
			session, err := client.NewSession()
			if err != nil {
				panic(err)
			}
			defer session.Close()
			// 在远程服务器上运行Docker命令，output就是容器信息
			output, err := session.CombinedOutput(cmd)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(output))
			start_time := time.Now()                //获取现在的时间
			dao.CreateRecord(container, start_time) //创建使用记录
			dao.UpdateContainerStatus(3, container) //容器正在使用
			c.JSON(http.StatusOK, gin.H{
				"msg": "容器正在使用中",
			})
		}
	}
}
func ExitContainer(c *gin.Context) { //关机
	//keywords := c.PostFormArray("keywords") //获取输入数据
	//userid_string := keywords[0]
	userid_string := c.PostForm("user_id")
	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型
	container := dao.FindContainer(user_id)               //找到对应容器
	if container.Container_status == 1 {                  //容器被删除
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "时间已超额",
		})
		return
	}
	config := &ssh.ClientConfig{
		User: "zhangn279", //服务器的账号
		Auth: []ssh.AuthMethod{
			ssh.Password("ssezhangneng@972"), //服务器密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	//连接ssh服务器
	client, err := ssh.Dial("tcp", "172.16.108.78:2022", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()
	//str := strconv.Itoa(container.Image_ID) //将整数转为字符串
	str := container.Container_id
	cmd := "docker stop " + str //关闭容器的命令
	//创建新的会话
	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// 在远程服务器上运行Docker命令，output就是容器信息
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	end_time := time.Now()
	outtime := dao.UpdateRecord(user_id, end_time)
	dao.UpdateContainerStatus(2, container)
	if outtime { //超时
		//删除容器禁止下次使用
		//ssh连接信息
		config := &ssh.ClientConfig{
			User: "zhangn279", //服务器的账号
			Auth: []ssh.AuthMethod{
				ssh.Password("ssezhangneng@972"), //服务器密码
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		//连接ssh服务器
		client, err := ssh.Dial("tcp", "172.16.108.78:2022", config)
		if err != nil {
			log.Fatal("Failed to dial: ", err)
		}
		defer client.Close()
		//cmd := "exit" //退出容器的命令
		//str := strconv.Itoa(container.Image_ID) //将整数转为字符串
		str := container.Container_id
		cmd := "docker rm " + str //删除容器
		//创建新的会话
		session, err := client.NewSession()
		if err != nil {
			panic(err)
		}
		defer session.Close()
		// 在远程服务器上运行Docker命令，output就是容器信息
		output, err := session.CombinedOutput(cmd)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(output))
		dao.UpdateContainerStatus(1, container) //将容器的状态表示已删除
		c.JSON(http.StatusOK, gin.H{
			"msg":     "容器已停止使用",
			"warning": "容器使用超出额度",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "容器已停止使用",
		})
	}
}
func DeleteContainer(c *gin.Context) { //删除服务器
	container, is_enter := dao.UseContainer(c)
	if !is_enter { //判断是否能进入
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码不正确",
		})
	} else {
		//ssh连接信息
		config := &ssh.ClientConfig{
			User: "zhangn279", //服务器的账号
			Auth: []ssh.AuthMethod{
				ssh.Password("ssezhangneng@972"), //服务器密码
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		//连接ssh服务器
		client, err := ssh.Dial("tcp", "172.16.108.78:2022", config)
		if err != nil {
			log.Fatal("Failed to dial: ", err)
		}
		defer client.Close()
		//cmd := "exit" //退出容器的命令
		//str := strconv.Itoa(container.Image_ID) //将整数转为字符串
		str := container.Container_id
		cmd := "docker rm " + str //删除容器
		//创建新的会话
		session, err := client.NewSession()
		if err != nil {
			panic(err)
		}
		defer session.Close()
		// 在远程服务器上运行Docker命令，output就是容器信息
		output, err := session.CombinedOutput(cmd)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(output))
		dao.UpdateContainerStatus(1, container) //将容器的状态表示已删除
		c.JSON(http.StatusOK, gin.H{
			"msg": "容器已删除",
		})
	}
}
