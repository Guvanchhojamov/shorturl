package tools

import (
	"context"
	"github.com/Guvanchhojamov/shorturl/gateway/services/shorturlrpc/rpc/shorturlservice"

	"github.com/Guvanchhojamov/shorturl/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/shorturl/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortLogic {
	return &ShortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortLogic) Short(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	//shorten logic
	in := &shorturlservice.ShortenReq{Url: req.Url}
	shortened, err := l.svcCtx.ShortUrlService.Shorten(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &types.ShortenResp{Shorted: shortened.Expand}, nil
	//shorten logic end
}
