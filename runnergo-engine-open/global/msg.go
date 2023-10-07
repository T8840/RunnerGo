package global

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnMsg(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
