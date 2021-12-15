package service

import (
	"fmt"
	"message-board/dao"
)

//通过原密码修改密码，传入用户名原密码和新密码，返回错误和一个布尔值.供api层调用
func ChangePassword (username,oldPassword,newPassword string)(error,bool){
//	读取数据库中的信息。
	user,err := dao.ReadUser(username)
	//判断原密码是否正确，返回一个布尔值。若布尔值为true，调用dao层的updatepassword函数
	fmt.Println(user)
	b := JudgePassword(oldPassword,user.Password,user.Username,"修改密码")
	if b == true{
		dao.UpdatePassword(username,newPassword)
	}
	return err,b
}

//通过密保修改密码
func ForgetPassword(username,answer,newPassword string)(error,bool,string){
	//读取数据库中的信息，将该用户的密保问题及问题导出。
	user,err := dao.ReadUser(username)
	if user.Mibao==""{
		var b bool
		return err,b,user.Mibao
	}
	//判断密保答案是否正确，调用JudgeMibao接口，接受一个bool值
	b := JudgeMibao(answer,user.Answer,username,"忘记密码")
	if b == true{
		dao.UpdatePassword(username,newPassword)
	}
	return err,b,user.Password

}