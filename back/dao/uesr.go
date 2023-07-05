package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User_Info struct {
	User_id       int    `json:"user_id" db:"user_id"`
	User_Name     string `json:"user_name" db:"user_name"`
	Real_name     string `json:"real_name" db:"real_name"`
	User_password string `json:"user_password" db:"user_password"`
	Remainder     string `json:"remainder" db:"remainder"`
}

type Admin_Info struct {
	Admin_id       int    `json:"user_id"`
	Admin_Name     string `json:"user_name"`
	Admin_RealName string `json:"real_Name"`
	Admin_Password string `json:"user_password"`
}

func FindUser(userid int, db *gorm.DB) (user User_Info) { //查找到用户名相同的用户
	var saveInfo User_Info
	//查找输入用户名是否存在
	tableName := db.Model(&User_Info{})
	fmt.Println("tablename:", tableName) //
	db.Table("user_info").Where("user_id = ?", userid).Take(&saveInfo)

	fmt.Printf("u:%#v\n", saveInfo)
	return saveInfo
}

func GetUserInfo(c *gin.Context) User_Info {
	db := Openmysql()
	userid_string := c.PostForm("user_id") //返回的是string类型

	user_id, _ := strconv.Atoi(userid_string) //要转化成int类型

	get_info := FindUser(user_id, db)

	return get_info
}

func FindUserPassword(userid int, db *gorm.DB) string { //查找到用户名相同的用户
	var saveInfo User_Info
	//查找输入用户名是否存在
	db.Table("user_info").Where("user_id = ?", userid).Take(&saveInfo)
	fmt.Printf("%+v\n", saveInfo)
	print("\n")
	print("*old in database:", saveInfo.User_password)

	return saveInfo.User_password
}

func ChangeUserPassword(c *gin.Context) bool {
	db := Openmysql()
	//add_user(db)
	/*
		keywords := c.PostFormArray("keywords")

		print(keywords[0], "    ")
		print(keywords[1], "    ")
		print(keywords[2], "    ")
		//userid_string := c.PostForm("id") //返回的是string类型
	*/
	userid_string := c.PostForm("user_id")
	get_olduserPassword := c.PostForm("user_password")
	user_id, _ := strconv.Atoi(userid_string) //要转化成int类型
	old_userPassword := FindUserPassword(user_id, db)

	hashedPassword := sha256.Sum256([]byte(get_olduserPassword))
	get_hashedPassword := hex.EncodeToString(hashedPassword[:])

	if get_hashedPassword != old_userPassword {
		return false
	}
	//能执行到这里说明输入的原密码与数据库一致，允许修改密码

	new_Password := c.PostForm("new_password")
	new_hashedPassword := sha256.Sum256([]byte(new_Password))
	new_dbpassword := hex.EncodeToString(new_hashedPassword[:])
	err := db.Table("user_info").Where("user_id = ?", user_id).Update("user_password", new_dbpassword).Error

	if err != nil {
		panic(err)
	} else {
		return true
	}

}

func FindRecord(userid int64, db *gorm.DB) (userecord []Used_Record) {
	var record []Used_Record
	//查找输入用户名是否存在
	db.Model(&Used_Record{})
	db.Table("Used_Record").Where("user_id = ?", userid).Find(&record)
	fmt.Printf("u:%#v\n", record)
	return record
}

func Get_Recording(c *gin.Context) []Used_Record {
	db := Openmysql()
	// add_time()                                            //测试
	userid_string := c.PostForm("id")                     //返回的是string类型
	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型
	recording_info := FindRecord(user_id, db)
	return recording_info
}

func add_time() {
	db := Openmysql()
	db.Model(&Used_Record{})
	var record Used_Record
	db.Table("used_record").Where("user_id= ?", 1).First(&record)
	dtime := record.End_time.Sub(record.Start_time).String()
	db.Table("used_record").Where("user_id= ?", 1).Update("rent_time", dtime)
	fmt.Println("时间间隔：", dtime)
	/*
		tt1 := time.Now()
		db.Table("used_record").Where("user_id= ?", 1).Update("start_time", tt1)
		db.Table("used_record").Where("user_id= ?", 1).Update("end_time", tt1)
		db.Table("used_record").Where("user_id= ?", 2).Update("start_time", tt1)
		db.Table("used_record").Where("user_id= ?", 2).Update("end_time", tt1)
	*/
}

func Cut_Remainder(record Used_Record) {
	db := Openmysql()
	db.Model(&User_Info{})
	var user User_Info
	db.Table("user_info").Where("user_id= ?", record.User_id).First(&user)
	remainder, _ := time.ParseDuration(user.Remainder)
	new_remainder := remainder - (record.End_time.Sub(record.Start_time))
	db.Table("user_info").Where("user_id= ?", record.User_id).Update("remainder", new_remainder.String())
}
