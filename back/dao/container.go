package dao

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Used_Record struct {
	User_id    int       `json:"使用者ID"`
	Machine_id int       `json:"租用的机器ID"`
	Start_time time.Time `json:"机器开始使用时间"`
	End_time   time.Time `json:"结束租用服务器时间"`
	Rent_time  string    `json:"本次租用服务器的时间"`
	Used_time  string    `json:"总共租用服务器时间"`
}

type Container struct {
	Image_ID           int    `json:"镜像ID"`
	Container_password string `json:"容器密码"`
	Container_port     int    `json:"容器端口"`
	Container_id       string `json:"容器ID"`
	Machine_id         int    `json:"对应机器ID"`
	Container_status   int    `json:"容器的状态"`
	User_id            int    `json:"创建容器的用户"`
}

func (Container) TableName() string {
	return "container"
}

func FindContainer(user_id int64) Container {
	db := Openmysql()
	defer db.Close() //把数据库的连接关闭掉
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&Container{})
	// u1 := Container{123, "222", 111, 1, 123}
	// db.Create(&u1)
	var record Container
	err := db.Table("containers").Where("user_id = ?", user_id).Take(&record)
	if err != nil {
		panic(err)
	} else {
		return record
	}
}
func UseContainer(c *gin.Context) (Container, bool) { //能否使用容器,即密码是否正确
	keywords := c.PostFormArray("keywords") //获取输入数据
	userid_string := keywords[0]
	password := keywords[1]
	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型
	container := FindContainer(user_id)                   //找到对应容器
	container_password := container.Container_password    //查找密码
	if password != container_password {                   //密码不相等
		return container, false
	} else {
		return container, true
	}
}
func CreateRecord(container Container, start_time time.Time) { //创建一条使用记录进入数据库
	db := Openmysql()
	defer db.Close() //把数据库的连接关闭掉
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&Used_Record{})
	var record Used_Record
	err := db.Table("used_record").Where("user_id = ?", container.User_id).Last(&record)
	var total_time time.Duration
	rent_time := time.Now().Sub(start_time)
	if err != nil { //之前没有使用过该容器
		total_time = rent_time
	} else {
		used_time, _ := time.ParseDuration(record.Used_time)
		total_time = used_time + rent_time
	}
	record = Used_Record{
		User_id:    container.User_id,
		Machine_id: container.Machine_id,
		Start_time: start_time,
		End_time:   start_time,
		Rent_time:  rent_time.String(), //获得时间差
		Used_time:  total_time.String(),
	}
	db.Create(&record) //将这个记录增加
}
func UpdateRecord(user_id int64, end_time time.Time) { //退出容器后更新使用记录
	db := Openmysql()
	var record Used_Record
	err := db.Table("used_record").Where("user_id = ?", user_id).Last(&record) //查找最新的记录,开始时间相同
	if err != nil {
		panic(err)
	} else { //更新记录

		db.Table("used_record").Where("user_id= ? and start_time = ?", user_id, record.Start_time).Update("end_time", end_time.String())
		db.Table("used_record").Where("user_id= ? and start_time = ?", user_id, record.Start_time).Update("rent_time", end_time.Sub(record.Start_time))
		used_time, _ := time.ParseDuration(record.Used_time)
		rent_time, _ := time.ParseDuration(record.Rent_time)
		new_rentTime := (used_time - rent_time + end_time.Sub(record.Start_time)).String()
		db.Table("used_record").Where("user_id= ? and start_time = ?", user_id, record.Start_time).Update("used_time", new_rentTime)
	}
}
func UpdateContainerStatus(status int, container Container) { //更新容器的状态
	db := Openmysql()
	db.Table("containers").Where("image_id= ?", container.Image_ID).Update("container_status", container.Container_status)
}
