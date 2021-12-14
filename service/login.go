package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"message-board/dao"
)

//创建登录主函数，传入用户输入信息用户名和密码哦，接收dao层的结构体和错误,将错误返回。
//供api层中的接口调用

func Login(un,pw string)(error,bool){
	u,err:=dao.ReadUser(un)
	b := judgePassword(pw,u.Password,u.Username)
	return err,b
}
//判断密码,在后台打印用户是否登录成功的信息信息,并返回一个布尔值，反应密码是否正确
func judgePassword(inputPw,mysqlPw,username string)bool{
	if inputPw != mysqlPw{
		fmt.Println(username,"登陆失败")
		b := false
		return b
	}else {
		fmt.Println(username,"已登录")
		b := true
		return b
	}
}