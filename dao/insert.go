package dao

//传入用户相关信息，将数据插入到数据库中。供service层的regist接口调用
func InsertUser(username,password,mibbao,answer string)error  {
	db,err := LinkMysql()
	sqlstr := "insert into users (username,password,mibao,answer) values(?,?,?,?)"
	_,err = db.Exec(sqlstr,username,password,mibbao,answer)
	return err
}