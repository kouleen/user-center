package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kouleen/idl/kitex_gen/user/userheaderservice"
	"github.com/kouleen/user-center/utils"
)

func main() {
	if v := os.Getenv("REDIS_DB"); v != "" {
		if iv, err := strconv.Atoi(v); err == nil {
			utils.InitRedis(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), iv)
		}
	}
	// etcd注册中心
	r, err := etcd.NewEtcdRegistry([]string{os.Getenv("ETCD_ENDPOINTS")})
	if err != nil {
		log.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", os.Getenv("ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}
	svr := userheaderservice.NewServer(new(UserHeaderServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: os.Getenv("SERVICE_NAME")}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
