package logic

import (
	"context"

	"github.com/Guvanchhojamov/shorturl/gateway/services/shorturlrpc/rpc/internal/svc"
	"github.com/Guvanchhojamov/shorturl/gateway/services/shorturlrpc/rpc/shorturl"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *shorturl.ExpandReq) (*shorturl.ExpandResp, error) {
	// todo: add your logic here and delete this line

	return &shorturl.ExpandResp{}, nil
}
