package dao

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	User_id      int64  `json:user_id"`
	User_Name    string `json:User_name"`
	Real_Name    string `json:real_Name`
	Password     string `json:password"`
	Expired_time string `json:expired_time`
}

func FindUser(userid int64) (user UserInfo) { //查找到用户名相同的用户
	var saveInfo UserInfo
	//查找输入用户名是否存在
	db := Openmysql()
	db.Where("user_id = ?", userid).Take(&saveInfo)
	print("++++++++++++++++++++")
	fmt.Printf("u:%#v\n", saveInfo)
	return saveInfo
}

func GetUserInfo(c *gin.Context) *UserInfo {
	db, err := gorm.Open("mysql", "root:0921@(127.0.0.1:3306)/web_database?charset=utf8mb4&parseTime=True&loc=Local")
	// db就是database
	if err != nil {
		panic(err)
	}
	defer db.Close() //把数据库的连接关闭掉

	db.AutoMigrate(&UserInfo{})

	//插入:
	now := time.Now()
	nowStr1 := now.Format("2006-01-02 15:04:05")

	u1 := UserInfo{1, "小明", "黄大明", "hdmpassword", nowStr1}
	u2 := UserInfo{2, "小红", "刘大红", "lxhpassword", nowStr1}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	fmt.Println("完成！")

	userid_string := c.PostForm("id") //返回的是string类型
	print("user_id----------------:" + userid_string)
	print("----------------")
	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型
	get_info := FindUser(user_id)
	/*
		if get_info.User_id == 0 { //该用户不存在则报错
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"fail": "该用户名不存在",
			})
			return
		}*/
	return &get_info

	/*
		now := time.Now()
		nowStr1 := now.Format("2006-01-02 15:04:05")


		u1 := AdminInfo{1, "小明", "黄大明", "hdmpassword", nowStr1}
		u2 := AdminInfo{2, "小红", "刘大红", "lxhpassword", nowStr1}
		// 创建记录
		db.Create(&u1)
		db.Create(&u2)
		fmt.Println("完成！")

		return &u1 //用来测试一下 返回的是u1数据
	*/
}
