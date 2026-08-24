package main

import (
	"log"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kouleen/idl/kitex_gen/user/userservice"
	"github.com/kouleen/user-center/middleware"
)

func main() {
	// etcd注册中心
	r, err := etcd.NewEtcdRegistry([]string{os.Getenv("ETCD_ENDPOINTS")})
	if err != nil {
		log.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", os.Getenv("ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: os.Getenv("SERVICE_NAME")}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
		server.WithMiddleware(middleware.RpcClientMiddleware),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
