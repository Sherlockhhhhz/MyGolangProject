package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// 启动服务
func Start(ctx context.Context, serviceName, host, port string, registerHandlesFunc func()) (context.Context, error) {
	registerHandlesFunc() // 启动函数
	ctx = startService(ctx, serviceName, host, port)
	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		log.Println(srv.ListenAndServe()) // 打印错误
		cancel()
	}()

	go func() {
		fmt.Printf("%v started, Press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s) // 等待用户的输入
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
