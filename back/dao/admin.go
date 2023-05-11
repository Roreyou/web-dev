package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/jinzhu/gorm"
)

func add_user(db *gorm.DB) {
	db.AutoMigrate(&User_Info{})
	//插入:
	hdmpassword := "hdmpassword2"
	lxhpassword := "lxhpassword2"

	hashedPassword_hdm := sha256.Sum256([]byte(hdmpassword))
	get_hdmPassword := hex.EncodeToString(hashedPassword_hdm[:])

	hashedPassword_lxh := sha256.Sum256([]byte(lxhpassword))
	get_lxhPassword := hex.EncodeToString(hashedPassword_lxh[:])

	u1 := User_Info{3, "小明2", "黄大明2", get_hdmPassword}
	u2 := User_Info{4, "小红2", "刘大红2", get_lxhPassword}
	// 创建记录

	db.Table("user_info").Create(&u1)
	db.Table("user_info").Create(&u2)

	fmt.Println("完成！")
}
