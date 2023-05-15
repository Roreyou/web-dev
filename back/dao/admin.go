package dao

import (
	"github.com/jinzhu/gorm"
)

func add_user(db *gorm.DB) {
	/*db.AutoMigrate(&User_Info{})
	//插入:
	hdmpassword := "hdm1"
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

	fmt.Println("完成！")*/
}
