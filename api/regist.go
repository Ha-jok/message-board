package api

import (
	"github.com/gin-gonic/gin"
	"message-board/service"
	"net/http"
)

//接收用户发送的信息，调用service中的regist接口，接收一个布尔值，根据布尔值，返回响应报文
func registMiddleware(c *gin.Context){
	//接收用户发送的信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	mibao := c.PostForm("mibao")
	answer := c.PostForm("answer")
	err,b := service.Regist(username,password,mibao,answer)
	if b == false{
		c.String(http.StatusOK,"用户名重复请使用其他用户名")
		return
	}
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error" : err,
			"tip" : "请联系管理员",
		})
		return
	}
	if len(password)<3{
		c.String(http.StatusOK,"密码长度不能小于3")
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"msg" : "注册成功",
		"username" : username,
		"password" : password,
	})
	return
}

func Regist(engine *gin.Engine){
	POST("/regist",engine,registMiddleware)
}

