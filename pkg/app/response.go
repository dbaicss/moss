package app

import (
	"github.com/gin-gonic/gin"
	"moss-service/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int			`json:"code"`
	Msg  string			`json:"msg"`
	Data interface{} 	`json:"data"`
}

// 返回Response格式∂
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode,Response{
		Code: 	errCode,
		Msg:	e.GetMsg(errCode),
		Data:	data,
	})
}