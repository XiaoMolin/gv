package chapter11

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gv/models"
	"net/http"
)



func ApiAxios(ctx *gin.Context){

	// 结构体
	user :=models.User{Id: 1,Name: "lxm",Age: 20}
	// 数组/切片吧
	arrs:=[]int{3,52,52,5,21}
	// 结构体数组
	// 结构体map

	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"成功",
		"user":user,
		"arrs":arrs,
	})

}


type Book struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
}

func GetBooks(ctx *gin.Context){
	books:=[]Book{
		{Id: 1,Name: "gin",Author: "lxm"},
		{Id: 2,Name: "gin2",Author: "lxm2"},
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ok",
		"books":books,
	})
}

func GetBookDetal(ctx *gin.Context){
	id :=ctx.Query("id")
	fmt.Println(id)
	// 假设查询了
	book:=Book{
		Id: 1,Name: "gin",Author: "lxm",
	}
	ctx.JSON(http.StatusOK,gin.H{
		"book":book,
	})
}