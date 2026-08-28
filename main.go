package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kouleen/common/bootstrap"
	"github.com/kouleen/common/middleware"
	"github.com/kouleen/idl/kitex_gen/rpc"
	"github.com/kouleen/idl/kitex_gen/user/userservice"
)

func main() {
	bootstrap.Run(rpc.USER_RPC_SERVER, func(option ...server.Option) server.Server {
		return userservice.NewServer(new(UserServiceImpl), option...)
	}, bootstrap.WithServerMiddleware(middleware.RpcServerMiddleware))
}
