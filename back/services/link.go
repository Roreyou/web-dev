package services

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func SSHlink() *ssh.Client {
	sshConfig := &ssh.ClientConfig{
		User: "zhangn279",
		Auth: []ssh.AuthMethod{
			ssh.Password("ssezhangneng@972"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接SSH服务器
	sshClient, err := ssh.Dial("tcp", "ssh.schoolresearch.one:21110", sshConfig)
	if err != nil {
		panic(err)
	}
	defer sshClient.Close()
	/*session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()*/

	return sshClient
}

/*func IfmachineAvailable(id string) bool {
	db := dao.Openmysql()
	var gpul []dao.Gpu
	sql := "select server_id,server_type,server_size,server_state,server_flag from gpus where server_id=1 "
	err := db.Select(&gpul, sql, 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return true
	}
	fmt.Println("select succ:", gpul)
	return false
}*/

func GreateDocker(mid string) int {
	/*if IfmachineAvailable(mid) {
		return 1
	}
	/*	// SSH连接信息*/
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
	//sshClient := SSHlink()
	fmt.Println("--------连接成功-------------")
	// 创建新的会话  15acc1083d3a
	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//session := SSHlink()
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
