package svc

import (
	"github.com/Guvanchhojamov/shorturl/gateway/api/internal/config"
	"github.com/Guvanchhojamov/shorturl/gateway/services/shorturlrpc/rpc/shorturlservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ShortUrlService shorturlservice.ShorturlService //manual code adding service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ShortUrlService: shorturlservice.NewShorturlService(zrpc.MustNewClient(c.ShortUrlRpcConfig)), //manual code service with configs
	}
}
