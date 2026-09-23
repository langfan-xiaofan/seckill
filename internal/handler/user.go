package handler

import (
	"net/http"
	"seckill/internal/dto"
	"seckill/internal/pkg/response"
	"seckill/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(mysql *gorm.DB) *UserHandler {
	return &UserHandler{
		svc: service.NewUserService(mysql),
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		response.Fail(c, http.StatusBadRequest, nil, "参数错误"+err.Error())
		return
	}
	err := h.svc.Register(c, req.Username, req.Password)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, nil, "内部服务错误"+err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, nil, "参数错误"+err.Error())
		return
	}
	res, err := h.svc.Login(c, req.Username, req.Password)
	if err != nil {
		response.Fail(c, http.StatusForbidden, nil, "账号或密码错误")
		return
	}
	response.Success(c, res)
	return
}
