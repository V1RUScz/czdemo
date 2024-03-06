package svc

import (
	"czdemo/application/applet/internal/config"
	"czdemo/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
	UserRPC  user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRPC := zrpc.MustNewClient(c.UserRPC)
	return &ServiceContext{
		Config:   c,
		BizRedis: redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		UserRPC:  user.NewUser(userRPC),
	}
}
