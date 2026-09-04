// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package timezone

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReorderTimezoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReorderTimezoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReorderTimezoneLogic {
	return &ReorderTimezoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReorderTimezoneLogic) ReorderTimezone(req *types.ReorderTimezoneRequest) (resp *types.ReorderTimezoneResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
