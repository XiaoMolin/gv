package middleware

import (
	"github.com/gin-gonic/gin"
)

// 解决跨域请求的问题
func CrosMiddleWare(ctx *gin.Context){
	ctx.Header("Access-Control-Allow-Origin","*")// 允许跨域请求

	ctx.Next()


}
