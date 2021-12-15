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

//忘记密码中间件，调用service中的接口，接收用户发送的信息，返回响应报文。
func forgetPasswordMiddleware(c *gin.Context){
	//接收用户发送的信息，用户名，服务-忘记密码，密保答案,新密码
	username := c.PostForm("username")
	answer := c.PostForm("answer")
	newPassword := c.PostForm("newPassword")
	//判定用户传入新密码长度是否符合规范
	if len(newPassword)<3 {
		c.String(http.StatusOK,"密码长度不能小于3")
		return
	}
	//调用service中的forgetpassword接口，接收错误和布尔值，返回相应的响应报文
	err,b,judgeString := service.ForgetPassword(username,answer,newPassword)
	//首先判断该用户是否设置密保
	if judgeString == "" {
		c.String(http.StatusOK,"该用户未设置密保")
		return
	}
	//若密保已设置，处理传递的错误
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : err,
			"tip" : "请联系管理员",
		})
		return
	}
	//判断是否修改完成
	if b == false {
		c.String(http.StatusOK,"修改失败，密保错误")
		return
	}
	if judgeString == newPassword{
		c.String(http.StatusOK,"密码不能和原来相同")
		return
	}
	if b == true {
		c.JSON(http.StatusOK,gin.H{
			"msg" : "修改成功",
			"newpassword" : newPassword,
		})
	}
}

func ForgetPassword(engine *gin.Engine){
	POST("/forgetpassword",engine,forgetPasswordMiddleware)
}