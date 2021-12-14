package main

import (
	"message-board/api"
)

//注册，登陆，修改密码
//发表留言，修改留言，删除留言（删除后标记为该留言已删除，不删除其评论或回复，或者删除后标记为该留言已删除，不删除其评论或回复），查看所有留言，查看一个留言及其下属评论
//发表评论，修改评论，删除评论

//登录功能：
//在api下的login.go进行路由接口，供main调用
//在service下的login.go进行判断，并返回判断结果，供api调用
//在dao文件夹下的read.go读取数据库信息，用于被service调用。
//在model文件夹下的user.go创建符合数据库的用户结构体，


func main(){
	//创建一个路由，调用api层下的router接口
	engine := api.CreateRout()
	//实现登录功能,调用api层的login接口
	api.Login(engine)
}