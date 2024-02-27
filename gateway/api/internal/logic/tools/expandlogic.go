package tools

import (
	"context"
	"github.com/Guvanchhojamov/shorturl/gateway/services/shorturlrpc/rpc/shorturl"

	"github.com/Guvanchhojamov/shorturl/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/shorturl/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (resp *types.ExpandResp, err error) {
	//manual code expand url
	in := &shorturl.ExpandReq{Expand: req.Shorten}
	expanded, err := l.svcCtx.ShortUrlService.Expand(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &types.ExpandResp{Url: expanded.Url}, nil
	//end
}
