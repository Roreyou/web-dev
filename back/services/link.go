package services

import (
	"back/dao"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func IfuserTimeout(uid string) bool {
	db := dao.Openmysql()
	var ur dao.Used_Record
	fmt.Println(uid)
	db.Table("used_record").Where("user_id = ?", uid).First(&ur)
	ut := ur.Used_time
	uh := strings.Split(ut, "h")
	h, _ := strconv.Atoi(uh[0])
	fmt.Println(h)
	return h < 10
}

func GreateDocker(mid string, uid string, did string) int {
	if !IfuserTimeout(uid) {
		return 1
	}
	sshConfig := &ssh.ClientConfig{
		User: "zhangn279",
		Auth: []ssh.AuthMethod{
			ssh.Password("123456"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接SSH服务器
	sshClient, err := ssh.Dial("tcp", "ssh.schoolresearch.one:21110", sshConfig)
	if err != nil {
		panic(err)
	}
	defer sshClient.Close()
	fmt.Println("--------连接成功-------------")
	// 创建新的会话  15acc1083d3a
	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 在远程服务器上运行Docker命令来创建一个新的容器
	//cmd := "docker run --name my-container -itd your_image_name /bin/bash"
	output, err := session.CombinedOutput("docker -v")
	if err != nil {
		panic(err)
	}
	fmt.Println("------------操作----------成功")
	fmt.Println(string(output))
	return 0
}
