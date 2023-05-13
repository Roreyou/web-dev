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
	Rent_time  time.Time `json:"本次租用服务器的时间"`
	Used_time  time.Time `json:"总共租用服务器时间"`
}

type Container struct {
	Image_ID           int    `json:"容器ID"`
	Container_password string `json:"容器密码"`
	Machine_id         int    `json:"对应机器ID"`
	Container_status   int    `json:"容器的状态"`
	User_id            int    `json:"创建容器的用户"`
}

func FindContainer(user_id int64) Container {
	db := Openmysql()
	var record Container
	err := db.Table("container").Where("user_id = ?", user_id).Take(&record)
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
