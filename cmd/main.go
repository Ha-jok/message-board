package main

import (
	"message-board/api"
)

//注册，登陆，修改密码
//发表留言，修改留言，删除留言（删除后标记为该留言已删除，不删除其评论或回复，或者删除后标记为该留言已删除，不删除其评论或回复），查看所有留言，查看一个留言及其下属评论
//发表评论，修改评论，删除评论

//登录功能：
//在api下的login.go进行路由接口，供main调用
//在service下的login.go进行判断密码是否正确，并返回判断结果，供api调用
//在dao文件夹下的read.go读取数据库信息，用于被service调用。
//在model文件夹下的user.go创建符合数据库的用户结构体，

//修改密码功能，
//实现：1.通过原密码修改密码。2.通过密保答案修改密码。
//直接调用api层包装的两个接口,changepassword和forgetpassword
//api层接收用户发送的信息，调用service层的两个接口changepassword和forgetpassword进行处理，返回响应报文。
//service层接收dao层读取的数据库信息，并调用dao层中的函数，readUser,updateUser
//dao层read.go读取信息，updata.go修改信息。


func main(){
	engine := api.CreateRout()
	//实现登录
	api.Login(engine)
	//实现修改密码功能，调用api层的ChangePassword接口和ForgetPassword接口
	api.ChangePassword(engine)
	engine.Run()
}