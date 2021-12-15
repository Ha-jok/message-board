package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"message-board/model"
)



//定义函数，传入用户名，读出该用户名相关信息，把得到的结构体和错误返回到service层
func ReadUser(u string) (model.User,error){
	db,err := LinkMysql()
	sqlstr:="select * from users where username=?"
	var user model.User
	err=db.QueryRow(sqlstr,u).Scan(&user.Username,&user.Password,&user.Mibao,&user.Answer)
	return user,err
}