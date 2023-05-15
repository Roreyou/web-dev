package services

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func GreateDocker() int {
	// SSH连接信息

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

	// 创建新的会话
	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 在远程服务器上运行Docker命令来创建一个新的容器
	//cmd := "docker run --name my-container -itd your_image_name /bin/bash"
	output, err := session.CombinedOutput("ls -lh")
	if err != nil {
		panic(err)
	}
	fmt.Println("------------连接----------成功")
	fmt.Println(string(output))
	return 0
}
