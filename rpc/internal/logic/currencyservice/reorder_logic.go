package currencyservicelogic

import (
	"context"

	"oa.98ent.com/p9/platform-base/rpc/internal/svc"
	"oa.98ent.com/p9/platform-base/rpc/pb/base/currency"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReorderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReorderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReorderLogic {
	return &ReorderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 调整货币排序
func (l *ReorderLogic) Reorder(in *currency.ReorderCurrencyRequest) (*currency.ReorderCurrencyResponse, error) {
	// todo: add your logic here and delete this line

	return &currency.ReorderCurrencyResponse{}, nil
}
