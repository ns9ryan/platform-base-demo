// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package timezone

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllTimezonesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAllTimezonesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllTimezonesLogic {
	return &ListAllTimezonesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAllTimezonesLogic) ListAllTimezones(req *types.ListAllTimezonesRequest) (resp *types.ListAllTimezonesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
