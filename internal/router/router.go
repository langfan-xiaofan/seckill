package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(router *gin.Engine, mysql *gorm.DB) {
	userRouter := router.Group("v1/auth")
	InitUserRouter(userRouter, mysql)
}
