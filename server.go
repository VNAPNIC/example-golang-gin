package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"healthcare-panel/common"
	"healthcare-panel/middleware"
	model "healthcare-panel/models"
	"healthcare-panel/routers"
	"healthcare-panel/utils/logging"
	redisUtil "healthcare-panel/utils/redis"
	"healthcare-panel/utils/setting"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger = logging.Setup("main-logger", nil)

func init() {
	setting.Setup()
	model.Setup()
	common.InitValidate()
	if setting.RedisSetting.Host != "" {
		redisUtil.Setup()
	}
}

// @title Healthcare panel
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header like: Bearer xxxx
// @name Authorization
func main() {
	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery())
	if setting.AppSetting.EnabledCORS {
		g.Use(middleware.CORS())
	}

	routersInit := routers.InitRouter(g)

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		// service connection
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatalln(err)
		}
	}()

	conn, _ := net.Dial("ip:icmp", "localhost")
	fmt.Println(conn.LocalAddr())
	logger.Printf("[info] start http server listening %s", endPoint)
	logger.Printf("[info] Actual pid is %d", os.Getpid())

	// Wait for an interrupt signal to gracefully shut down the server (set a timeout of 5 seconds)
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}

	logger.Println("Server exiting")
}
