package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DataBase struct {
		MysqlDns string
	}

	//rpc
	UserRpc zrpc.RpcClientConf
}
