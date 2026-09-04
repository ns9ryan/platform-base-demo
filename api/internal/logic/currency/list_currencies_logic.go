// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package currency

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCurrenciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCurrenciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCurrenciesLogic {
	return &ListCurrenciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCurrenciesLogic) ListCurrencies(req *types.ListCurrenciesRequest) (resp *types.ListCurrenciesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
