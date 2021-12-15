package dao

import (
	_ "github.com/go-sql-driver/mysql"
)

//修改密码，传入用户名和新密码。返回错误。供service层调用
func UpdatePassword(username,newPassword string)error{
	db,err := LinkMysql()
	sqlstr := "update users set password=? where username=?"
	_,err = db.Exec(sqlstr,newPassword,username)
	return err
}



