package router

import (
	"seckill/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRouter(router *gin.RouterGroup, mysql *gorm.DB) {
	userHandler := handler.NewUserHandler(mysql)
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
}
