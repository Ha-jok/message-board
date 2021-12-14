package service

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"message-board/dao"
)

//创建登录主函数，传入用户输入信息用户名和密码哦，接收dao层的结构体和错误,将错误返回。
func login(un,pw string)error{
	u,err:=dao.readUser(un)
	if u.password!=pw{
		fmt.Println("您输入的密码错误，请重新输入")
	}


}