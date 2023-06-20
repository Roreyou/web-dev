package services

import (
	"back/dao"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

var c int

func IfuserTimeout(uid string) bool { //检查用户是否超额
	db := dao.Openmysql()
	var ur dao.User_Info
	db.Table("user_info").Where("user_id = ?", uid).First(&ur)
	ut := ur.Remainder
	uh := strings.Split(ut, "h")
	h, _ := strconv.Atoi(uh[0])
	um := strings.Split(uh[1], "m")
	m, _ := strconv.Atoi(um[0])
	db.Table("container").Where("user_id = ? ", uid).Count(&c)
	return h == 0 && m < 5 //小于五分钟就超额了
}

func StorecontainerInfo(imd int, psw string, cid string, mid int, uid int) {
	db := dao.Openmysql()

	c := dao.Container{
		Image_ID:           imd,
		Container_password: psw,
		Container_port:     20000,
		Container_id:       cid,
		Machine_id:         mid,
		Container_status:   2,
		User_id:            uid,
		Container_ip:       "172.16.108.78",
	}
	if err := db.Create(&c).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
}

func GreateDocker(mid string, uid string, did string, pwd string) int {
	if IfuserTimeout(uid) {
		return 1
	}

	if c > 0 {
		return 3
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
	var iid string
	if did == "1" {
		iid = "7e84df504bd9"
	} else if did == "2" {
		iid = "czq/tf2.8:v1.1"
	}

	cmd := "docker create --gpus " + `"` + "device=" + mid + `"` + " -it -p 20000:22 --name tensorflow" + uid + " " + iid + " bash"
	commands := []string{
		cmd,
		"docker start tensorflow" + uid,
		"docker exec tensorflow" + uid + " bash -c " + `"` + "echo root:" + pwd + "|chpasswd" + `"`, //修改密码
		"docker stop tensorflow" + uid,
	}
	command := strings.Join(commands, "; ")
	fmt.Println(cmd)
	output, err := session.CombinedOutput(command)
	if err != nil {
		fmt.Println("创建容器失败", err)
		return 2
	}
	cid := string(output)
	cid = cid[0:12]
	fmt.Println(cid)
	if cid == "Error respon" {
		return 2
	}
	session.Close()
	fmt.Println("------------容器创建----------成功")
	sid, _ := strconv.Atoi(mid)
	iuid, _ := strconv.Atoi(uid)
	StorecontainerInfo(1, pwd, cid, sid, iuid)
	return 0
}
