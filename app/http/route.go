package http

import (
	"github.com/wangyulu/web-go/app/http/module/demo"
	"github.com/wangyulu/web-go/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
