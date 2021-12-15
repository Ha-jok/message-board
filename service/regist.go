package service

import (
	"message-board/dao"
)


func Regist(username,password,mibao,answer string)(error,bool){
	err := dao.InsertUser(username,password,mibao,answer)
	if err != nil {
		b := JudgeUsernameDuplicate(username,err.Error())
		return err,b
	}
	b := true
	return err,b
}

