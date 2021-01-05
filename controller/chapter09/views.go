package chapter09

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionTest(ctx *gin.Context){

	// 初始化session对象

	session := sessions.Default(ctx)

	// 设置session
	session.Set("name","lxm")

	// 获取session
	name :=session.Get("name")
	fmt.Println("=============")
	fmt.Println(name)

	// 删除指定key的session
	session.Delete("name")

	// 删除所有的session
	session.Clear()

	// 保存session
	_ = session.Save()
}