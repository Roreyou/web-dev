package controller

import (
	"net/http"

	"back/dao"

	"back/services"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.JSON(http.StatusOK, dao.GetGpuInfo(c))
}

func Link(c *gin.Context) {
	c.JSON(http.StatusOK, services.GreateDocker())
}
