package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Add_user(c *gin.Context) bool {
	db := Openmysql()
	db.AutoMigrate(&User_Info{})
	//插入id:
	userid := c.PostForm("user_id")    //返回的是string类型
	user_id, _ := strconv.Atoi(userid) //要转化成int类型

	userpassword := "123456"
	//插入user_name
	user_name := c.PostForm("user_name") //返回的是string类型
	//c插入real_name
	real_name := c.PostForm("real_name")

	hashedPassword_hdm := sha256.Sum256([]byte(userpassword))
	save_Password := hex.EncodeToString(hashedPassword_hdm[:])

	u1 := User_Info{user_id, user_name, real_name, save_Password, "10h"}
	// 创建记录
	db.Table("user_info").Create(&u1)
	fmt.Println(u1)
	fmt.Println("完成！")

	return true
}

func Delete_usr(c *gin.Context) bool {
	db := Openmysql()
	db.AutoMigrate(&User_Info{})
	userid := c.PostForm("user_id")    //返回的是string类型
	user_id, _ := strconv.Atoi(userid) //要转化成int类型
	fmt.Println(user_id)
	// 查询要删除的用户是否存在
	var user User_Info
	result := db.Table("user_info").Where("user_id = ?", user_id).First(&user)
	fmt.Println(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	// 删除用户数据
	db.Debug().Table("user_info").Where("user_id = ?", user_id).Delete(&user)
	//db.Table("user_info").Delete(&user)
	db.Close()
	return true

}

func Init_password(c *gin.Context) bool {
	db := Openmysql()
	userid_string := c.PostForm("user_id")
	user_id, _ := strconv.Atoi(userid_string) //要转化成int类型
	user := FindUser(user_id, db)
	if user.User_id == 0 { //该用户不存在则报错
		return false
	}
	new_Password := "123456" //初始化密码
	fmt.Println("重置密码设置为：", new_Password)
	new_hashedPassword := sha256.Sum256([]byte(new_Password))
	new_dbpassword := hex.EncodeToString(new_hashedPassword[:])
	err := db.Table("user_info").Where("user_id = ?", user_id).Update("user_password", new_dbpassword).Error

	if err != nil {
		panic(err)
	} else {
		return true
	}
}

func Show_alluser(c *gin.Context) []User_Info {
	db := Openmysql()
	var saveInfo []User_Info
	db.Table("user_info").Find(&saveInfo)
	fmt.Println(saveInfo)
	return saveInfo
}

func FindAllCont() (cont []Container) {
	db := Openmysql()
	db.Model(&Container{})
	db.Table("container").Find(&cont)
	return cont
}
func FindAllMachine() (server []Server_Info) { //返回服务器信息
	db := Openmysql()
	defer db.Close() //把数据库的连接关闭掉
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&Server_Info{})
	db.Table("server_info").Find(&server)
	return server
}

func AddMachine(sid string, server_type string, size string) {
	server_id, _ := strconv.ParseInt(sid, 10, 64) //要转化成int64类型
	server_size, _ := strconv.ParseInt(size, 10, 64)
	db := Openmysql()
	s := Server_Info{
		Server_id:    int(server_id),
		Server_type:  server_type,
		Server_size:  int(server_size),
		Server_state: 0,
		Server_flag:  false,
	}
	if err := db.Create(&s).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
}

func DeleteMachine(sid string) {
	server_id, _ := strconv.ParseInt(sid, 10, 64) //要转化成int64类型
	db := Openmysql()
	db.AutoMigrate(&Server_Info{})
	db.Table("server_info").Where("server_id = ?", server_id).Delete(&Server_Info{})
}
