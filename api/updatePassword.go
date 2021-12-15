package api

import (
	"github.com/gin-gonic/gin"
	"message-board/service"
	"net/http"
)

//创建中间件函数，调用service的接口，接收用户发送的信息，返回响应报文
func changePasswordMiddleware(c *gin.Context){
	//接收用户发送的信息
	username := c.PostForm("username")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	if len(newPassword)<3 {
		c.String(http.StatusOK,"密码长度不能小于3")
		return
	}
	if oldPassword==newPassword{
		c.String(http.StatusOK,"两次密码不能相同")
		return
	}
	//调用service中的change接口，接收错误和返回值
	err,b := service.ChangePassword(username,oldPassword,newPassword)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : err,
			"tip" : "请联系管理员",
		})
		return
	}
	if b == true {
		c.JSON(http.StatusOK,gin.H{
			"msg" : "修改成功",
			"newpassword" : newPassword,
		})
	}else {
		c.String(http.StatusOK,"修改失败，密码错误")
	}
}
//创建修改密码的路由接口
func ChangePassword (engine *gin.Engine){
	POST("/changeword",engine,changePasswordMiddleware)
}