package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMsg(c *gin.Context, msg string) {
	url := "error?msg=" + msg
	c.Redirect(302, url)
}

func Error(c *gin.Context) {
	msg := c.Query("msg")
	c.HTML(http.StatusOK, "error.html", gin.H{
		"msg": msg,
	})
}
