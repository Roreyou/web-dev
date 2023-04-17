package controller

import (
	"net/http"

	"back/dao"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.JSON(http.StatusOK, dao.GetGpuInfo(c))
}

func Link(c *gin.Context) {
	c.JSON(http.StatusOK, "将用户与对应机器相连")
}
