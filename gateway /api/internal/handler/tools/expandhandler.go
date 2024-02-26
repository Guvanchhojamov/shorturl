package tools

import (
	"net/http"

	"github.com/Guvanchhojamov/shorturl/gateway/api/internal/logic/tools"
	"github.com/Guvanchhojamov/shorturl/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/shorturl/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExpandHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tools.NewExpandLogic(r.Context(), svcCtx)
		resp, err := l.Expand(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
