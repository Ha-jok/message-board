package service

import "fmt"

//判断密码,在后台打印用户是否成功的信息信息,并返回一个布尔值，反应密码是否正确，
func JudgePassword(inputPw,mysqlPw,username,service string)bool{
	if inputPw != mysqlPw{
		fmt.Println(username,"密码错误",service)
		b := false
		return b
	}else {
		fmt.Println(username,"密码正确",service)
		b := true
		return b
	}
}

//判断密保，在后台打印是否成功的信息，返回一个布尔值，反应密保答案是否正确，
