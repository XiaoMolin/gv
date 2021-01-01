package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/**/*")
	r.Static("/static","static")
	//r.LoadHTMLFiles("index.html")
	r.GET("/", func(context *gin.Context) {
		//context.String(200, "asdd")
		name:="林小墨"
		context.HTML(200,"index/index.html",name)
	})


	_ = r.Run()
}
