package dao

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User_Info struct {
	User_id       int64  `json:user_id, db:user_id`
	User_Name     string `json:User_name,db:user_name"`
	User_RealName string `json:real_Name,db:real_name`
	User_Password string `json:user_password",db:user_password`
}

type Admin_Info struct {
	Admin_id       int64  `json:user_id"`
	Admin_Name     string `json:User_name"`
	Admin_RealName string `json:real_Name`
	Admin_Password string `json:user_password"`
}

func FindUser(userid int64, db *gorm.DB) (user User_Info) { //查找到用户名相同的用户
	var saveInfo User_Info
	//查找输入用户名是否存在
	tableName := db.Model(&User_Info{})
	fmt.Println("tablename:", tableName) //
	db.Table("user_info").Where("user_id = ?", userid).Take(&saveInfo)

	fmt.Printf("u:%#v\n", saveInfo)
	return saveInfo
}

func GetUserInfo(c *gin.Context) *User_Info {
	db := Openmysql()
	userid_string := c.PostForm("id") //返回的是string类型

	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型

	get_info := FindUser(user_id, db)

	if get_info.User_id == 0 { //该用户不存在则报错
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "该用户名不存在",
		})
	}
	return &get_info
}

func FindUserPassword(userid int64, db *gorm.DB) string { //查找到用户名相同的用户
	var saveInfo User_Info
	//查找输入用户名是否存在
	db.Table("user_info").Where("user_id = ?", userid).Take(&saveInfo)
	fmt.Printf("%+v\n", saveInfo)
	print("\n")
	print("*old in database:", saveInfo.User_Password)
	return saveInfo.User_Password
}

func ChangeUserPassword(c *gin.Context) bool {
	db := Openmysql()

	keywords := c.PostFormArray("keywords")

	print(keywords[0], "    ")
	print(keywords[1], "    ")
	print(keywords[2], "    ")

	userid_string := keywords[0]
	get_olduserPassword := keywords[1]
	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型
	old_userPassword := FindUserPassword(user_id, db)

	print("old_userPassword: ", old_userPassword)
	fmt.Printf("\n")
	print(keywords[2])
	if get_olduserPassword != old_userPassword {
		return false
	}
	//能执行到这里说明输入的原密码与数据库一致，允许修改密码

	new_Password := keywords[2]

	err := db.Table("user_info").Where("user_id = ?", user_id).Update("user_password", new_Password).Error

	if err != nil {
		panic(err)
	} else {
		return true
	}

}
