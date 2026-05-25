package logic

import (
	"context"

	"github.com/Guvanchhojamov/shorturl/gateway/services/shorturlrpc/rpc/internal/svc"
	"github.com/Guvanchhojamov/shorturl/gateway/services/shorturlrpc/rpc/shorturl"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *shorturl.ShortenReq) (*shorturl.ShortenResp, error) {
	// todo: add your logic here and delete this line

	return &shorturl.ShortenResp{}, nil
}
