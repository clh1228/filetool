package main

import (
	"flag"
	"fmt"

	"filetool/api/config"
	"filetool/api/internal/handler"
	"filetool/api/internal/svc"
	database "filetool/database"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/rdata-sync-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Setup Db(Sqlite)
	database.SetupDatabase()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	// 启动redis Listen节点
	// go func() {
	// 	redisapi.InitP3000RedisListen(c)
	// }()

	// 启动消息订阅服务
	// go func() {
	// 	p3000.StartMessage(c)
	// }()

	server.Start()
}
