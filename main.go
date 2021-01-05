package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gv/router"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("lxm"))
	// 使用session中间件
	r.Use(sessions.Sessions("gin-session", store))

	r.LoadHTMLGlob("template/**/*")
	r.Static("/static", "static")

	router.Router(r)
	s := &http.Server{
		Addr:         ":8090",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	_ = s.ListenAndServe()
}
