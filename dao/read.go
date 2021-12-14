package dao

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"message-board/model"
)

//定义一个字符串，供各个函数调用，来连接数据库
var str = "root:2002@/messageboard"

//定义函数，传入用户名，读出该用户名相关信息，把得到的结构体和错误返回到service层
func ReadUser(u string) (model.User,error){
	db,err:=sql.Open("mysql",str)
	sqlstr:="select * from users where username=?"
	var user model.User
	err=db.QueryRow(sqlstr,u).Scan(&user.Username,&user.Password,&user.Mibao,&user.Answer)
	return user,err
}