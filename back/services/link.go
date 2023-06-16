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
			ssh.Password("ssezhangneng@972"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接SSH服务器
	sshClient, err := ssh.Dial("tcp", "ssh.schoolresearch.one:21006", sshConfig)
	if err != nil {
		panic(err)
	}
	defer sshClient.Close()
	fmt.Println("--------连接成功-------------")
	// 创建新的会话
	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 在远程服务器上运行Docker命令来创建一个新的容器
	cmd := "docker create --gpus " + `"device=2"` + " -it -p 20000:22 --name tensorflow_test 7e84df504bd9 bash"
	fmt.Println(cmd)
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println("------------操作----------成功")
	cid := string(output)
	cid = cid[1:12]
	fmt.Println(cid)
	return 0
}
