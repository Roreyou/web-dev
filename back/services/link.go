package services

import (
	"back/dao"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

var c int

func IfuserTimeout(uid string) bool { //检查用户是否超额
	db := dao.Openmysql()
	var ur dao.Used_Record
	db.Table("used_record").Where("user_id = ?", uid).First(&ur)
	ut := ur.Used_time
	uh := strings.Split(ut, "h")
	h, _ := strconv.Atoi(uh[0])
	fmt.Println(h)
	db.Table("container").Count(&c)
	fmt.Print("c:")
	fmt.Println(c)
	return h < 10
}

func Repasswd(psw string) { //修改容器密码
	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("123456"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// 连接SSH服务器
	sshClient, err := ssh.Dial("tcp", "172.16.108.78:20000", sshConfig)
	if err != nil {
		panic(err)
	}
	defer sshClient.Close()
	fmt.Println("--------容器连接成功-------------")
	// 创建新的会话
	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// 远程连接docker并修改密码
	commands := []string{
		"echo root:" + psw + " | chpasswd",
		"exit",
	}
	command := strings.Join(commands, "; ")

	if err := session.Run(command); err != nil {
		panic("Failed to login to swarm registry: " + err.Error())
	}
	fmt.Println("------------密码修改----------成功")
	session.Close()

}

func StorecontainerInfo(imd int, psw string, cid string, mid int) {
	db := dao.Openmysql()
	c := dao.Container{
		Image_ID:           imd,
		Container_password: psw,
		Container_port:     20000,
		Container_id:       cid,
		Machine_id:         mid,
		Container_status:   2,
		User_id:            1,
	}
	if err := db.Create(&c).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
}

func GreateDocker(mid string, uid string, did string, pwd string) int {
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
	tid := fmt.Sprintf("%d", c)
	cmd := "docker create --gpus " + `"` + "device=" + mid + `"` + " -it -p 20000:22 --name tensorflow" + tid + " 7e84df504bd9 bash"
	commands := []string{
		cmd,
		"docker start tensorflow" + tid,
		//"docker ps",
	}
	command := strings.Join(commands, "; ")
	fmt.Println(cmd)
	output, err := session.CombinedOutput(command)
	if err != nil {
		panic(err)
	}
	fmt.Println("------------容器创建----------成功")
	cid := string(output)
	cid = cid[0:12]
	fmt.Println(cid)
	session.Close()
	time.Sleep(1.5 * 1000 * time.Millisecond)
	Repasswd(pwd)
	sid, _ := strconv.Atoi(mid)
	StorecontainerInfo(1, pwd, cid, sid)
	Tail(tid)
	return 0
}

func Tail(tid string) {
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
	if err := session.Run("docker stop tensorflow" + tid); err != nil {
		panic(err)
	}
	defer sshClient.Close()
}
