package router

import (
	"github.com/gin-gonic/gin"
	"gv/controller/chapter09"
	"gv/controller/chapter11"
	"gv/middleware"
)

func Router(router *gin.Engine){
	chap09:=router.Group("/chapter09")
	chap11:=router.Group("/chapter11")
	chap11.Use(middleware.CrosMiddleWare)
	chapter09.Router(chap09)
	chapter11.Router(chap11)

}