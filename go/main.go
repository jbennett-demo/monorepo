package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func newRouter(log *zap.Logger) *gin.Engine {
	r := gin.New()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.POST("/tokens", func(c *gin.Context) {
		id := uuid.NewString()
		log.Info("token issued", zap.String("id", id))
		c.JSON(http.StatusCreated, gin.H{"id": id})
	})
	return r
}

func main() {
	viper.SetDefault("addr", ":8080")
	log, _ := zap.NewProduction()
	defer log.Sync()

	if err := newRouter(log).Run(viper.GetString("addr")); err != nil {
		log.Fatal("server stopped", zap.Error(err))
	}
}
