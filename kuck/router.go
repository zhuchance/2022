package main

import (
	"github.com/EDDYCJY/go-gin-example/controller"
	"github.com/gin-gonic/gin"
)

func CollerRoute(r *gin.Engine) *gin.Engine  {
	r.POST("/api/auth/register", controller.Register)
	return r
}