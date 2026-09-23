package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"seckill/internal/config"
	"seckill/internal/infra"
	"seckill/internal/router"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Cmd() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	err = config.Init()
	if err != nil {
		logger.Error(err.Error())
	}
	mysql, err := infra.InitMysql()
	if err != nil {
		logger.Error(err.Error())
	}
	redis, err := infra.InitRedis()
	if err != nil {
		logger.Error(err.Error())
	}
	defer redis.Close()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	r := gin.Default()
	router.InitRouter(r, mysql)
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	errCh := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()
	select {
	case err := <-errCh:
		panic(err)
	case <-ctx.Done():
		stop()
	}
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	err = server.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error(err.Error())
		server.Close()
	}
	logger.Info(fmt.Sprintf("server shutdown success."))
}
