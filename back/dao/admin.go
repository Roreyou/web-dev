package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func add_user(db *gorm.DB) {
	db.AutoMigrate(&User_Info{})
	//插入:
	hdmpassword := "hdm"
	lxhpassword := "lxh1"

	hashedPassword_hdm := sha256.Sum256([]byte(hdmpassword))
	get_hdmPassword := hex.EncodeToString(hashedPassword_hdm[:])

	hashedPassword_lxh := sha256.Sum256([]byte(lxhpassword))
	get_lxhPassword := hex.EncodeToString(hashedPassword_lxh[:])

	u1 := User_Info{1, "小明", "黄大明", get_hdmPassword}
	u2 := User_Info{2, "小红", "刘大红", get_lxhPassword}
	// 创建记录

	db.Table("user_info").Create(&u1)
	db.Table("user_info").Create(&u2)

	fmt.Println("完成！")
}

func Init_password(c *gin.Context) bool {
	db := Openmysql()
	userid_string := c.PostForm("user_id")
	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型
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
