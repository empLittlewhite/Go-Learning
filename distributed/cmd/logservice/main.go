// 这是入口文件
package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/emplittlewhite/distributed/log"
	"github.com/emplittlewhite/distributed/service"
)

func main() {
	log.Run("./distributed.log")
	// 按道理是应该使用配置文件传入的，但是入门教程就算了写死吧
	host, port := "localhost", "4000"
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	// 等待ctx的信号传入
	<-ctx.Done()
	fmt.Println("shutting down log service. ")
}
