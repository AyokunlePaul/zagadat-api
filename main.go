package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	zapLogger, _ := zap.NewDevelopment()
	router := gin.New()
	router.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(zapLogger, true))

	router.GET("/zadadis", func(context *gin.Context) {
		context.String(http.StatusOK, "Zagadat")
	})

	zapLogger.Fatal(router.Run(":8080").Error())
}
