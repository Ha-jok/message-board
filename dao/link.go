package dao

import (
	"database/sql"
)

//连接数据库
func LinkMysql()(*sql.DB,error){
	//定义一个字符串，供各个函数调用，来连接数据库
	var Str = "root:2002@/messageboard"
	db,error := sql.Open("mysql",Str)
	return db,error
}