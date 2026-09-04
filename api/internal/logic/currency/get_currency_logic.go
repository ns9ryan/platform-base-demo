// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package currency

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrencyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrencyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrencyLogic {
	return &GetCurrencyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrencyLogic) GetCurrency(req *types.GetCurrencyRequest) (resp *types.GetCurrencyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
