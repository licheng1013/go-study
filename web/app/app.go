package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web/common"
)

func Run() {
	//下面都是优雅停机方式
	server := &http.Server{
		Addr:    common.AppConfig.Addr,
		Handler: common.App,
	}
	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("关闭失败:" + err.Error())
	}
	log.Println("服务器退出")
}
