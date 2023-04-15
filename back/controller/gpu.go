package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gpu struct {
}

func (g *Gpu) Show(c *gin.Context) {
	c.JSON(http.StatusOK, "展示可用机器列表")
}

func (g *Gpu) Link(c *gin.Context) {
	c.JSON(http.StatusOK, "将用户与对应机器相连")
}
