package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DataBase struct {
		MysqlDns string
	}
	// Auth struct {
	// 	AccessSecret string
	// 	AccessExpire int64
	// }
}
