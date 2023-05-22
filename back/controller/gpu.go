package controller

import (
	"fmt"
	"net/http"

	"back/dao"

	"back/services"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.JSON(http.StatusOK, dao.GetGpuInfo(c))
}

func Link(c *gin.Context) {
	machineid := c.Query("id")
	fmt.Println("看一下id")
	fmt.Println(machineid)
	c.JSON(http.StatusOK, services.GreateDocker(machineid))
}
