package dao

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Used_Record struct {
	User_id    int       `json:"使用者ID" db:"user_id"`
	Machine_id int       `json:"租用的机器ID" db:"machine_id"`
	Start_time time.Time `json:"机器开始使用时间" db:"start_time"`
	End_time   time.Time `json:"结束租用服务器时间" db:"end_time"`
	Rent_time  string    `json:"本次租用服务器的时间" db:"rent_time"`
	Used_time  string    `json:"总共租用服务器时间" db:"used_time"`
}

type Container struct {
	User_id            int    `json:"创建容器的用户" db:"user_id"`
	Image_ID           int    `json:"镜像ID" db:"image_id"`
	Container_port     int    `json:"容器端口" db:"container_port"`
	Container_status   int    `json:"容器的状态" db:"container_status"`
	Container_password string `json:"容器密码" db:"container_password"`
	Container_id       string `json:"容器ID" db:"container_id"`
	Container_ip       string `json:"容器所在ip" db:"container_ip"`
	Machine_id         int    `json:"对应机器ID" db:"machine_id"`
}

func (Container) TableName() string {
	return "container"
}
func (Used_Record) TableName() string {
	return "used_record"
}

func FindContainer(user_id int64) Container {
	db := Openmysql()
	defer db.Close() //把数据库的连接关闭掉
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&Container{})
	// u1 := Container{123, "222", 111, 1, 123}
	// db.Create(&u1)
	var record []Container
	db.Table("container").Where("user_id = ?", user_id).Find(&record)
	fmt.Println(record)
	// if err != nil {
	// 	panic(err)
	// }
	var container Container
	for i, newcontainer := range record {
		if newcontainer.Container_status != 1 {
			container = newcontainer
		}
		fmt.Println(i)
	}
	return container
}

func UseContainer(c *gin.Context) (Container, bool) { //能否使用容器,即密码是否正确
	//keywords := c.PostFormArray("keywords") //获取输入数据
	//userid_string := keywords[0]
	//password := keywords[1]
	userid_string := c.PostForm("user_id")
	password := c.PostForm("container_password")
	user_id, _ := strconv.ParseInt(userid_string, 10, 64) //要转化成int64类型
	container := FindContainer(user_id)                   //找到对应容器
	container_password := container.Container_password    //查找密码
	if password != container_password {                   //密码不相等
		return container, false
	} else {
		return container, true
	}
}

// func FindAllContainer(user_id int64) (containers []Container) { //查找该用户的所有使用容器
//
//		db := Openmysql()
//		defer db.Close() //把数据库的连接关闭掉
//		//2、把模型与数据库中的表对应起来
//		db.AutoMigrate(&Container{})
//		var container []Container
//		db.Table("container").Where("user_id = ?", user_id).Take(container)
//		return container
//	}
func FindMachine(machine_id int) (server Server_Info) { //返回服务器信息
	db := Openmysql()
	defer db.Close() //把数据库的连接关闭掉
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&Used_Record{})
	db.Table("server_info").Where("server_id = ?", machine_id).Take(&server)
	return server
}
func CreateRecord(container Container, start_time time.Time) { //创建一条使用记录进入数据库
	db := Openmysql()
	defer db.Close() //把数据库的连接关闭掉
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&Used_Record{})
	var allrecord []Used_Record
	var record Used_Record
	db.Table("used_record").Where("user_id = ?", container.User_id).Find(&allrecord)
	var total_time time.Duration
	rent_time := time.Now().Sub(start_time)
	// if err != nil {
	// 	panic(err)
	// }
	if len(allrecord) == 0 { //之前没有使用过该容器
		total_time = rent_time
	} else {
		record = allrecord[0]
		for i, r := range allrecord {
			if record.Start_time.Before(r.Start_time) {
				record = r
				fmt.Println(i)
				fmt.Println(record)
			}
		}
		used_time, _ := time.ParseDuration(record.Used_time)
		total_time = used_time + rent_time
		if record.Machine_id == container.Machine_id { //再次使用该机器
			newrent_time, _ := time.ParseDuration(record.Rent_time)
			rent_time += newrent_time
		}
	}
	record = Used_Record{
		User_id:    container.User_id,
		Machine_id: container.Machine_id,
		Start_time: start_time,
		End_time:   start_time,
		Rent_time:  rent_time.String(), //获得时间差
		Used_time:  total_time.String(),
	}
	fmt.Println(record)
	db.Create(&record) //将这个记录增加
}
func UpdateRecord(user_id int64, end_time time.Time) bool { //退出容器后更新使用记录
	db := Openmysql()
	//var record Used_Record
	//db.Table("used_record").Where("user_id = ?", user_id).First(&record) //查找最新的记录,开始时间相同
	// if err != nil {
	// 	panic(err)
	// } else { //更新记录
	var allrecord []Used_Record
	var record Used_Record
	db.Table("used_record").Where("user_id = ?", user_id).Find(&allrecord)
	record = allrecord[0]
	fmt.Println(record)
	for i, r := range allrecord {
		if record.Start_time.Before(r.Start_time) {
			record = r
			fmt.Println(i)
			fmt.Println(record)
		}
	}
	oldrecord := record
	if len(allrecord) >= 2 {
		oldrecord = allrecord[len(allrecord)-2]
	}
	end_time.Local().Second()
	fmt.Println(record)
	db.Table("used_record").Where("user_id= ? and start_time = ?", user_id, record.Start_time).Update("end_time", end_time)
	db.Table("used_record").Where("user_id= ? and start_time = ?", user_id, record.Start_time).Update("rent_time", (end_time.Sub(record.Start_time)).String())
	used_time, _ := time.ParseDuration(oldrecord.Used_time)
	rent_time, _ := time.ParseDuration(record.Rent_time)
	new_usedTime := (used_time + end_time.Sub(record.Start_time)).String()
	new_renttime := (rent_time + end_time.Sub(record.Start_time)).String()
	db.Table("used_record").Where("user_id= ? and start_time = ?", user_id, record.Start_time).Update("used_time", new_usedTime)
	//}
	db.Table("used_record").Where("user_id= ? and start_time = ?", user_id, record.Start_time).Update("rent_time", new_renttime)
	db.Table("used_record").Where("user_id= ? and start_time = ?", user_id, record.Start_time).Last(&record)
	return IsOutTime(record)
}
func UpdateContainerStatus(status int, container Container) { //更新容器的状态
	db := Openmysql()
	db.Table("container").Where("container_id = ?", container.Container_id).Update("container_status", status)
}
func IsOutTime(record Used_Record) bool { //容器使用是否超出额度
	used_hour := strings.Split(record.Used_time, "h")
	hour, _ := strconv.Atoi(used_hour[0])
	Cut_Remainder(record)
	return hour >= 10
}
func IsContainerUsing() bool { //是否有正在使用的容器
	db := Openmysql()
	var container Container
	db.Table("container").Where("container_status= ?", 3).First(&container)
	// if err != nil {
	// 	panic(err)
	// }
	if (container == Container{}) { //没有正在使用的容器
		return false
	} else {
		return true
	}
}

func ChangserverFlag(flag bool) {
	db := Openmysql()
	fmt.Println("在改了在改了")
	db.Table("server_info").Update("server_flag", flag)
}
