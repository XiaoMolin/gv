package chapter11

import (
	"github.com/gin-gonic/gin"
)

func Router(chap09 *gin.RouterGroup){
	chap09.GET("/api_axios",ApiAxios)
	chap09.GET("/get_books",GetBooks)
	chap09.GET("/get_book_detail",GetBookDetal)
}
