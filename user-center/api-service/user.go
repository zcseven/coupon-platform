package main

import (
	"flag"
	"fmt"

	"coupon-platform/user-center/api-service/internal/config"
	"coupon-platform/user-center/api-service/internal/handler"
	"coupon-platform/user-center/api-service/internal/middleware"
	"coupon-platform/user-center/api-service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// //全局中间
	// server.Use(func(next http.HandlerFunc) http.HandlerFunc {
	// 	return func(w http.ResponseWriter, r *http.Request) {

	// 		next(w, r)
	// 	}
	// })
	server.Use(middleware.NewCorsMiddleware().Handle)
	server.Use(middleware.NewRequestMiddleware().Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
