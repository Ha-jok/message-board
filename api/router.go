package api

import (
	"github.com/gin-gonic/gin"
)

//创建路由,返回路由引擎指针，供main调用，
//并包装不同的请求方法,api文件夹下其他的接口结合

func CreateRout()*gin.Engine{
	engine := gin.Default()
	return engine
}

//func GET(engine *gin.Engine,relitivePath string,function func (context *gin.Context),c *gin.Context){
//	engine.GET(relitivePath, function(c) )
//	engine.Run("8080")
//}

func POST(relitivePath string,engine *gin.Engine,function func (context *gin.Context),c *gin.Context){
	engine.POST(relitivePath, function)
	engine.Run()
}

