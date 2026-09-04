// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package timezone

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTimezonesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTimezonesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTimezonesLogic {
	return &ListTimezonesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTimezonesLogic) ListTimezones(req *types.ListTimezonesRequest) (resp *types.ListTimezonesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
